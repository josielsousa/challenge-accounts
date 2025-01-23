package accounts

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
	"github.com/josielsousa/challenge-accounts/app/domain/erring"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/hash"
	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres/pgtest"
)

func TestRepository_Insert(t *testing.T) {
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
		checkErr  func(t *testing.T, err error)
		beforeRun func(t *testing.T, db *pgxpool.Pool)
	}{
		{
			name: "should save account with successfully",
			args: func(t *testing.T) args {
				t.Helper()

				acc := entities.Account{
					ID:        "",
					Name:      "Teste",
					Balance:   350_00,
					CPF:       newCpf,
					Secret:    secretHash,
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				}

				return args{
					acc: acc,
				}
			},
			beforeRun: nil,
			checkErr: func(t *testing.T, err error) {
				t.Helper()

				require.NoError(t, err)
			},
		},
		{
			name: "should return an postgres error when save new account, because column name is invalid",
			args: func(t *testing.T) args {
				t.Helper()

				acc := entities.Account{
					ID:        "",
					Name:      "Teste",
					Balance:   350_00,
					CPF:       newCpf,
					Secret:    secretHash,
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				}

				return args{
					acc: acc,
				}
			},
			beforeRun: func(t *testing.T, db *pgxpool.Pool) {
				t.Helper()

				_, err := db.Exec(context.Background(), "ALTER TABLE accounts RENAME COLUMN name TO username")
				require.NoError(t, err)
			},
			checkErr: func(t *testing.T, err error) {
				t.Helper()

				var pgErr *pgconn.PgError
				require.ErrorAs(t, err, &pgErr)
				assert.Equal(t, `column "name" of relation "accounts" does not exist`, pgErr.Message)
			},
		},
		{
			name: "should return an error when save a duplicate account",
			args: func(t *testing.T) args {
				t.Helper()

				acc := entities.Account{
					ID:        "",
					Name:      "Teste",
					Balance:   350_00,
					CPF:       newCpf,
					Secret:    secretHash,
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				}

				return args{
					acc: acc,
				}
			},
			beforeRun: func(t *testing.T, dbTest *pgxpool.Pool) {
				t.Helper()

				{
					acc := entities.Account{
						ID:        "",
						Name:      "Teste",
						Balance:   350_00,
						CPF:       newCpf,
						Secret:    secretHash,
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}

					err = pgtest.AccountsInsert(t, dbTest, acc)
					require.NoError(t, err)
				}
			},
			checkErr: func(t *testing.T, err error) {
				t.Helper()

				require.ErrorIs(t, err, erring.ErrAccountAlreadyExists)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			pgPool := pgtest.NewDB(t)
			repo := NewRepository(pgPool)

			if tt.beforeRun != nil {
				tt.beforeRun(t, pgPool)
			}

			args := tt.args(t)

			err := repo.Insert(context.Background(), args.acc)
			tt.checkErr(t, err)
		})
	}
}
