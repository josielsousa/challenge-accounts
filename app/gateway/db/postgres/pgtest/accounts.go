package pgtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
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

	secret, err := acc.GetSecretHashed()
	require.NoError(t, err)

	row := db.QueryRow(
		context.Background(),
		query,
		acc.ID,
		acc.Name,
		acc.CPF.Value(),
		secret,
		acc.Balance,
		acc.CreatedAt,
		acc.UpdatedAt,
	)

	err = row.Scan(
		&acc.ID,
		&acc.CreatedAt,
		&acc.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("%s-> %s:%w", op, "on insert account", err)
	}

	return nil
}

func GetAccount(t *testing.T, db *pgxpool.Pool, id string) (accounts.Account, string, error) {
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

	var (
		numCPF string
		sec    string
		acc    accounts.Account
	)

	row := db.QueryRow(
		context.Background(),
		query,
		id,
	)

	err := row.Scan(
		&acc.ID,
		&acc.Name,
		&numCPF,
		&sec,
		&acc.Balance,
		&acc.CreatedAt,
		&acc.UpdatedAt,
	)
	if err != nil {
		return accounts.Account{}, "", fmt.Errorf("%s-> %s:%w", op, "on query account", err)
	}

	if len(numCPF) > 0 {
		accPF, err := cpf.NewCPF(numCPF)
		require.NoError(t, err)

		acc.CPF = accPF
	}

	acc.SetSecret(sec)

	return acc, sec, nil
}
