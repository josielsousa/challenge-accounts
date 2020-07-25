package service_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/josielsousa/challenge-accounts/repo/model"
	"github.com/josielsousa/challenge-accounts/repo/model/mocks"
	srv "github.com/josielsousa/challenge-accounts/service"
	"github.com/josielsousa/challenge-accounts/types"
)

const (
	ErrorScenarioError         = "Deveria retornar status 500; retornou %d"
	ErrorScenarioSuccess       = "Deveria retornar status 200; retornou %d"
	ErrorScenarioErrorNotFound = "Deveria retornar status 404; retornou %d"
)

var (
	stgAcc      *mocks.MockAccountStorage
	logAcc      *types.MockAPILogProvider
	srvAcc      *srv.AccountService
	accountTest model.Account
)

func mockLogAcc() {
	logAcc = &types.MockAPILogProvider{}
	logAcc.OnInfo = func(info string) {}
	logAcc.OnError = func(trace string, erro error) {}
}

func mockAccountStorage() model.AccountStorage {
	accountTest = model.Account{
		Cpf:      "XXXX",
		Name:     "Teste Pessoa",
		Ballance: 99.99,
		Secret:   "xxSecretXx",
	}

	stgAcc = &mocks.MockAccountStorage{
		OnGetAccount: func(id string) (*model.Account, error) {
			return &accountTest, nil
		},

		OnGetAccountByCPF: func(cpf string) (*model.Account, error) {
			return &accountTest, nil
		},

		OnUpdate: func(account model.Account) (*model.Account, error) {
			return &accountTest, nil
		},

		OnInsert: func(account model.Account) (*model.Account, error) {
			return &accountTest, nil
		},

		OnGetAllAccounts: func() (accounts []model.Account, err error) {
			accounts = make([]model.Account, 0)
			accounts = append(accounts, accountTest)
			return accounts, nil
		},
	}

	return stgAcc
}

func setupAccountService() *srv.AccountService {
	mockLogAcc()
	mockAccountStorage()
	return srv.NewAccountService(stgAcc, logAcc)
}

func TestServiceInsertAccount(t *testing.T) {
	t.Run("Teste Inserir account sucesso", func(t *testing.T) {
		//FakeBody para request
		srvAcc = setupAccountService()
		body, _ := json.Marshal(accountTest)

		bytesBody := bytes.NewReader(body)
		mockReq := httptest.NewRequest(http.MethodPost, "http://localhost:3000/accounts", bytesBody)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		srvAcc.InsertAccount(mockRps, mockReq)
		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusCreated {
			t.Errorf(ErrorScenarioSuccess, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste Inserir account erro inesperado", func(t *testing.T) {
		//FakeBody para request
		srvAcc = setupAccountService()
		body, _ := json.Marshal(accountTest)

		bytesBody := bytes.NewReader(body)
		mockReq := httptest.NewRequest(http.MethodPost, "http://localhost:3000/accounts", bytesBody)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		//Força o retorno de um erro inesperado ao realizar a persistência no banco de dados.
		stgAcc.OnInsert = func(account model.Account) (*model.Account, error) {
			return nil, errors.New("Erro Inesperado")
		}

		srvAcc.InsertAccount(mockRps, mockReq)

		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusInternalServerError {
			t.Errorf(ErrorScenarioError, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste Inserir account erro on GetAccountByCPF", func(t *testing.T) {
		//FakeBody para request
		srvAcc = setupAccountService()
		body, _ := json.Marshal(accountTest)

		bytesBody := bytes.NewReader(body)
		mockReq := httptest.NewRequest(http.MethodPost, "http://localhost:3000/accounts", bytesBody)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		//Força o retorno da account
		stgAcc.OnGetAccountByCPF = func(id string) (*model.Account, error) {
			return nil, errors.New(types.ErrorUnexpected)
		}

		srvAcc.InsertAccount(mockRps, mockReq)

		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusInternalServerError {
			t.Errorf(ErrorScenarioError, mockRps.Result().StatusCode)
		}
	})
}

func TestServiceGetAllAccount(t *testing.T) {
	t.Run("Teste Get All account sucesso", func(t *testing.T) {
		//FakeBody para request
		srvAcc = setupAccountService()
		mockReq := httptest.NewRequest(http.MethodGet, "http://localhost:3000/accounts", nil)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		srvAcc.GetAllAccounts(mockRps, mockReq)
		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusOK {
			t.Errorf(ErrorScenarioSuccess, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste Get All account lista vazia", func(t *testing.T) {
		srvAcc = setupAccountService()
		mockReq := httptest.NewRequest(http.MethodGet, "http://localhost:3000/accounts", nil)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		//Força o retorno da account
		stgAcc.OnGetAllAccounts = func() ([]model.Account, error) {
			accounts := make([]model.Account, 0)
			return accounts, nil
		}

		srvAcc.GetAllAccounts(mockRps, mockReq)

		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusNoContent {
			t.Errorf(ErrorScenarioErrorNotFound, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste Get All account erro inesperado", func(t *testing.T) {
		srvAcc = setupAccountService()
		mockReq := httptest.NewRequest(http.MethodGet, "http://localhost:3000/accounts", nil)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		//Força o retorno da account
		stgAcc.OnGetAllAccounts = func() ([]model.Account, error) {
			return nil, errors.New("Erro Inesperado")
		}

		srvAcc.GetAllAccounts(mockRps, mockReq)

		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusInternalServerError {
			t.Errorf(ErrorScenarioError, mockRps.Result().StatusCode)
		}
	})
}

func TestServiceGetAccountBallance(t *testing.T) {
	t.Run("Teste Get account sucesso", func(t *testing.T) {
		//FakeBody para request
		srvAcc = setupAccountService()
		mockReq := httptest.NewRequest(http.MethodGet, "http://localhost:3000/accounts/1/ballance", nil)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		//Gera o id
		accountTest.ID = uuid.New().String()

		//Força o retorno da account
		stgAcc.OnGetAccount = func(id string) (*model.Account, error) {
			return &accountTest, nil
		}

		srvAcc.GetAccountBallance(mockRps, mockReq)
		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusOK {
			t.Errorf(ErrorScenarioSuccess, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste Get account erro não encontrado", func(t *testing.T) {
		srvAcc = setupAccountService()
		mockReq := httptest.NewRequest(http.MethodGet, "http://localhost:3000/accounts/1/ballance", nil)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		//Força não encontrar a account
		stgAcc.OnGetAccount = func(id string) (*model.Account, error) {
			return nil, nil
		}

		srvAcc.GetAccountBallance(mockRps, mockReq)

		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusNotFound {
			t.Errorf(ErrorScenarioErrorNotFound, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste Get account erro inesperado", func(t *testing.T) {
		srvAcc = setupAccountService()
		mockReq := httptest.NewRequest(http.MethodGet, "http://localhost:3000/accounts/1/ballance", nil)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		//Força não encontrar a account
		stgAcc.OnGetAccount = func(id string) (*model.Account, error) {
			return nil, errors.New("Erro Inesperado")
		}

		srvAcc.GetAccountBallance(mockRps, mockReq)

		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusInternalServerError {
			t.Errorf(ErrorScenarioError, mockRps.Result().StatusCode)
		}
	})
}
