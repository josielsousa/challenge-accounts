package service_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/josielsousa/challenge-accounts/repo/db"
	"github.com/josielsousa/challenge-accounts/repo/model"
	"github.com/josielsousa/challenge-accounts/repo/model/mocks"
	srv "github.com/josielsousa/challenge-accounts/service"
	"github.com/josielsousa/challenge-accounts/types"
	_ "github.com/mattn/go-sqlite3"
)

const (
	ErrorScenarioSuccessNoContent = "Deveria retornar status 204; retornou %d"
)

var (
	srvDb               *db.Service
	transferTest        model.Transfer
	accountTransferTest model.Account
	claims              *types.Claims
	srvTransfer         *srv.TransferService
	logTransfer         *types.MockAPILogProvider
	stgAccTransfer      *mocks.MockAccountStorage
	stgTransfer         *mocks.MockTransferStorage
)

func mockLogTransfer() {
	logTransfer = &types.MockAPILogProvider{}
	logTransfer.OnInfo = func(info string) {}
	logTransfer.OnError = func(trace string, erro error) {}
}

func mockTransferStorage() *db.Service {
	originalID := uuid.New().String()
	claims = &types.Claims{
		Username:  "18447253082",
		AccountID: originalID,
	}

	accountTransferTest = model.Account{
		ID:      originalID,
		Cpf:     "18447253082",
		Name:    "Teste Pessoa",
		Balance: 99.99,
		Secret:  secret,
	}

	transferTest = model.Transfer{
		Amount:               0.99,
		AccountOriginID:      uuid.New().String(),
		AccountDestinationID: uuid.New().String(),
	}

	stgTransfer = &mocks.MockTransferStorage{
		OnInsert: func(transfer model.Transfer) (*model.Transfer, error) {
			return &transferTest, nil
		},

		OnGetAllTransfers: func(accountId string) ([]model.Transfer, error) {
			transfers := make([]model.Transfer, 0)
			transfers = append(transfers, transferTest)
			return transfers, nil
		},
	}

	stgAccTransfer = &mocks.MockAccountStorage{
		OnGetAccount: func(id string) (*model.Account, error) {
			return &accountTransferTest, nil
		},

		OnUpdate: func(account model.Account) (*model.Account, error) {
			return &accountTransferTest, nil
		},
	}

	srvDb, _ = db.Open(db.Gorm)
	srvDb.Account = stgAccTransfer
	srvDb.Transfer = stgTransfer
	return srvDb
}

func setupTransferService() *srv.TransferService {
	mockLogTransfer()
	mockTransferStorage()
	return srv.NewTransferService(srvDb, logTransfer)
}

func getMockRequestTransfer(method string) (*httptest.ResponseRecorder, *http.Request) {
	body, _ := json.Marshal(transferTest)
	bytesBody := bytes.NewReader(body)
	mockReq := httptest.NewRequest(method, "http://localhost:3000/transfers", bytesBody)
	mockRps := httptest.NewRecorder()
	return mockRps, mockReq
}

func TestDoTransfers(t *testing.T) {
	t.Run("Teste realizar uma transferência com sucesso", func(t *testing.T) {
		//FakeBody para request
		srvTransfer = setupTransferService()
		mockRps, mockReq := getMockRequestTransfer(http.MethodPost)

		srvTransfer.DoTransfer(mockRps, mockReq, claims)
		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusCreated {
			t.Errorf(ErrorScenarioCreated, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste realizar uma transferência com erro parâmetros de entrada 422", func(t *testing.T) {
		//FakeBody para request
		srvTransfer = setupTransferService()
		transferTest = model.Transfer{}
		mockRps, mockReq := getMockRequestTransfer(http.MethodPost)

		srvTransfer.DoTransfer(mockRps, mockReq, claims)
		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusUnprocessableEntity {
			t.Errorf(ErrorScenarioUnprocessableEntity, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste realizar uma transferência com erro 500", func(t *testing.T) {
		//FakeBody para request
		srvTransfer = setupTransferService()
		mockRps, mockReq := getMockRequestTransfer(http.MethodPost)

		//Força o retorno da account
		stgAccTransfer.OnGetAccount = func(id string) (*model.Account, error) {
			return nil, errors.New(types.ErrorUnexpected)
		}

		srvTransfer.DoTransfer(mockRps, mockReq, claims)
		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusInternalServerError {
			t.Errorf(ErrorScenarioError, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste realizar uma transferência com erro not found 404", func(t *testing.T) {
		//FakeBody para request
		srvTransfer = setupTransferService()
		mockRps, mockReq := getMockRequestTransfer(http.MethodPost)

		//Força o retorno da account
		stgAccTransfer.OnGetAccount = func(id string) (*model.Account, error) {
			return nil, nil
		}

		srvTransfer.DoTransfer(mockRps, mockReq, claims)
		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusNotFound {
			t.Errorf(ErrorScenarioErrorNotFound, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste realizar uma transferência com saldo insuficiente", func(t *testing.T) {
		//FakeBody para request
		srvTransfer = setupTransferService()
		transferTest.Amount = 0.5
		accountTransferTest.Balance = 0.2
		mockRps, mockReq := getMockRequestTransfer(http.MethodPost)

		srvTransfer.DoTransfer(mockRps, mockReq, claims)
		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusUnprocessableEntity {
			t.Errorf(ErrorScenarioUnprocessableEntity, mockRps.Result().StatusCode)
		}
	})
}

func TestGetAllTransfers(t *testing.T) {
	t.Run("Teste recuperar todas as transfências com sucesso", func(t *testing.T) {
		//FakeBody para request
		srvTransfer = setupTransferService()
		mockRps, mockReq := getMockRequestTransfer(http.MethodGet)

		srvTransfer.GetAllTransfers(mockRps, mockReq, claims)
		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusOK {
			t.Errorf(ErrorScenarioSuccess, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste recuperar todas as transfências com sucesso, porem sem dados", func(t *testing.T) {
		//FakeBody para request
		//FakeBody para request
		srvTransfer = setupTransferService()
		mockRps, mockReq := getMockRequestTransfer(http.MethodGet)

		//Força o retorno da transfers vazia
		stgTransfer.OnGetAllTransfers = func(accountId string) ([]model.Transfer, error) {
			transfers := make([]model.Transfer, 0)
			return transfers, nil
		}

		srvTransfer.GetAllTransfers(mockRps, mockReq, claims)
		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusNoContent {
			t.Errorf(ErrorScenarioSuccessNoContent, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste recuperar todas as transfências com erro", func(t *testing.T) {
		//FakeBody para request
		//FakeBody para request
		srvTransfer = setupTransferService()
		mockRps, mockReq := getMockRequestTransfer(http.MethodGet)

		//Força o retorno da transfers vazia
		stgTransfer.OnGetAllTransfers = func(accountId string) ([]model.Transfer, error) {
			transfers := make([]model.Transfer, 0)
			return transfers, errors.New("Erro Inesperado")
		}

		srvTransfer.GetAllTransfers(mockRps, mockReq, claims)
		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusInternalServerError {
			t.Errorf(ErrorScenarioError, mockRps.Result().StatusCode)
		}
	})
}
