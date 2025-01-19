package transfers

import (
	"context"
	"time"

	accE "github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	trfE "github.com/josielsousa/challenge-accounts/app/domain/entities/transfers"
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
	Insert(ctx context.Context, transfer trfE.TransferData) error
	ListTransfers(ctx context.Context, accOriginID string) ([]trfE.Transfer, error)
}

// AccountRepository - Interface que define as assinaturas para o repository de
// accounts.
type AccountRepository interface {
	GetByID(ctx context.Context, id string) (accE.Account, error)
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
