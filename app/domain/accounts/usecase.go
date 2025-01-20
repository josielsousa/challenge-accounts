package accounts

import (
	"context"
	"time"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
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
	GetAll(ctx context.Context) ([]entities.Account, error)
	GetByID(ctx context.Context, id string) (entities.Account, error)
	GetByCPF(ctx context.Context, cpf string) (entities.Account, error)
	Insert(ctx context.Context, account entities.Account) error
}

type Usecase struct {
	R Repository
}

func NewUsecase(r Repository) *Usecase {
	return &Usecase{R: r}
}
