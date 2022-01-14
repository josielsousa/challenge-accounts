package transfers

import (
	"context"
	"fmt"

	"github.com/josielsousa/challenge-accounts/app/domain/entities/transfers"
)

const (
	queryGetAllTransfersByAccOriginID = `
	SELECT 
		id,
		amount,
		account_origin_id,
		account_destination_id,
		created_at,
		updated_at
	FROM transfers 
	WHERE account_origin_id = $1 
	`
)

func (r *Repository) ListTransfers(ctx context.Context, accOriginID string) ([]transfers.Transfer, error) {
	const op = `Repository.Transfers.ListTransfers`

	trfs := make([]transfers.Transfer, 0)

	rows, err := r.db.Query(context.Background(), queryGetAllTransfersByAccOriginID, accOriginID)
	if err != nil {
		return nil, fmt.Errorf("%s -> %s: %w", op, "on get all transfers by account origin", err)
	}
	defer rows.Close()

	for rows.Next() {
		var trf transfers.Transfer
		rows.Scan(
			&trf.ID,
			&trf.Amount,
			&trf.AccountOriginID,
			&trf.AccountDestinationID,
			&trf.CreatedAt,
			&trf.UpdatedAt,
		)

		trfs = append(trfs, trf)
	}

	return trfs, nil
}
