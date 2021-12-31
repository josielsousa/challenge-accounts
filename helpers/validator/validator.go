package validator

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"

	"github.com/thedevsaddam/govalidator"

	httpHelper "github.com/josielsousa/challenge-accounts/helpers/http"
	"github.com/josielsousa/challenge-accounts/repo/model"
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
	govalidator.AddCustomRule("string", func(field, rule, message string, value interface{}) error {
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

	govalidator.AddCustomRule("cpf", func(field, rule, message string, value interface{}) error {
		// Parâmetros iniciais para validação:
		//	sizeWithoutDigits = Tamanho do CPF sem os dígitos informados.
		//	position = Peso inicial para validação dos dígitos.
		const (
			sizeWithoutDigits = 9
			position          = 10
		)

		helper := &Helper{}
		val := value.(string)

		if !helper.IsCPFValid(val, sizeWithoutDigits, position) {
			if message == "" {
				message = "The CPF is invalid."
			}

			return fmt.Errorf(message)
		}
		return nil
	})
}

// ValidateDataAccount - Realiza a validação dos dados enviados para as request de `account`.
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

// ValidateDataTransfer - Realiza a validação dos dados enviados para as request de `transfer`.
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

// ValidateDataLogin - Realiza a validação dos dados enviados para as request de `login`.
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

// IsCPFValid - Verifica se o CPF informado é válido, calculando os dígitos verificadores.
func (h *Helper) IsCPFValid(doc string, size, position int) bool {
	// Removes special characters.
	h.removeSpecialChars(&doc)

	// Se o documento estiver vazio, retorna falso
	if len(doc) <= 0 {
		return false
	}

	// Documento não é valido quando todos os dígitos forem iguais.
	if h.allEquals(doc) {
		return false
	}

	// Calcula o primeiro digito verificador.
	data := doc[:size]
	digit := calculateDigit(data, position)

	// Calcula o segundo digito verificador.
	data = data + digit
	digit = calculateDigit(data, position+1)

	return doc == data+digit
}

// removeSpecialChars - Remove os caracteres não númericos.
func (h *Helper) removeSpecialChars(value *string) string {
	if value == nil {
		return ""
	}

	reg, _ := regexp.Compile("[^0-9]+")
	normalized := reg.ReplaceAllString(*value, "")
	return normalized
}

// allEquals - Verifica se todos os dígitos do CPF são iguais, e.g.: 111.111.111-11
func (h *Helper) allEquals(value string) bool {
	base := value[0]
	for i := 1; i < len(value); i++ {
		if base != value[i] {
			return false
		}
	}

	return true
}

// calculateDigit - Calcula o digito verificador do documento informado conforme seu tipo.
//	position - representa o peso para a regra de cálculo do digito verificador.
// CPF pesos: 10, 9, 8, 7, 6, 5, 4, 3, 2
func calculateDigit(value string, position int) string {
	var sum int
	for _, r := range value {
		sum += int(r-'0') * position
		position--

		if position < 2 {
			position = 9
		}
	}

	sum %= 11
	if sum < 2 {
		return "0"
	}

	return strconv.Itoa(11 - sum)
}
