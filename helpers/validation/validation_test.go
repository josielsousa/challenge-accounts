package validation_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/josielsousa/challenge-accounts/helpers/validation"
	"github.com/josielsousa/challenge-accounts/repo/model"
	"github.com/josielsousa/challenge-accounts/types"
)

var (
	accountTest    model.Account
	transferTest   model.Transfer
	credentialTest types.Credentials
	validationHlp  *validation.Helper
)

func init() {
	validation.InitCustomRule()
}

func setup() {
	accountTest = model.Account{
		Cpf:      "XXXX",
		Name:     "Teste Pessoa",
		Ballance: 99.99,
		Secret:   "xxSecretXx",
	}

	transferTest = model.Transfer{
		Amount:               0.5,
		AccountDestinationID: "xxx-xxx-xxx",
	}

	credentialTest = types.Credentials{
		Cpf:    "xxxx",
		Secret: "xXSecretXx",
	}

	validationHlp = validation.NewHelper()
}

func TestValidateDataAccount(t *testing.T) {
	setup()

	t.Run("Teste validate data account com sucesso", func(t *testing.T) {
		body, _ := json.Marshal(accountTest)
		bytesBody := bytes.NewReader(body)
		mockReq := httptest.NewRequest(http.MethodPost, "http://localhost:3000/accounts", bytesBody)
		mockRps := httptest.NewRecorder()

		acc := validationHlp.ValidateDataAccount(mockRps, mockReq)
		if acc == nil {
			t.Error("Error on validate account request")
		}
	})

	t.Run("Teste validate data account com error", func(t *testing.T) {
		accountTest.Cpf = ""
		body, _ := json.Marshal(accountTest)

		bytesBody := bytes.NewReader(body)
		mockReq := httptest.NewRequest(http.MethodPost, "http://localhost:3000/accounts", bytesBody)
		mockRps := httptest.NewRecorder()

		acc := validationHlp.ValidateDataAccount(mockRps, mockReq)
		if acc != nil || mockRps.Result().StatusCode != http.StatusUnprocessableEntity {
			t.Errorf("Error on validate account, required field not validate.")
		}
	})
}

func TestValidateDataTransfer(t *testing.T) {
	setup()

	t.Run("Teste validate data transfer com sucesso", func(t *testing.T) {
		body, _ := json.Marshal(transferTest)
		bytesBody := bytes.NewReader(body)
		mockReq := httptest.NewRequest(http.MethodPost, "http://localhost:3000/transfers", bytesBody)
		mockRps := httptest.NewRecorder()

		transfer := validationHlp.ValidateDataTransfer(mockRps, mockReq)
		if transfer == nil {
			t.Error("Error on validate transfer request")
		}
	})

	t.Run("Teste validate data transfer com error", func(t *testing.T) {
		transferTest.AccountDestinationID = ""
		body, _ := json.Marshal(transferTest)

		bytesBody := bytes.NewReader(body)
		mockReq := httptest.NewRequest(http.MethodPost, "http://localhost:3000/transfers", bytesBody)
		mockRps := httptest.NewRecorder()

		transfer := validationHlp.ValidateDataTransfer(mockRps, mockReq)
		if transfer != nil || mockRps.Result().StatusCode != http.StatusUnprocessableEntity {
			t.Errorf("Error on validate transfer, required field not validate.")
		}
	})
}

func TestValidateDataLogin(t *testing.T) {
	setup()

	t.Run("Teste validate data login com sucesso", func(t *testing.T) {
		body, _ := json.Marshal(credentialTest)
		bytesBody := bytes.NewReader(body)
		mockReq := httptest.NewRequest(http.MethodPost, "http://localhost:3000/login", bytesBody)
		mockRps := httptest.NewRecorder()

		credential := validationHlp.ValidateDataLogin(mockRps, mockReq)
		if credential == nil {
			t.Error("Error on validate login request")
		}
	})

	t.Run("Teste validate data login com error", func(t *testing.T) {
		credentialTest.Secret = ""
		body, _ := json.Marshal(transferTest)

		bytesBody := bytes.NewReader(body)
		mockReq := httptest.NewRequest(http.MethodPost, "http://localhost:3000/login", bytesBody)
		mockRps := httptest.NewRecorder()

		credential := validationHlp.ValidateDataLogin(mockRps, mockReq)
		if credential != nil || mockRps.Result().StatusCode != http.StatusUnprocessableEntity {
			t.Errorf("Error on validate login, required field not validate.")
		}
	})
}
