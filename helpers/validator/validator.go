package validator

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/thedevsaddam/govalidator"

	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/entities/transfers"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
	httpHelper "github.com/josielsousa/challenge-accounts/helpers/http"
	"github.com/josielsousa/challenge-accounts/types"
)

// Helper - Define a struct para o helper de validação.
type Helper struct {
	httpHlp *httpHelper.Helper
}

// NewHelper - Instância helper validação.
func NewHelper() *Helper {
	return &Helper{
		httpHlp: httpHelper.NewHelper(),
	}
}

// InitCustomRule - Inicializa a regra geral para verificar os campos do tipo `string`.
func InitCustomRule() {
	govalidator.AddCustomRule("string", func(field, _, message string, value interface{}) error {
		if value == nil {
			return nil
		}

		typeField := reflect.TypeOf(value).String()

		err := fmt.Errorf("o campo %s deve ser um string", field)
		if message != "" {
			err = errors.New(message)
		}

		if typeField != "string" {
			return err
		}

		return nil
	})

	govalidator.AddCustomRule("cpf", func(_, _, _ string, value interface{}) error {
		val, ok := value.(string)
		if !ok {
			return errors.New("o campo cpf deve ser uma string")
		}

		_, err := cpf.NewCPF(val)
		if err != nil {
			return errors.New("número de cpf inválido")
		}

		return nil
	})
}

// ValidateDataAccount - Realiza a validação dos dados enviados para as request de `account`.
func (h *Helper) ValidateDataAccount(writer http.ResponseWriter, req *http.Request) *accounts.Account {
	var account accounts.Account
	opts := govalidator.Options{
		Request:         req,
		RequiredDefault: true,
		Data:            &account,
		Rules:           validateRulesAccount,
		Messages:        validateMessagesAccount,
	}

	err := govalidator.New(opts).ValidateJSON()
	if len(err) > 0 {
		h.httpHlp.ThrowError(writer, http.StatusUnprocessableEntity, err)

		return nil
	}

	return &account
}

// ValidateDataTransfer - Realiza a validação dos dados enviados para as request de `transfer`.
func (h *Helper) ValidateDataTransfer(writer http.ResponseWriter, req *http.Request) *transfers.Transfer {
	var transfer transfers.Transfer
	opts := govalidator.Options{
		Request:         req,
		RequiredDefault: true,
		Data:            &transfer,
		Rules:           validateRulesTransfer,
		Messages:        validateMessagesTransfer,
	}

	err := govalidator.New(opts).ValidateJSON()
	if len(err) > 0 {
		h.httpHlp.ThrowError(writer, http.StatusUnprocessableEntity, err)

		return nil
	}

	return &transfer
}

// ValidateDataLogin - Realiza a validação dos dados enviados para as request de `login`.
func (h *Helper) ValidateDataLogin(writer http.ResponseWriter, req *http.Request) *types.Credentials {
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
		h.httpHlp.ThrowError(writer, http.StatusUnprocessableEntity, err)

		return nil
	}

	return &credentials
}
