package transfers

import (
	"context"
	"fmt"
)

func (u Usecase) ListTransfersAccount(
	ctx context.Context,
	accOriginID string,
) ([]TransferOutput, error) {
	const op = `transfers.ListTransfersAccount`

	allTrfs, err := u.R.ListTransfers(ctx, accOriginID)
	if err != nil {
		return nil, fmt.Errorf("%s -> on list transfers: %w", op, err)
	}

	out := make([]TransferOutput, len(allTrfs))

	for i, transfer := range allTrfs {
		out[i] = TransferOutput{
			ID:                   transfer.ID,
			Amount:               transfer.Amount,
			AccountOriginID:      transfer.AccountOriginID,
			AccountDestinationID: transfer.AccountDestinationID,
			CreatedAt:            transfer.CreatedAt,
		}
	}

	return out, nil
}
