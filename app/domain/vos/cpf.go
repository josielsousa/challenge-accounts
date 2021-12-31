package vos

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const (
	firstOneDigitPosition   = 10
	regexPatternOnlyNumbers = `[^0-9]+`
	regexPatternCPFMasked   = `^(\d{3})(\d{3})(\d{3})(\d{2})$`
)

var ErrInvalid = errors.New("invalid cpf number")

type CPF struct {
	value string
}

func (c CPF) String() string {
	return c.value
}

func NewCPF(cpf string) (CPF, error) {
	c := CPF{
		value: removeSpecialChars(cpf),
	}

	if ok := c.isValid(); !ok {
		return CPF{}, ErrInvalid
	}

	return c, nil
}

// removeSpecialChars - Remove os caracteres não númericos.
func removeSpecialChars(value string) string {
	exp, err := regexp.Compile(regexPatternOnlyNumbers)
	if err != nil {
		return ""
	}

	return exp.ReplaceAllString(value, "")
}

// allEquals - Verifica se todos os dígitos do CPF são iguais, e.g.: 111.111.111-11
func (c CPF) allEquals() bool {
	base := c.value[0]
	for i := 1; i < len(c.value); i++ {
		if base != c.value[i] {
			return false
		}
	}

	return true
}

// calculateDigit - Calcula o digito verificador do documento informado conforme seu tipo.
//	position - representa o peso para a regra de cálculo do digito verificador.
// CPF pesos: 10, 9, 8, 7, 6, 5, 4, 3, 2
func (c CPF) calculateDigit(position int) string {
	var sum int
	data := c.value[:position-1]

	for _, digit := range data {
		sum += int(digit-'0') * position
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

// IsValid - Verifica se o CPF informado é válido, calculando os dígitos verificadores.
func (c CPF) isValid() bool {
	if len(c.value) != 11 || c.allEquals() {
		return false
	}

	doc := c.value[:9]
	firstOne := c.calculateDigit(firstOneDigitPosition)
	secondOne := c.calculateDigit(firstOneDigitPosition + 1)

	doc = fmt.Sprintf("%s%s%s", doc, firstOne, secondOne)
	return c.value == doc
}

// mask retorna o CPF informado no padrão (000.000.000-00).
func (c CPF) Mask() string {
	exp, err := regexp.Compile(regexPatternCPFMasked)
	if err != nil {
		return ""
	}

	return exp.ReplaceAllString(c.value, "$1.$2.$3-$4")
}
