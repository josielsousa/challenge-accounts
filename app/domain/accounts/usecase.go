package accounts

import (
	"context"
	"time"

	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
)

// go :generate moq -fmt goimports -out usecase_mock.go . Usecase
type Usecase interface {
	Create(ctx context.Context, acc AccountInput) error
	GetAllAccounts(ctx context.Context) ([]AccountOutput, error)
	GetAccountBalance(ctx context.Context, accountID string) (int, error)
}

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
