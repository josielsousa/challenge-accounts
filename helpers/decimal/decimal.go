package decimal

import (
	"github.com/shopspring/decimal"
)

// Add - Adiciona o `amount` informado, no balance atual.
func Add(balance, amount float64) float64 {
	amountDecimal := decimal.NewFromFloat(amount)
	balanceDecimal := decimal.NewFromFloat(balance)

	subtotal := balanceDecimal.Add(amountDecimal)
	total, _ := subtotal.Float64()
	return total
}

// Sub - Subtrai o `amount` informado, no balance atual.
func Sub(balance, amount float64) float64 {
	amountDecimal := decimal.NewFromFloat(amount)
	balanceDecimal := decimal.NewFromFloat(balance)

	subtotal := balanceDecimal.Sub(amountDecimal)
	total, _ := subtotal.Float64()
	return total
}
