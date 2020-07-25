package service

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
	"github.com/josielsousa/challenge-accounts/types"
)

const (
	ErrorScenarioError         = "Deveria retornar status 500; retornou %d"
	ErrorScenarioSuccess       = "Deveria retornar status 200; retornou %d"
	ErrorScenarioErrorNotFound = "Deveria retornar status 404; retornou %d"
)

var (
	stg         *mocks.MockAccountStorage
	logAcc      *types.MockAPILogProvider
	service     *AccountService
	accountTest model.Account
)

func mockLog() {
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

	stg = &mocks.MockAccountStorage{
		OnGetAccount: func(id string) (*model.Account, error) {
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

	return stg
}

func setup() *AccountService {
	mockLog()
	mockAccountStorage()
	return NewAccountService(stg, logAcc)
}

func TestServiceInsertAccount(t *testing.T) {
	t.Run("Teste Inserir account sucesso", func(t *testing.T) {
		//FakeBody para request
		service = setup()
		body, _ := json.Marshal(accountTest)

		bytesBody := bytes.NewReader(body)
		mockReq := httptest.NewRequest(http.MethodPost, "http://localhost:8080/accounts", bytesBody)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		service.InsertAccount(mockRps, mockReq)
		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusCreated {
			t.Errorf(ErrorScenarioSuccess, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste Inserir account erro inesperado", func(t *testing.T) {
		//FakeBody para request
		service = setup()
		body, _ := json.Marshal(accountTest)

		bytesBody := bytes.NewReader(body)
		mockReq := httptest.NewRequest(http.MethodPost, "http://localhost:8080/accounts", bytesBody)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		//Força o retorno de um erro inesperado ao realizar a persistência no banco de dados.
		stg.OnInsert = func(account model.Account) (*model.Account, error) {
			return nil, errors.New("Erro Inesperado")
		}

		service.InsertAccount(mockRps, mockReq)

		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusInternalServerError {
			t.Errorf(ErrorScenarioError, mockRps.Result().StatusCode)
		}
	})
}

func TestServiceUpdateAccount(t *testing.T) {
	t.Run("Teste Update account sucesso", func(t *testing.T) {
		//FakeBody para request
		service = setup()
		body, _ := json.Marshal(accountTest)

		bytesBody := bytes.NewReader(body)
		mockReq := httptest.NewRequest(http.MethodPut, "http://localhost:8080/accounts", bytesBody)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		service.UpdateAccount(mockRps, mockReq)
		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusOK {
			t.Errorf(ErrorScenarioSuccess, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste Update account erro inesperado", func(t *testing.T) {
		//FakeBody para request
		service = setup()
		body, _ := json.Marshal(accountTest)

		bytesBody := bytes.NewReader(body)
		mockReq := httptest.NewRequest(http.MethodPut, "http://localhost:8080/accounts", bytesBody)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		//Força o retorno de um erro inesperado ao realizar a persistência no banco de dados.
		stg.OnUpdate = func(account model.Account) (*model.Account, error) {
			return nil, errors.New("Erro Inesperado")
		}

		service.UpdateAccount(mockRps, mockReq)

		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusInternalServerError {
			t.Errorf(ErrorScenarioError, mockRps.Result().StatusCode)
		}
	})
}

func TestServiceGetAccount(t *testing.T) {
	t.Run("Teste Get account sucesso", func(t *testing.T) {
		//FakeBody para request
		service = setup()
		mockReq := httptest.NewRequest(http.MethodGet, "http://localhost:8080/accounts", nil)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		//Gera o id
		accountTest.ID = uuid.New().String()

		//Força o retorno da account
		stg.OnGetAccount = func(id string) (*model.Account, error) {
			return &accountTest, nil
		}

		service.GetAccount(mockRps, mockReq)
		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusOK {
			t.Errorf(ErrorScenarioSuccess, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste Get account erro não encontrado", func(t *testing.T) {
		service = setup()
		mockReq := httptest.NewRequest(http.MethodGet, "http://localhost:8080/accounts", nil)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		//Força não encontrar a account
		stg.OnGetAccount = func(id string) (*model.Account, error) {
			return nil, nil
		}

		service.GetAccount(mockRps, mockReq)

		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusNotFound {
			t.Errorf(ErrorScenarioErrorNotFound, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste Get account erro inesperado", func(t *testing.T) {
		service = setup()
		mockReq := httptest.NewRequest(http.MethodGet, "http://localhost:8080/accounts", nil)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		//Força não encontrar a account
		stg.OnGetAccount = func(id string) (*model.Account, error) {
			return nil, errors.New("Erro Inesperado")
		}

		service.GetAccount(mockRps, mockReq)

		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusInternalServerError {
			t.Errorf(ErrorScenarioError, mockRps.Result().StatusCode)
		}
	})
}

func TestServiceGetAllAccount(t *testing.T) {
	t.Run("Teste Get All account sucesso", func(t *testing.T) {
		//FakeBody para request
		service = setup()
		mockReq := httptest.NewRequest(http.MethodGet, "http://localhost:8080/accounts", nil)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		service.GetAllAccounts(mockRps, mockReq)
		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusOK {
			t.Errorf(ErrorScenarioSuccess, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste Get All account lista vazia", func(t *testing.T) {
		service = setup()
		mockReq := httptest.NewRequest(http.MethodGet, "http://localhost:8080/accounts", nil)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		//Força o retorno da account
		stg.OnGetAllAccounts = func() ([]model.Account, error) {
			accounts := make([]model.Account, 0)
			return accounts, nil
		}

		service.GetAllAccounts(mockRps, mockReq)

		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusNoContent {
			t.Errorf(ErrorScenarioErrorNotFound, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste Get All account erro inesperado", func(t *testing.T) {
		service = setup()
		mockReq := httptest.NewRequest(http.MethodGet, "http://localhost:8080/accounts", nil)

		//Mock writer para teste
		mockRps := httptest.NewRecorder()

		//Força o retorno da account
		stg.OnGetAllAccounts = func() ([]model.Account, error) {
			return nil, errors.New("Erro Inesperado")
		}

		service.GetAllAccounts(mockRps, mockReq)

		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusInternalServerError {
			t.Errorf(ErrorScenarioError, mockRps.Result().StatusCode)
		}
	})
}
