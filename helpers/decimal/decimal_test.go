package decimal_test

import (
	"testing"

	"github.com/josielsousa/challenge-accounts/helpers/decimal"
)

const (
	ErrorOnAdd = "O resultado %.2f foi diferente do esperado %.2f"
)

func TestDecimalAdd(t *testing.T) {
	t.Run("Teste decimal soma", func(t *testing.T) {
		amount := float64(.2)
		balance := float64(.1)
		expected := float64(.3)

		result := decimal.Add(balance, amount)
		if result > expected || result < expected {
			t.Errorf(ErrorOnAdd, result, expected)
		}
	})
}

func TestDecimalSub(t *testing.T) {
	t.Run("Teste decimal soma", func(t *testing.T) {
		amount := float64(.1)
		balance := float64(.4)
		expected := float64(.3)

		result := decimal.Sub(amount, balance)
		if result > expected || result < expected {
			t.Errorf(ErrorOnAdd, result, expected)
		}
	})
}
