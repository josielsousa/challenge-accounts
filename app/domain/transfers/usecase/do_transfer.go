package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/entities/transfers"
	trfUC "github.com/josielsousa/challenge-accounts/app/domain/transfers"
)

func (t Transfer) DoTransfer(ctx context.Context, input trfUC.TransferInput) error {
	const op = `Usecase.Transfer.DoTransfer`

	if input.Amount <= 0 {
		return accounts.ErrInvalidAmount
	}

	accOri, err := t.accRepo.GetByID(ctx, input.AccountOriginID)
	if err != nil {
		return accounts.ErrAccountOriginNotFound
	}

	if accOri.Balance < input.Amount {
		return accounts.ErrInsufficientFunds
	}

	accDest, err := t.accRepo.GetByID(ctx, input.AccountDestinationID)
	if err != nil {
		return accounts.ErrAccountDestinationNotFound
	}

	err = t.repo.Insert(ctx, transfers.TransferData{
		Transfer: transfers.Transfer{
			ID:                   uuid.NewString(),
			Amount:               input.Amount,
			CreatedAt:            time.Now().Local(),
			AccountOriginID:      accOri.ID,
			AccountDestinationID: accDest.ID,
		},
		AccountDestination: accDest,
		AccountOrigin:      accOri,
	})
	if err != nil {
		return fmt.Errorf("%s-> %s: %w", op, "on do transfer", err)
	}

	return nil
}
