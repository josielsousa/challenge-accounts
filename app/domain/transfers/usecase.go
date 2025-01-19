package transfers

import (
	"context"
	"time"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
)

type Usecase struct {
	AR AccountRepository
	R  Repository
}

// Repository - Interface que define as assinaturas para o repository de
// transfers.
//
//go:generate moq -rm -out usecase_mock.go . Repository AccountRepository
type Repository interface {
	Insert(ctx context.Context, transfer entities.TransferData) error
	ListTransfers(ctx context.Context, accOriginID string) ([]entities.Transfer, error)
}

// AccountRepository - Interface que define as assinaturas para o repository de
// accounts.
type AccountRepository interface {
	// TODO: add update balance inside a transactioner
	GetByID(ctx context.Context, id string) (entities.Account, error)
}

func NewUsecase(
	trfRepo Repository,
	accRepo AccountRepository,
) *Usecase {
	return &Usecase{
		R:  trfRepo,
		AR: accRepo,
	}
}

type TransferInput struct {
	AccountOriginID      string
	AccountDestinationID string
	Amount               int
}

type TransferOutput struct {
	ID                   string
	AccountOriginID      string
	AccountDestinationID string
	Amount               int
	CreatedAt            time.Time
}
