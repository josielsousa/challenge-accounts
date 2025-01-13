package transfers

import (
	"time"

	accE "github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	trfE "github.com/josielsousa/challenge-accounts/app/domain/entities/transfers"
)

type Usecase struct {
	AR accE.Repository
	R  trfE.Repository
}

func NewUsecase(
	trfRepo trfE.Repository,
	accRepo accE.Repository,
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
