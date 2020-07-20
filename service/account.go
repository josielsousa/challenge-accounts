package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
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
		s.logger.Error("Error on decode account: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	createdAt := time.Now()
	account.CreatedAt = &createdAt
	account.ID = uuid.New().String()

	acc, err := s.stgAccount.Insert(account)
	if err != nil {
		s.logger.Error("Error on insert an account: ", err)
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
// 	200: Quando existir accounts para serem retornadas
//	404: Quando não encontrar accounts.
//	500: Erro inesperado durante o processamento da requisição
func (s *AccountService) GetAccount(w http.ResponseWriter, req *http.Request) {
	//Recebe os parâmetros da request e seleciona o ID
	params := s.httpHlp.GetParams(req)
	id := params["id"]

	s.logger.Info(fmt.Sprintf("ID: %s", id))

	account, err := s.getAccount(id)
	if err != nil {
		s.logger.Error("Error on get account: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	success := account != nil && len(account.ID) > 0
	statusCode := http.StatusOK
	if !success {
		statusCode = http.StatusNotFound
	}

	s.httpHlp.ThrowSuccess(w, statusCode, types.SuccessResponse{Success: success, Data: account})
}

//JSONDecoder - Realiza o parser do body recebido da request.
func (s *AccountService) getAccount(id string) (*model.Account, error) {
	account, err := s.stgAccount.GetAccount(id)
	if err != nil {
		notFound, _ := regexp.MatchString(types.ErrorRecordNotFound, err.Error())
		if !notFound {
			return nil, err
		}

		return &model.Account{}, nil
	}

	return account, nil
}

//JSONDecoder - Realiza o parser do body recebido da request.
func (s *AccountService) JSONDecoder(req *http.Request) (model.Account, error) {
	account := model.Account{}
	err := json.NewDecoder(req.Body).Decode(&account)
	return account, err
}
