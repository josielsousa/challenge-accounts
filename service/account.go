package service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	httpHelper "github.com/josielsousa/challenge-accounts/helpers/http"
	"github.com/josielsousa/challenge-accounts/repo/model"
	"github.com/josielsousa/challenge-accounts/types"
)

// AccountService - Implementação do service para as accounts.
type AccountService struct {
	stgAccount model.AccountStorage
	logger     types.APILogProvider
	httpHlp    *httpHelper.Helper
}

//NewAccountService - Instância o service com a dependência `log` inicializada.
func NewAccountService(stgAccount model.AccountStorage, log types.APILogProvider) *AccountService {
	return &AccountService{
		logger:     log,
		stgAccount: stgAccount,
		httpHlp:    httpHelper.NewHelper(),
	}
}

//InsertAccount - Realiza a inserção de uma account conforme os dados do `body` da requisição
//	200: Sucesso na inserção
//	500: Erro inesperado durante o processamento da requisição
func (s *AccountService) InsertAccount(w http.ResponseWriter, req *http.Request) {
	account, err := s.JSONDecoder(req)
	if err != nil {
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	account.ID = uuid.New().String()
	account.CreatedAt = time.Now()

	acc, err := s.stgAccount.Insert(account)
	if err != nil {
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	s.httpHlp.ThrowSuccess(w, http.StatusOK, types.SuccessResponse{Success: true, Data: acc})
}

//UpdateAccount - Realiza a atualização de uma account conforme os dados do `body` da requisição
//	200: Sucesso na atualização
//	500: Erro inesperado durante o processamento da requisição
func (s *AccountService) UpdateAccount(w http.ResponseWriter, req *http.Request) {
	account, err := s.JSONDecoder(req)
	if err != nil {
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	acc, err := s.stgAccount.Update(account)
	if err != nil {
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
// 	200: Quando existir accounts para serem retornadas
//	404: Quando não encontrar accounts.
//	500: Erro inesperado durante o processamento da requisição
func (s *AccountService) GetAccount(w http.ResponseWriter, req *http.Request) {
	//Recebe os parâmetros da request e seleciona o ID
	params := s.httpHlp.GetParams(req)
	id := params["id"]

	account, err := s.stgAccount.GetAccount(id)
	if err != nil {
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	statusCode := http.StatusOK
	if account == nil || len(account.ID) <= 0 {
		statusCode = http.StatusNotFound
	}

	s.httpHlp.ThrowSuccess(w, statusCode, types.SuccessResponse{Success: true, Data: account})
}

//JSONDecoder - Realiza o parser do body recebido da request.
func (s *AccountService) JSONDecoder(req *http.Request) (model.Account, error) {
	account := model.Account{}
	err := json.NewDecoder(req.Body).Decode(&account)
	return account, err
}
