package transfers

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
	"github.com/josielsousa/challenge-accounts/app/domain/erring"
)

func (u Usecase) DoTransfer(
	ctx context.Context,
	input TransferInput,
) (TransferOutput, error) {
	const op = `transfers.DoTransfer`

	if input.Amount <= 0 {
		return TransferOutput{}, erring.ErrInvalidAmount
	}

	accOri, err := u.AR.GetByID(ctx, input.AccountOriginID)
	if err != nil {
		return TransferOutput{}, fmt.Errorf(
			"on get account origin: %w -> %w",
			err,
			erring.ErrAccountOriginNotFound,
		)
	}

	if accOri.Balance < input.Amount {
		return TransferOutput{}, erring.ErrInsufficientFunds
	}

	accDest, err := u.AR.GetByID(ctx, input.AccountDestinationID)
	if err != nil {
		return TransferOutput{}, fmt.Errorf(
			"on get account destination: %w -> %w",
			err,
			erring.ErrAccountDestinationNotFound,
		)
	}

	err = accOri.Withdraw(input.Amount)
	if err != nil {
		return TransferOutput{}, fmt.Errorf("%s-> %s: %w", op, "on withdraw", err)
	}

	err = accDest.Deposit(input.Amount)
	if err != nil {
		return TransferOutput{}, fmt.Errorf("%s-> %s: %w", op, "on deposit", err)
	}

	transferData := entities.TransferData{
		Transfer: entities.Transfer{
			ID:                   uuid.NewString(),
			Amount:               input.Amount,
			CreatedAt:            time.Now(),
			AccountOriginID:      accOri.ID,
			AccountDestinationID: accDest.ID,
		},
		AccountDestination: entities.AccountData{
			ID:      accDest.ID,
			Balance: accDest.Balance,
		},
		AccountOrigin: entities.AccountData{
			ID:      accOri.ID,
			Balance: accOri.Balance,
		},
	}

	err = u.R.Insert(ctx, transferData)
	if err != nil {
		return TransferOutput{}, fmt.Errorf("%s-> %s: %w", op, "on do transfer", err)
	}

	return TransferOutput{
		ID:                   transferData.ID,
		AccountOriginID:      transferData.AccountOriginID,
		AccountDestinationID: transferData.AccountDestinationID,
		Amount:               transferData.Amount,
		CreatedAt:            transferData.CreatedAt,
	}, nil
}
