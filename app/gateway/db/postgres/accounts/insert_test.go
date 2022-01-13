package accounts

import (
	"context"
	"testing"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
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
		ctx context.Context
		acc accounts.Account
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
				acc := accounts.Account{
					Name:    "Teste",
					Balance: 350_00,
					CPF:     newCpf,
					Secret:  secretHash,
				}

				return args{
					acc: acc,
					ctx: context.Background(),
				}
			},
			checkErr: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "should return an postgres error when save new account, because column name is invalid",
			args: func(t *testing.T) args {
				acc := accounts.Account{
					Name:    "Teste",
					Balance: 350_00,
					CPF:     newCpf,
					Secret:  secretHash,
				}

				return args{
					acc: acc,
					ctx: context.Background(),
				}
			},
			beforeRun: func(t *testing.T, db *pgxpool.Pool) {
				{
					_, err := db.Exec(context.Background(), "ALTER TABLE accounts RENAME COLUMN name TO username")
					require.NoError(t, err)
				}
			},
			checkErr: func(t *testing.T, err error) {
				var pgErr *pgconn.PgError
				assert.ErrorAs(t, err, &pgErr)
				assert.Equal(t, `column "name" of relation "accounts" does not exist`, pgErr.Message)
			},
		},
		{
			name: "should return an error when save a duplicate account",
			args: func(t *testing.T) args {
				acc := accounts.Account{
					Name:    "Teste",
					Balance: 350_00,
					CPF:     newCpf,
					Secret:  secretHash,
				}

				return args{
					acc: acc,
					ctx: context.Background(),
				}
			},
			beforeRun: func(t *testing.T, db *pgxpool.Pool) {
				{
					acc := accounts.Account{
						Name:    "Teste",
						Balance: 350_00,
						CPF:     newCpf,
						Secret:  secretHash,
					}

					err = pgtest.AccountsInsert(t, db, acc)
					require.NoError(t, err)
				}
			},
			checkErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, accounts.ErrAccountAlreadyExists)
			},
		},
	}
	for _, tt := range tests {
		tt := tt // capture range variable

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			pgPool := pgtest.NewDB(t)
			r := NewRepository(pgPool)

			if tt.beforeRun != nil {
				tt.beforeRun(t, pgPool)
			}

			args := tt.args(t)

			err := r.Insert(args.ctx, args.acc)
			tt.checkErr(t, err)
		})
	}
}
