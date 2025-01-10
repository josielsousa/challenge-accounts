package validator_test

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
//
// 	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
// 	"github.com/josielsousa/challenge-accounts/app/domain/entities/transfers"
// 	"github.com/josielsousa/challenge-accounts/helpers/validator"
// 	"github.com/josielsousa/challenge-accounts/types"
// )
//
// var (
// 	accountTest    accounts.Account
// 	transferTest   transfers.Transfer
// 	credentialTest types.Credentials
// 	validatorHlp   *validator.Helper
// )
//
// func init() {
// 	validator.InitCustomRule()
// }
//
// func setup() {
// 	accountTest = accounts.Account{
// 		Cpf:     "18447253082",
// 		Name:    "Teste Pessoa",
// 		Balance: 99.99,
// 		Secret:  "xxSecretXx",
// 	}
//
// 	transferTest = transfers.Transfer{
// 		Amount:               0.5,
// 		AccountDestinationID: "xxx-xxx-xxx",
// 	}
//
// 	credentialTest = types.Credentials{
// 		Cpf:    "18447253082",
// 		Secret: "xXSecretXx",
// 	}
//
// 	validatorHlp = validator.NewHelper()
// }
//
// func TestValidateDataAccount(t *testing.T) {
// 	setup()
//
// 	t.Run("Teste validate data account com sucesso", func(t *testing.T) {
// 		body, _ := json.Marshal(accountTest)
// 		bytesBody := bytes.NewReader(body)
// 		mockReq := httptest.NewRequest(http.MethodPost, "http://localhost:3000/accounts", bytesBody)
// 		mockRps := httptest.NewRecorder()
//
// 		acc := validatorHlp.ValidateDataAccount(mockRps, mockReq)
// 		if acc == nil {
// 			t.Error("Error on validate account request")
// 		}
// 	})
//
// 	t.Run("Teste validate data account com error", func(t *testing.T) {
// 		accountTest.Cpf = ""
// 		body, _ := json.Marshal(accountTest)
//
// 		bytesBody := bytes.NewReader(body)
// 		mockReq := httptest.NewRequest(http.MethodPost, "http://localhost:3000/accounts", bytesBody)
// 		mockRps := httptest.NewRecorder()
//
// 		acc := validatorHlp.ValidateDataAccount(mockRps, mockReq)
// 		if acc != nil || mockRps.Result().StatusCode != http.StatusUnprocessableEntity {
// 			t.Errorf("Error on validate account, required field not validate.")
// 		}
// 	})
// }
//
// func TestValidateDataTransfer(t *testing.T) {
// 	setup()
//
// 	t.Run("Teste validate data transfer com sucesso", func(t *testing.T) {
// 		body, _ := json.Marshal(transferTest)
// 		bytesBody := bytes.NewReader(body)
// 		mockReq := httptest.NewRequest(http.MethodPost, "http://localhost:3000/transfers", bytesBody)
// 		mockRps := httptest.NewRecorder()
//
// 		transfer := validatorHlp.ValidateDataTransfer(mockRps, mockReq)
// 		if transfer == nil {
// 			t.Error("Error on validate transfer request")
// 		}
// 	})
//
// 	t.Run("Teste validate data transfer com error", func(t *testing.T) {
// 		transferTest.AccountDestinationID = ""
// 		body, _ := json.Marshal(transferTest)
//
// 		bytesBody := bytes.NewReader(body)
// 		mockReq := httptest.NewRequest(http.MethodPost, "http://localhost:3000/transfers", bytesBody)
// 		mockRps := httptest.NewRecorder()
//
// 		transfer := validatorHlp.ValidateDataTransfer(mockRps, mockReq)
// 		if transfer != nil || mockRps.Result().StatusCode != http.StatusUnprocessableEntity {
// 			t.Errorf("Error on validate transfer, required field not validate.")
// 		}
// 	})
// }
//
// func TestValidateDataLogin(t *testing.T) {
// 	setup()
//
// 	t.Run("Teste validate data login com sucesso", func(t *testing.T) {
// 		body, _ := json.Marshal(credentialTest)
// 		bytesBody := bytes.NewReader(body)
// 		mockReq := httptest.NewRequest(http.MethodPost, "http://localhost:3000/login", bytesBody)
// 		mockRps := httptest.NewRecorder()
//
// 		credential := validatorHlp.ValidateDataLogin(mockRps, mockReq)
// 		if credential == nil {
// 			t.Error("Error on validate login request")
// 		}
// 	})
//
// 	t.Run("Teste validate data login com error", func(t *testing.T) {
// 		credentialTest.Secret = ""
// 		body, _ := json.Marshal(transferTest)
//
// 		bytesBody := bytes.NewReader(body)
// 		mockReq := httptest.NewRequest(http.MethodPost, "http://localhost:3000/login", bytesBody)
// 		mockRps := httptest.NewRecorder()
//
// 		credential := validatorHlp.ValidateDataLogin(mockRps, mockReq)
// 		if credential != nil || mockRps.Result().StatusCode != http.StatusUnprocessableEntity {
// 			t.Errorf("Error on validate login, required field not validate.")
// 		}
// 	})
// }
