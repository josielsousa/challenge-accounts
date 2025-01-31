package transfers

import (
	"context"
	"fmt"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
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

func (r *Repository) ListTransfers(
	ctx context.Context, accOriginID string,
) ([]entities.Transfer, error) {
	const op = `Repository.Transfers.ListTransfers`

	trfs := make([]entities.Transfer, 0)

	rows, err := r.db.Query(ctx, queryGetAllTransfersByAccOriginID, accOriginID)
	if err != nil {
		return nil, fmt.Errorf(
			"%s -> on get all transfers by account origin: %w", op, err,
		)
	}
	defer rows.Close()

	for rows.Next() {
		var trf entities.Transfer

		err := rows.Scan(
			&trf.ID,
			&trf.Amount,
			&trf.AccountOriginID,
			&trf.AccountDestinationID,
			&trf.CreatedAt,
			&trf.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s -> on scan row: %w", op, err)
		}

		trfs = append(trfs, trf)
	}

	return trfs, nil
}
