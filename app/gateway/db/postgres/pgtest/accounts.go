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

func AccountsInsert(db *pgxpool.Pool, t *testing.T, acc accounts.Account) error {
	t.Helper()

	const op = `Pgtest.AccountsInsert`
	if len(acc.ID) <= 0 {
		acc.ID = uuid.NewString()
	}

	if acc.CreatedAt.IsZero() {
		acc.CreatedAt = time.Now().In(time.Local)
	}

	query := `
        INSERT INTO accounts(
            id,
            name,
            cpf,
            secret,
            balance,
            created_at
        ) VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, created_at, updated_at
    `

	row := db.QueryRow(
		context.Background(),
		query,
		acc.ID,
		acc.Name,
		acc.CPF.String(),
		acc.Secret.String(),
		acc.Balance,
		acc.CreatedAt,
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

func GetAccount(db *pgxpool.Pool, t *testing.T, id string) (accounts.Account, string, error) {
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
	var numCpf, sec string

	row := db.QueryRow(
		context.Background(),
		query,
		id,
	)

	err := row.Scan(
		&acc.ID,
		&acc.Name,
		&numCpf,
		&sec,
		&acc.Balance,
		&acc.CreatedAt,
		&acc.UpdatedAt,
	)
	if err != nil {
		return accounts.Account{}, "", fmt.Errorf("%s-> %s:%w", op, "on query account", err)
	}

	if len(numCpf) > 0 {
		accPF, err := cpf.NewCPF(numCpf)
		require.NoError(t, err)

		acc.CPF = accPF
	}

	return acc, sec, nil
}
