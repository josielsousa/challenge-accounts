package service_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/josielsousa/challenge-accounts/helpers/auth"
	"github.com/josielsousa/challenge-accounts/repo/model"
	"github.com/josielsousa/challenge-accounts/repo/model/mocks"
	srv "github.com/josielsousa/challenge-accounts/service"
	"github.com/josielsousa/challenge-accounts/types"
)

const (
	ErrorGetNewToken           = "Erro inesperado ao recuperar novo token"
	ErrorTokenExpired          = "Token expirado. Deveria retornar status 401; retornou %d"
	ErrorTokenInvalid          = "Token inválido. Deveria retornar status 401; retornou %d"
	ErrorScenarioUnauthorized  = "Deveria retornar status 401; retornou %d"
	ErrorTokenErrorUnexpected  = "Token inválido. Deveria retornar status 500; retornou %d"
	ErrorTokenSignatureInvalid = "Token inválido. Deveria retornar status 401; retornou %d"
)

var (
	secret          = "xxSecretXx"
	stgAccAuth      *mocks.MockAccountStorage
	logAuth         *types.MockAPILogProvider
	srvAuth         *srv.AuthService
	accountAuthTest model.Account
	credentialTest  types.Credentials
)

func mockLogAuth() {
	logAuth = &types.MockAPILogProvider{}
	logAuth.OnInfo = func(info string) {}
	logAuth.OnError = func(trace string, erro error) {}
}

func mockAccAuthStorage() model.AccountStorage {
	accountAuthTest = model.Account{
		ID:      uuid.New().String(),
		Cpf:     "18447253082",
		Name:    "Teste Pessoa",
		Balance: 99.99,
		Secret:  secret,
	}

	stgAccAuth = &mocks.MockAccountStorage{
		OnGetAccountByCPF: func(cpf string) (*model.Account, error) {
			return &accountAuthTest, nil
		},
	}

	return stgAccAuth
}

func setupAuthService() *srv.AuthService {
	mockLogAuth()
	mockAccAuthStorage()

	credentialTest = types.Credentials{
		Cpf:    "18447253082",
		Secret: secret,
	}

	secretHashed, _ := auth.NewHelper().Hash(secret)
	accountAuthTest.Secret = string(secretHashed)

	return srv.NewAuthService(stgAccAuth, logAuth)
}

func getMockRequestAuth() (*httptest.ResponseRecorder, *http.Request) {
	body, _ := json.Marshal(credentialTest)
	bytesBody := bytes.NewReader(body)
	mockReq := httptest.NewRequest(http.MethodPost, "http://localhost:3000/login", bytesBody)
	mockRps := httptest.NewRecorder()
	return mockRps, mockReq
}

func TestAuthServiceLogin(t *testing.T) {
	t.Run("Teste login sucesso", func(t *testing.T) {
		//FakeBody para request
		srvAuth = setupAuthService()
		mockRps, mockReq := getMockRequestAuth()

		srvAuth.Login(mockRps, mockReq)
		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusOK {
			t.Errorf(ErrorScenarioSuccess, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste login secret diferente error", func(t *testing.T) {
		//FakeBody para request
		srvAuth = setupAuthService()

		credentialTest.Secret = "new_secret"
		mockRps, mockReq := getMockRequestAuth()

		srvAuth.Login(mockRps, mockReq)
		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusUnauthorized {
			t.Errorf(ErrorScenarioUnauthorized, mockRps.Result().StatusCode)
		}
	})
}

func TestAuthServiceValidateToken(t *testing.T) {
	t.Run("Teste validate token sucesso", func(t *testing.T) {
		//FakeBody para request
		srvAuth = setupAuthService()
		mockRps, mockReq := getMockRequestAuth()

		//Tempo de expiração do token
		jwtKey := []byte("api-challenge-accounts")
		expirationTime := time.Now().Add(1 * time.Minute)
		accessToken, err := srvAuth.GetToken(&accountAuthTest, jwtKey, expirationTime)
		if err != nil {
			t.Error(ErrorGetNewToken, err)
			return
		}

		mockReq.Header.Add("Access-Token", accessToken.Token)
		retAuth := srvAuth.ValidateToken(func(wNext http.ResponseWriter, reqNext *http.Request, claims *types.Claims) {

		})

		retAuth(mockRps, mockReq)

		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusOK {
			t.Errorf(ErrorScenarioSuccess, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste validate token expirado", func(t *testing.T) {
		//FakeBody para request
		srvAuth = setupAuthService()
		mockRps, mockReq := getMockRequestAuth()

		//Tempo de expiração do token
		jwtKey := []byte("api-challenge-accounts")
		expirationTime := time.Now().Add(1 * time.Millisecond)
		accessToken, err := srvAuth.GetToken(&accountAuthTest, jwtKey, expirationTime)
		if err != nil {
			t.Error(ErrorGetNewToken, err)
			return
		}

		//Força a expiração do token
		time.Sleep(2 * time.Millisecond)

		mockReq.Header.Add("Access-Token", accessToken.Token)
		retAuth := srvAuth.ValidateToken(func(wNext http.ResponseWriter, reqNext *http.Request, claims *types.Claims) {

		})

		retAuth(mockRps, mockReq)

		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusUnauthorized {
			t.Errorf(ErrorTokenExpired, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste validate token erro inesperado", func(t *testing.T) {
		//FakeBody para request
		srvAuth = setupAuthService()
		mockRps, mockReq := getMockRequestAuth()

		mockReq.Header.Add("Access-Token", "token_invalid")
		retAuth := srvAuth.ValidateToken(func(wNext http.ResponseWriter, reqNext *http.Request, claims *types.Claims) {

		})

		retAuth(mockRps, mockReq)

		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusInternalServerError {
			t.Errorf(ErrorTokenErrorUnexpected, mockRps.Result().StatusCode)
		}
	})

	t.Run("Teste validate token signature invalido", func(t *testing.T) {
		//FakeBody para request
		srvAuth = setupAuthService()
		mockRps, mockReq := getMockRequestAuth()

		tokenInvalid := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjAxNjgxNjY4MTYwIiwiYWNjb3VudF9pZCI6ImY5MmFiYzg2LTU2ZGEtNDMxZi04YTAyLWM0ODlkZDI2NWUwMyIsImV4cGlyZXNfYXQiOjE1OTU2ODg3NTV9.EalkFJT0IyjR1RfPfW5Rbsx1jxhviF1lPOrsGqkqq:)`
		mockReq.Header.Add("Access-Token", tokenInvalid)
		retAuth := srvAuth.ValidateToken(func(wNext http.ResponseWriter, reqNext *http.Request, claims *types.Claims) {

		})

		retAuth(mockRps, mockReq)

		//Verificação do comportamento de acordo com o cenário
		if mockRps.Result().StatusCode != http.StatusUnauthorized {
			t.Errorf(ErrorTokenSignatureInvalid, mockRps.Result().StatusCode)
		}
	})
}
