package service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/josielsousa/challenge-accounts/helpers/auth"
	httpHelper "github.com/josielsousa/challenge-accounts/helpers/http"
	"github.com/josielsousa/challenge-accounts/repo/model"
	"github.com/josielsousa/challenge-accounts/types"
)

// AccountService - Implementação do service para as accounts.
type AccountService struct {
	authHlp    *auth.Helper
	httpHlp    *httpHelper.Helper
	stgAccount model.AccountStorage
	logger     types.APILogProvider
}

//NewAccountService - Instância o service com a dependência `log` inicializada.
func NewAccountService(stgAccount model.AccountStorage, log types.APILogProvider) *AccountService {
	return &AccountService{
		logger:     log,
		stgAccount: stgAccount,
		authHlp:    auth.NewHelper(),
		httpHlp:    httpHelper.NewHelper(),
	}
}

//InsertAccount - Realiza a inserção de uma account conforme os dados do `body` da requisição
//	200: Sucesso na inserção
//	500: Erro inesperado durante o processamento da requisição
func (s *AccountService) InsertAccount(w http.ResponseWriter, req *http.Request) {
	account, err := s.JSONDecoder(req)
	if err != nil {
		s.logger.Error("Error on decode account: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	createdAt := time.Now()
	account.CreatedAt = &createdAt
	account.ID = uuid.New().String()

	secretHash, err := s.authHlp.Hash(account.Secret)
	if err != nil {
		s.logger.Error("Error on hash secret account: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	account.Secret = string(secretHash)
	acc, err := s.stgAccount.Insert(account)
	if err != nil {
		s.logger.Error("Error on insert an account: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	s.httpHlp.ThrowSuccess(w, http.StatusCreated, types.SuccessResponse{Success: true, Data: acc})
}

//UpdateAccount - Realiza a atualização de uma account conforme os dados do `body` da requisição
//	200: Sucesso na atualização
//	500: Erro inesperado durante o processamento da requisição
func (s *AccountService) UpdateAccount(w http.ResponseWriter, req *http.Request) {
	account, err := s.JSONDecoder(req)
	if err != nil {
		s.logger.Error("Error on decode account: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	acc, err := s.stgAccount.Update(account)
	if err != nil {
		s.logger.Error("Error on update an account: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	s.httpHlp.ThrowSuccess(w, http.StatusOK, types.SuccessResponse{Success: true, Data: acc})
}

//GetAllAccounts - Retorna as informações de todas as contas se não existir retorna []
// 	200: Quando existir accounts para serem retornadas
//	204: Quando não encontrar accounts.
//	500: Erro inesperado durante o processamento da requisição
func (s *AccountService) GetAllAccounts(w http.ResponseWriter, req *http.Request) {
	accounts, err := s.stgAccount.GetAllAccounts()
	if err != nil {
		s.logger.Error("Error on get all accounts: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	statusCode := http.StatusOK
	if len(accounts) <= 0 {
		statusCode = http.StatusNoContent
	}

	s.httpHlp.ThrowSuccess(w, statusCode, types.SuccessResponse{Success: true, Data: accounts})
}

//GetAccount - Retorna as informações da account, conforme o id informado.
// 	200: Quando existir account para ser retornada
//	404: Quando não encontrar a account.
//	500: Erro inesperado durante o processamento da requisição
func (s *AccountService) GetAccount(w http.ResponseWriter, req *http.Request) {
	//Recebe os parâmetros da request e seleciona o ID
	params := s.httpHlp.GetParams(req)
	id := params["id"]

	account, err := s.stgAccount.GetAccount(id)
	if err != nil {
		s.logger.Error("Error on get account: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	success := account != nil && len(account.ID) > 0
	if !success {
		s.httpHlp.ThrowError(w, http.StatusNotFound, types.ErroAccountNotFound)
		return
	}

	s.httpHlp.ThrowSuccess(w, http.StatusOK, types.SuccessResponse{Success: success, Data: account})
}

//GetAccountBallance - Retorna as informações da account, conforme o id informado.
// 	200: Quando existir account para ser retornada
//	404: Quando não encontrar a account.
//	500: Erro inesperado durante o processamento da requisição
func (s *AccountService) GetAccountBallance(w http.ResponseWriter, req *http.Request) {
	//Recebe os parâmetros da request e seleciona o ID
	params := s.httpHlp.GetParams(req)
	id := params["id"]

	account, err := s.stgAccount.GetAccount(id)
	if err != nil {
		s.logger.Error("Error on get account: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	success := account != nil && len(account.ID) > 0
	if !success {
		s.httpHlp.ThrowError(w, http.StatusNotFound, types.ErroAccountNotFound)
		return
	}

	accountBallance := model.Account{Ballance: account.Ballance}
	s.httpHlp.ThrowSuccess(w, http.StatusOK, types.SuccessResponse{Success: success, Data: accountBallance})
}

//JSONDecoder - Realiza o parser do body recebido da request.
func (s *AccountService) JSONDecoder(req *http.Request) (model.Account, error) {
	account := model.Account{}
	err := json.NewDecoder(req.Body).Decode(&account)
	return account, err
}
