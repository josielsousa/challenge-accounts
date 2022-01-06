package accounts

import (
	"errors"
	"time"

	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/hash"
)

var (
	ErrInvalidAmount        = errors.New("invalid amount")
	ErrAccountAlreadyExists = errors.New("account already exists")
	ErrInsufficientFunds    = errors.New("insufficient amount")
)

// Account - Estrutura da entidade `account`
type Account struct {
	secret       string
	hashedSecret string
	ID           string
	Name         string
	Balance      int
	CPF          cpf.CPF
	CreatedAt    time.Time
	UpdatedAt    time.Time
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

func (a *Account) SetSecret(secret string) {
	a.secret = secret
}

func (a *Account) SetHashedSecret(hashedSecret string) {
	a.hashedSecret = hashedSecret
}

func (a *Account) GetSecretHashed() (string, error) {
	if len(a.hashedSecret) <= 0 {
		hs, err := hash.GenHash(a.secret)
		if err != nil {
			return "", err
		}

		a.hashedSecret = hs
	}

	return a.hashedSecret, nil
}
