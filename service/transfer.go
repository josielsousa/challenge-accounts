package service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	httpHelper "github.com/josielsousa/challenge-accounts/helpers/http"
	"github.com/josielsousa/challenge-accounts/repo/db"
	"github.com/josielsousa/challenge-accounts/repo/model"
	"github.com/josielsousa/challenge-accounts/types"
)

// Mensagens domínio de transferencias.
const (
	ErrorOriginAccountNotFound      = "Conta de origem não encontrada"
	ErrorDestinationAccountNotFound = "Conta de destino não encontrada"
	ErrorInsufficientOriginBallance = "Conta de origem sem saldo disponível"
)

// TransferService - Implementação do service para as transfers.
type TransferService struct {
	stg     *db.Service
	logger  types.APILogProvider
	httpHlp *httpHelper.Helper
}

//NewTransferService - Instância o service com a dependência `log` inicializada.
func NewTransferService(stg *db.Service, log types.APILogProvider) *TransferService {
	return &TransferService{
		stg:     stg,
		logger:  log,
		httpHlp: httpHelper.NewHelper(),
	}
}

//DoTransfer - Realiza a transferência entre as `transfers` conforme a requisição
//	200: Sucesso na inserção
//	404: Quando a conta origem não for encontrada
//	404: Quando a conta destino não for encontrada
//	422: Quando não houver saldo disponível
//	500: Erro inesperado durante o processamento da requisição
func (s *TransferService) DoTransfer(w http.ResponseWriter, req *http.Request, claims *types.Claims) {
	transfer, err := s.JSONDecoder(req)
	if err != nil {
		s.logger.Error("Error on decode transfer: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	accOrigin, err := s.stg.Account.GetAccount(claims.AccountID)
	if err != nil {
		s.logger.Error("Error on recovery account origin transfer: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	if accOrigin == nil {
		s.httpHlp.ThrowError(w, http.StatusNotFound, ErrorOriginAccountNotFound)
		return
	}

	accDestination, err := s.stg.Account.GetAccount(transfer.AccountDestinationID)
	if err != nil {
		s.logger.Error("Error on recovery account destination transfer: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	if accDestination == nil {
		s.httpHlp.ThrowError(w, http.StatusNotFound, ErrorDestinationAccountNotFound)
		return
	}

	if accOrigin.Ballance <= 0 || accOrigin.Ballance < transfer.Amount {
		s.logger.Info(ErrorInsufficientOriginBallance)
		s.httpHlp.ThrowError(w, http.StatusUnprocessableEntity, ErrorInsufficientOriginBallance)
		return
	}

	//Inicia a transação
	tx := s.stg.BeginTransaction()
	accOrigin.Ballance -= transfer.Amount

	_, err = tx.Account.Update(*accOrigin)
	if err != nil {
		s.logger.Error("Error on update ballance account origin transfer: ", err)
		s.httpHlp.ThrowError(w, http.StatusNotFound, types.ErrorUnexpected)
		tx.Rollback()
		return
	}

	accDestination.Ballance += transfer.Amount
	_, err = tx.Account.Update(*accDestination)
	if err != nil {
		s.logger.Error("Error on update ballance account destination transfer: ", err)
		s.httpHlp.ThrowError(w, http.StatusNotFound, types.ErrorUnexpected)
		tx.Rollback()
		return
	}

	createdAt := time.Now()
	transfer.CreatedAt = &createdAt
	transfer.ID = uuid.New().String()

	trf, err := tx.Transfer.Insert(transfer)
	if err != nil {
		s.logger.Error("Error on insert an transfer: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		tx.Rollback()
		return
	}

	tx.Commit()
	s.httpHlp.ThrowSuccess(w, http.StatusCreated, types.SuccessResponse{Success: true, Data: trf})
}

//GetAllTransfers - Retorna as informações de todas as contas se não existir retorna []
// 	200: Quando existir transfers para serem retornadas
//	204: Quando não encontrar transfers.
//	500: Erro inesperado durante o processamento da requisição
func (s *TransferService) GetAllTransfers(w http.ResponseWriter, req *http.Request, claims *types.Claims) {
	transfers, err := s.stg.Transfer.GetAllTransfers(claims.AccountID)
	if err != nil {
		s.logger.Error("Error on get all transfers: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	statusCode := http.StatusOK
	if len(transfers) <= 0 {
		statusCode = http.StatusNoContent
	}

	s.httpHlp.ThrowSuccess(w, statusCode, types.SuccessResponse{Success: true, Data: transfers})
}

//JSONDecoder - Realiza o parser do body recebido da request.
func (s *TransferService) JSONDecoder(req *http.Request) (model.Transfer, error) {
	transfer := model.Transfer{}
	err := json.NewDecoder(req.Body).Decode(&transfer)
	return transfer, err
}
