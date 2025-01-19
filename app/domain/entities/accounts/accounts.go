package accounts

import (
	"time"

	"github.com/josielsousa/challenge-accounts/app/domain/erring"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/hash"
)

// Account - Estrutura da entidade `account`.
type Account struct {
	ID        string
	Name      string
	Balance   int
	CPF       cpf.CPF
	Secret    hash.Hash
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a *Account) Deposit(amount int) error {
	if amount <= 0 {
		return erring.ErrInvalidAmount
	}

	a.Balance += amount

	return nil
}

func (a *Account) Withdraw(amount int) error {
	if amount <= 0 {
		return erring.ErrInvalidAmount
	}

	if amount > a.Balance {
		return erring.ErrInsufficientFunds
	}

	a.Balance -= amount

	return nil
}
