package service

import (
	"net/http"

	httpHelper "github.com/josielsousa/challenge-accounts/helpers/http"
	"github.com/josielsousa/challenge-accounts/helpers/validator"
	"github.com/josielsousa/challenge-accounts/types"
)

// Mensagens domínio de transferencias.
const (
	ErrorOriginAccountNotFound      = "Conta de origem não encontrada"
	ErrorDestinationAccountNotFound = "Conta de destino não encontrada"
	ErrorInsufficientOriginBalance  = "Conta de origem sem saldo disponível"
)

// TransferService - Implementação do service para as transfers.
type TransferService struct {
	stg          any
	logger       types.APILogProvider
	httpHlp      *httpHelper.Helper
	validatorHlp *validator.Helper
}

// NewTransferService - Instância o service com a dependência `log` inicializada.
func NewTransferService(stg any, log types.APILogProvider) *TransferService {
	return &TransferService{
		stg:          stg,
		logger:       log,
		httpHlp:      httpHelper.NewHelper(),
		validatorHlp: validator.NewHelper(),
	}
}

// DoTransfer - Realiza a transferência entre as `accounts` conforme os dados enviados na requisição.
//
//	200: Sucesso na inserção
//	400: Quando o `token` estiver vazio / nulo
//	401: Quando o `token` fornecido for inválido.
//	404: Quando a `account` origem não for encontrada
//	404: Quando a `account` destino não for encontrada
//	422: Quando não houver saldo disponível
//	500: Erro inesperado durante o processamento da requisição
func (s *TransferService) DoTransfer(w http.ResponseWriter, req *http.Request, claims *types.Claims) {
	// transfer := s.validatorHlp.ValidateDataTransfer(w, req)
	// if transfer == nil {
	// 	return
	// }
	//
	// accOrigin, err := s.stg.Account.GetAccount(claims.AccountID)
	// if err != nil {
	// 	s.logger.Error("Error on recovery account origin transfer: ", err)
	// 	s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
	// 	return
	// }
	//
	// successAccOrigin := accOrigin != nil && len(accOrigin.ID) > 0
	// if !successAccOrigin {
	// 	s.httpHlp.ThrowError(w, http.StatusNotFound, ErrorOriginAccountNotFound)
	// 	return
	// }
	//
	// accDestination, err := s.stg.Account.GetAccount(transfer.AccountDestinationID)
	// if err != nil {
	// 	s.logger.Error("Error on recovery account destination transfer: ", err)
	// 	s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
	// 	return
	// }
	//
	// successAccDestination := accDestination != nil && len(accDestination.ID) > 0
	// if !successAccDestination {
	// 	s.httpHlp.ThrowError(w, http.StatusNotFound, ErrorDestinationAccountNotFound)
	// 	return
	// }
	//
	// if accOrigin.Balance <= 0 || accOrigin.Balance < transfer.Amount {
	// 	s.logger.Info(ErrorInsufficientOriginBalance)
	// 	s.httpHlp.ThrowError(w, http.StatusUnprocessableEntity, ErrorInsufficientOriginBalance)
	// 	return
	// }
	//
	// // Inicia a transação
	// tx := s.stg.BeginTransaction()
	// accOrigin.Balance = decimal.Sub(transfer.Amount, accOrigin.Balance)
	//
	// _, err = tx.Account.Update(*accOrigin)
	// if err != nil {
	// 	s.logger.Error("Error on update balance account origin transfer: ", err)
	// 	s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
	// 	tx.Rollback()
	// 	return
	// }
	//
	// accDestination.Balance = decimal.Add(accDestination.Balance, transfer.Amount)
	// _, err = tx.Account.Update(*accDestination)
	// if err != nil {
	// 	s.logger.Error("Error on update balance account destination transfer: ", err)
	// 	s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
	// 	tx.Rollback()
	// 	return
	// }
	//
	// createdAt := time.Now()
	// transfer.CreatedAt = &createdAt
	// transfer.ID = uuid.New().String()
	// transfer.AccountOriginID = claims.AccountID
	//
	// trf, err := tx.Transfer.Insert(*transfer)
	// if err != nil {
	// 	s.logger.Error("Error on insert an transfer: ", err)
	// 	s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
	// 	tx.Rollback()
	// 	return
	// }
	//
	// tx.Commit()
	// s.httpHlp.ThrowSuccess(w, http.StatusCreated, types.SuccessResponse{Success: true, Data: trf})
}

// GetAllTransfers - Retorna as informações de todas as contas se não existir retorna []
//
//	200: Quando existir transfers para serem retornadas
//	204: Quando não encontrar transfers.
//	400: Quando o token estiver vazio / nulo
//	401: Quando o `token` fornecido for inválido.
//	500: Erro inesperado durante o processamento da requisição
func (s *TransferService) GetAllTransfers(w http.ResponseWriter, req *http.Request, claims *types.Claims) {
	// transfers, err := s.stg.Transfer.GetAllTransfers(claims.AccountID)
	// if err != nil {
	// 	s.logger.Error("Error on get all transfers: ", err)
	// 	s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
	// 	return
	// }
	//
	// statusCode := http.StatusOK
	// if len(transfers) <= 0 {
	// 	statusCode = http.StatusNoContent
	// }
	//
	// s.httpHlp.ThrowSuccess(w, statusCode, types.SuccessResponse{Success: true, Data: transfers})
}
