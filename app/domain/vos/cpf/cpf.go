package cpf

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/josielsousa/challenge-accounts/app/common"
)

const (
	firstOneDigitPosition = 10
	maxCPFLen             = 11

	weigthPosition = 2
	dividend       = 2
)

var (
	ErrInvalid = errors.New("invalid cpf number")

	regexPatternOnlyNumbers = regexp.MustCompile(`[^0-9]+`)
	regexPatternCPFMasked   = regexp.MustCompile(`^(\d{3})(\d{3})(\d{3})(\d{2})$`)
)

type CPF struct {
	value string
}

func NewCPF(numCPF string) (CPF, error) {
	cpf := CPF{
		value: removeSpecialChars(numCPF),
	}

	if ok := cpf.isValid(); !ok {
		return CPF{}, ErrInvalid
	}

	return cpf, nil
}

// removeSpecialChars - Remove os caracteres não númericos.
func removeSpecialChars(value string) string {
	return regexPatternOnlyNumbers.ReplaceAllString(value, "")
}

// allEquals - Verifica se todos os dígitos do CPF são iguais, e.g.: 111.111.111-11.
func (c *CPF) allEquals() bool {
	base := c.value[0]
	for i := range len(c.value) {
		if base != c.value[i] {
			return false
		}
	}

	return true
}

// calculateDigit - Calcula o digito verificador do documento informado conforme seu tipo.
//
//	position - representa o peso para a regra de cálculo do digito verificador.
//
// CPF pesos: 10, 9, 8, 7, 6, 5, 4, 3, 2.
func (c *CPF) calculateDigit(position int) string {
	var sum int

	data := c.value[:position-1]

	for _, digit := range data {
		sum += int(digit-'0') * position
		position--

		if position < weigthPosition {
			position = 9
		}
	}

	sum %= maxCPFLen
	if sum < dividend {
		return "0"
	}

	return strconv.Itoa(maxCPFLen - sum)
}

// IsValid - Verifica se o CPF informado é válido, calculando os dígitos verificadores.
func (c *CPF) isValid() bool {
	if len(c.value) != maxCPFLen || c.allEquals() {
		return false
	}

	doc := c.value[:9]
	firstOne := c.calculateDigit(firstOneDigitPosition)
	secondOne := c.calculateDigit(firstOneDigitPosition + 1)

	doc = fmt.Sprintf("%s%s%s", doc, firstOne, secondOne)

	return c.value == doc
}

// mask retorna o CPF informado no padrão (000.000.000-00).
func (c *CPF) Mask() string {
	return regexPatternCPFMasked.ReplaceAllString(c.value, "$1.$2.$3-$4")
}

// Scan implements database/sql/driver Scanner interface.
func (c *CPF) Scan(value interface{}) error {
	if value == nil {
		*c = CPF{
			value: "",
		}

		return common.ErrScanValueNil
	}

	if value, ok := value.(string); ok {
		cpf, err := NewCPF(value)
		if err != nil {
			return err
		}

		*c = cpf

		return nil
	}

	return common.ErrScanValueIsNotString
}

func (c *CPF) Value() string {
	return c.value
}
