package accounts

import (
	"context"
	"time"

	accE "github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
)

type AccountInput struct {
	Name    string
	Balance int
	CPF     string
	Secret  string
}

type AccountOutput struct {
	ID        string
	Name      string
	Balance   int
	CPF       cpf.CPF
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Repository - Interface que define as assinaturas para o repository de accounts.
//
//go:generate moq -fmt goimports -out usecase_mock.go . Repository
type Repository interface {
	GetAll(ctx context.Context) ([]accE.Account, error)
	GetByID(ctx context.Context, id string) (accE.Account, error)
	GetByCPF(ctx context.Context, cpf string) (accE.Account, error)
	Insert(ctx context.Context, account accE.Account) error
	UpdateBalance(ctx context.Context, accountID string, balance int) error
}

type Account struct {
	R Repository
}

func NewUsecase(r Repository) *Account {
	return &Account{R: r}
}
