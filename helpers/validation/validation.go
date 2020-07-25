package validation

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	httpHelper "github.com/josielsousa/challenge-accounts/helpers/http"
	"github.com/josielsousa/challenge-accounts/repo/model"
	"github.com/josielsousa/challenge-accounts/types"
	"github.com/thedevsaddam/govalidator"
)

//Helper - Define a struct para o helper de validação.
type Helper struct {
	httpHlp *httpHelper.Helper
}

//NewHelper - Instância helper validação.
func NewHelper() *Helper {
	return &Helper{
		httpHlp: httpHelper.NewHelper(),
	}
}

// InitCustomRule - Inicializa a regra geral para verificar os campos do tipo `string`.
func InitCustomRule() {
	govalidator.AddCustomRule("string", func(field string, rule string, message string, value interface{}) error {
		if value == nil {
			return nil
		}

		typeField := reflect.TypeOf(value).String()
		err := fmt.Errorf(ErrorFieldNotString, field)
		if message != "" {
			err = errors.New(message)
		}

		if typeField != "string" {
			return err
		}

		return nil
	})
}

//ValidateDataAccount - Realiza a validação dos dados enviados para as request de `account`.
func (h *Helper) ValidateDataAccount(w http.ResponseWriter, req *http.Request) *model.Account {
	var account model.Account
	opts := govalidator.Options{
		Request:         req,
		RequiredDefault: true,
		Data:            &account,
		Rules:           validateRulesAccount,
		Messages:        validateMessagesAccount,
	}

	err := govalidator.New(opts).ValidateJSON()
	if len(err) > 0 {
		h.httpHlp.ThrowError(w, http.StatusUnprocessableEntity, err)
		return nil
	}

	return &account
}

//ValidateDataTransfer - Realiza a validação dos dados enviados para as request de `transfer`.
func (h *Helper) ValidateDataTransfer(w http.ResponseWriter, req *http.Request) *model.Transfer {
	var transfer model.Transfer
	opts := govalidator.Options{
		Request:         req,
		RequiredDefault: true,
		Data:            &transfer,
		Rules:           validateRulesTransfer,
		Messages:        validateMessagesTransfer,
	}

	err := govalidator.New(opts).ValidateJSON()
	if len(err) > 0 {
		h.httpHlp.ThrowError(w, http.StatusUnprocessableEntity, err)
		return nil
	}

	return &transfer
}

//ValidateDataLogin - Realiza a validação dos dados enviados para as request de `login`.
func (h *Helper) ValidateDataLogin(w http.ResponseWriter, req *http.Request) *types.Credentials {
	var credentials types.Credentials
	opts := govalidator.Options{
		Request:         req,
		RequiredDefault: true,
		Data:            &credentials,
		Rules:           validateRulesLogin,
		Messages:        validateMessagesLogin,
	}

	err := govalidator.New(opts).ValidateJSON()
	if len(err) > 0 {
		h.httpHlp.ThrowError(w, http.StatusUnprocessableEntity, err)
		return nil
	}

	return &credentials
}
