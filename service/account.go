package service

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/josielsousa/challenge-accounts/helpers/auth"
	httpHelper "github.com/josielsousa/challenge-accounts/helpers/http"
	"github.com/josielsousa/challenge-accounts/helpers/validator"
	"github.com/josielsousa/challenge-accounts/repo/model"
	"github.com/josielsousa/challenge-accounts/types"
)

// Mesnagens de validação para `accounts`.
const (
	ErrorAccountExist = "Já existe uma conta criada com o CPF informado."
)

//Inicializa as regras customizadas.
func init() {
	validator.InitCustomRule()
}

// AccountService - Implementação do service para as accounts.
type AccountService struct {
	authHlp      *auth.Helper
	httpHlp      *httpHelper.Helper
	validatorHlp *validator.Helper
	stgAccount   model.AccountStorage
	logger       types.APILogProvider
}

//NewAccountService - Instância o service com a dependência `log` inicializada.
func NewAccountService(stgAccount model.AccountStorage, log types.APILogProvider) *AccountService {
	return &AccountService{
		logger:       log,
		stgAccount:   stgAccount,
		authHlp:      auth.NewHelper(),
		httpHlp:      httpHelper.NewHelper(),
		validatorHlp: validator.NewHelper(),
	}
}

//InsertAccount - Realiza a inserção de uma account conforme os dados do `body` da requisição
//	200: Sucesso na inserção
//	422: Erro - Os dados de entrada são válidos porém existe uma `account` para o CPF informado.
//	500: Erro inesperado durante o processamento da requisição
func (s *AccountService) InsertAccount(w http.ResponseWriter, req *http.Request) {
	account := s.validatorHlp.ValidateDataAccount(w, req)
	if account == nil {
		return
	}

	accExist, err := s.stgAccount.GetAccountByCPF(account.Cpf)
	if err != nil {
		s.logger.Error("Error on verify account exists: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	exist := accExist != nil && len(accExist.ID) > 0
	if exist {
		s.logger.Info("Account exists")
		s.httpHlp.ThrowError(w, http.StatusUnprocessableEntity, ErrorAccountExist)
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
	acc, err := s.stgAccount.Insert(*account)
	if err != nil {
		s.logger.Error("Error on insert an account: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	s.httpHlp.ThrowSuccess(w, http.StatusCreated, types.SuccessResponse{Success: true, Data: acc})
}

//GetAllAccounts - Retorna as informações de todas as contas se não existir retorna []
//	200: Quando existir accounts para serem retornadas
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

//GetAccountBalance - Retorna as informações da account, conforme o id informado.
//	200: Quando existir account para ser retornada
//	404: Quando não encontrar a account.
//	500: Erro inesperado durante o processamento da requisição
func (s *AccountService) GetAccountBalance(w http.ResponseWriter, req *http.Request) {
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

	accountBalance := model.AccountBalance{Balance: account.Balance}
	s.httpHlp.ThrowSuccess(w, http.StatusOK, types.SuccessResponse{Success: success, Data: accountBalance})
}
