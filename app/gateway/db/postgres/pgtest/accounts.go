package pgtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
)

func AccountsInsert(t *testing.T, db *pgxpool.Pool, acc accounts.Account) error {
	t.Helper()

	const op = `Pgtest.AccountsInsert`
	if len(acc.ID) <= 0 {
		acc.ID = uuid.NewString()
	}

	if acc.CreatedAt.IsZero() {
		acc.CreatedAt = time.Now().In(time.Local)
	}

	if acc.UpdatedAt.IsZero() {
		acc.UpdatedAt = time.Now().In(time.Local)
	}

	query := `
        INSERT INTO accounts(
            id,
            name,
            cpf,
            secret,
            balance,
            created_at,
			updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id, created_at, updated_at
    `

	row := db.QueryRow(
		context.Background(),
		query,
		acc.ID,
		acc.Name,
		acc.CPF.Value(),
		acc.Secret.Value(),
		acc.Balance,
		acc.CreatedAt,
		acc.UpdatedAt,
	)

	err := row.Scan(
		&acc.ID,
		&acc.CreatedAt,
		&acc.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("%s-> %s:%w", op, "on insert account", err)
	}

	return nil
}

func GetAccount(t *testing.T, db *pgxpool.Pool, id string) (accounts.Account, error) {
	t.Helper()

	const op = `Pgtest.GetAccount`

	query := `
        SELECT 
            id,
            name,
            cpf,
            secret,
            balance,
            created_at,
			updated_at
		FROM  accounts 
		WHERE id = $1 
    `

	var acc accounts.Account

	row := db.QueryRow(
		context.Background(),
		query,
		id,
	)

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
		return accounts.Account{}, fmt.Errorf("%s-> %s:%w", op, "on query account", err)
	}

	return acc, nil
}