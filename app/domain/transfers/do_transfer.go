package transfers

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	accE "github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	trfE "github.com/josielsousa/challenge-accounts/app/domain/entities/transfers"
)

func (u Usecase) DoTransfer(ctx context.Context, input TransferInput) error {
	const op = `transfers.DoTransfer`

	if input.Amount <= 0 {
		return accE.ErrInvalidAmount
	}

	accOri, err := u.AR.GetByID(ctx, input.AccountOriginID)
	if err != nil {
		return fmt.Errorf("on get account origin: %w -> %w", err, accE.ErrAccountOriginNotFound)
	}

	if accOri.Balance < input.Amount {
		return accE.ErrInsufficientFunds
	}

	accDest, err := u.AR.GetByID(ctx, input.AccountDestinationID)
	if err != nil {
		return fmt.Errorf("on get account destination: %w -> %w", err, accE.ErrAccountDestinationNotFound)
	}

	err = accOri.Withdraw(input.Amount)
	if err != nil {
		return fmt.Errorf("%s-> %s: %w", op, "on withdraw", err)
	}

	err = accDest.Deposit(input.Amount)
	if err != nil {
		return fmt.Errorf("%s-> %s: %w", op, "on deposit", err)
	}

	err = u.R.Insert(ctx, trfE.TransferData{
		Transfer: trfE.Transfer{
			ID:                   uuid.NewString(),
			Amount:               input.Amount,
			CreatedAt:            time.Now(),
			AccountOriginID:      accOri.ID,
			AccountDestinationID: accDest.ID,
		},
		AccountDestination: trfE.AccountData{
			ID:      accDest.ID,
			Balance: accDest.Balance,
		},
		AccountOrigin: trfE.AccountData{
			ID:      accOri.ID,
			Balance: accOri.Balance,
		},
	})
	if err != nil {
		return fmt.Errorf("%s-> %s: %w", op, "on do transfer", err)
	}

	return nil
}
