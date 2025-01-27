package accounts

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
	"github.com/josielsousa/challenge-accounts/app/domain/erring"
)

const (
	queryByCPF = `
		SELECT
			id,
			name,
			cpf,
			secret,
			balance,
			created_at,
			updated_at
		FROM accounts
		WHERE cpf = $1
	`

	queryByID = `
		SELECT
			id,
			name,
			cpf,
			secret,
			balance,
			created_at,
			updated_at
		FROM accounts
		WHERE id = $1
	`
)

func (r *Repository) GetByCPF(
	ctx context.Context, numCPF string,
) (entities.Account, error) {
	const op = `Repository.Accounts.GetByCPF`

	acc, err := r.getAccount(ctx, numCPF, queryByCPF)
	if err != nil {
		return entities.Account{}, fmt.Errorf(
			"%s-> on query by CPF: %w", op, err)
	}

	return acc, nil
}

func (r *Repository) GetByID(
	ctx context.Context, id string,
) (entities.Account, error) {
	const operation = `Repository.Accounts.GetAccountByCPF`

	acc, err := r.getAccount(ctx, id, queryByID)
	if err != nil {
		return entities.Account{}, fmt.Errorf(
			"%s-> on query by id: %w", operation, err)
	}

	return acc, nil
}

func (r *Repository) getAccount(
	ctx context.Context, param, query string,
) (entities.Account, error) {
	const op = `Repository.Accounts.getAccount`

	row := r.db.QueryRow(
		ctx,
		query,
		param,
	)

	var acc entities.Account

	err := row.Scan(
		&acc.ID,
		&acc.Name,
		&acc.CPF,
		&acc.Secret,
		&acc.Balance,
		&acc.CreatedAt,
		&acc.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entities.Account{}, fmt.Errorf(
				"%s -> on get account no rows returned: %w",
				op, erring.ErrAccountNotFound,
			)
		}

		return entities.Account{}, fmt.Errorf(
			"%s-> on get account: %w", op, err,
		)
	}

	return acc, nil
}
