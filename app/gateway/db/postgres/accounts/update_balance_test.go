package accounts

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/hash"
	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres/pgtest"
)

func TestRepository_UpdateBalance(t *testing.T) {
	t.Parallel()

	secretHash, err := hash.NewHash("the#$%PassWoRd")
	require.NoError(t, err)

	newCpf, err := cpf.NewCPF("88350057017")
	require.NoError(t, err)

	type args struct {
		acc entities.Account
	}

	tests := []struct {
		name      string
		args      func(t *testing.T) args
		wantErr   error
		beforeRun func(t *testing.T, db *pgxpool.Pool)
		check     func(t *testing.T, db *pgxpool.Pool)
	}{
		{
			name: "test case name here",
			args: func(t *testing.T) args {
				t.Helper()

				acc := entities.Account{
					ID:        "cdd3e9ed-b33b-4b18-b5a4-31a791969a30",
					Name:      "",
					Balance:   250_00,
					CPF:       cpf.CPF{},
					Secret:    hash.Hash{},
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				}

				return args{
					acc: acc,
				}
			},
			beforeRun: func(t *testing.T, db *pgxpool.Pool) {
				t.Helper()

				{
					acc := entities.Account{
						ID:        "cdd3e9ed-b33b-4b18-b5a4-31a791969a30",
						Name:      "Teste",
						Balance:   350_00,
						CPF:       newCpf,
						Secret:    secretHash,
						CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
						UpdatedAt: time.Time{},
					}

					err = pgtest.AccountsInsert(t, db, acc)
					require.NoError(t, err)
				}
			},
			check: func(t *testing.T, db *pgxpool.Pool) {
				t.Helper()

				{
					got, err := pgtest.GetAccount(t, db, "cdd3e9ed-b33b-4b18-b5a4-31a791969a30")
					require.NoError(t, err)

					newCpf, err := cpf.NewCPF("88350057017")
					require.NoError(t, err)

					expected := entities.Account{
						ID:        "cdd3e9ed-b33b-4b18-b5a4-31a791969a30",
						Name:      "Teste",
						Balance:   250_00,
						CPF:       newCpf,
						Secret:    secretHash,
						CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
						UpdatedAt: time.Date(2022, time.January, 15, 1, 0, 0, 0, time.Local),
					}

					got.UpdatedAt = time.Time{}
					expected.UpdatedAt = time.Time{}

					assert.Equal(t, expected, got)
				}
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			pgPool := pgtest.NewDB(t)
			repo := NewRepository(pgPool)

			args := tt.args(t)

			if tt.beforeRun != nil {
				tt.beforeRun(t, pgPool)
			}

			ctx := context.Background()

			tx, err := pgPool.Begin(ctx)
			require.NoError(t, err)

			err = repo.UpdateBalance(ctx, tx, args.acc.ID, args.acc.Balance)
			require.ErrorIs(t, err, tt.wantErr)

			err = tx.Commit(ctx)
			require.NoError(t, err)

			if tt.check != nil {
				tt.check(t, pgPool)
			}
		})
	}
}
