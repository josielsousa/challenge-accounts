package usecase

import (
	"context"
	"fmt"

	trfUC "github.com/josielsousa/challenge-accounts/app/domain/transfers"
)

func (t Transfer) ListTransfersAccount(ctx context.Context, accOriginID string) ([]trfUC.TransferOutput, error) {
	const op = `transfers.ListTransfersAccount`

	transfers, err := t.repo.ListTransfers(ctx, accOriginID)
	if err != nil {
		return nil, fmt.Errorf("%s -> %s: %w", op, "on list transfers", err)
	}

	out := make([]trfUC.TransferOutput, 0, len(transfers))

	for i, transfer := range transfers {
		out[i] = trfUC.TransferOutput{
			ID:                   transfer.ID,
			Amount:               transfer.Amount,
			AccountOriginID:      transfer.AccountOriginID,
			AccountDestinationID: transfer.AccountDestinationID,
			CreatedAt:            transfer.CreatedAt,
		}
	}

	return out, nil
}
