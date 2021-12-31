package accounts

import (
	"errors"
	"time"

	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/hash"
)

var (
	ErrInvalidAmount     = errors.New("invalid amount")
	ErrInsufficientFunds = errors.New("insufficient amount")
)

// Account - Estrutura da entidade `account`
type Account struct {
	ID        string
	Name      string
	Balance   int
	Secret    hash.Hash
	CPF       cpf.CPF
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a *Account) Deposit(amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	if amount > a.Balance {
		return ErrInsufficientFunds
	}

	a.Balance -= amount
	return nil
}
