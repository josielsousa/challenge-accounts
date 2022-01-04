package accounts

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/hash"
	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres/pgtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepository_Update(t *testing.T) {
	t.Parallel()

	secPassword := "the#$%PassWoRd"

	type args struct {
		ctx context.Context
		acc accounts.Account
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
				secret, err := hash.NewHash(secPassword)
				require.NoError(t, err)

				newCpf, err := cpf.NewCPF("88350057017")
				require.NoError(t, err)

				return args{
					ctx: context.Background(),
					acc: accounts.Account{
						ID:        "cdd3e9ed-b33b-4b18-b5a4-31a791969a30",
						Name:      "Teste",
						Balance:   350_00,
						Secret:    secret,
						CPF:       newCpf,
						CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
						UpdatedAt: time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
					},
				}
			},
			beforeRun: func(t *testing.T, db *pgxpool.Pool) {
				{
					secret, err := hash.NewHash(secPassword)
					require.NoError(t, err)

					newCpf, err := cpf.NewCPF("88350057017")
					require.NoError(t, err)

					err = pgtest.AccountsInsert(db, t, accounts.Account{
						ID:        "cdd3e9ed-b33b-4b18-b5a4-31a791969a30",
						Name:      "Teste",
						Balance:   350_00,
						Secret:    secret,
						CPF:       newCpf,
						CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
					})
					require.NoError(t, err)
				}
			},
			check: func(t *testing.T, db *pgxpool.Pool) {
				{
					got, gotSecret, err := pgtest.GetAccount(db, t, "cdd3e9ed-b33b-4b18-b5a4-31a791969a30")
					require.NoError(t, err)

					secret, err := hash.NewHash(secPassword)
					require.NoError(t, err)

					newCpf, err := cpf.NewCPF("88350057017")
					require.NoError(t, err)

					expected := accounts.Account{
						ID:        "cdd3e9ed-b33b-4b18-b5a4-31a791969a30",
						Name:      "Teste",
						Balance:   350_00,
						Secret:    secret,
						CPF:       newCpf,
						CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
						UpdatedAt: time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
					}

					err = hash.CompareHashedAndSecret(gotSecret, secPassword)
					require.NoError(t, err)

					got.Secret = secret
					assert.Equal(t, expected, got)
				}
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt := tt // capture range variable

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			pgPool := pgtest.NewDB(t)
			r := NewRepository(pgPool)

			args := tt.args(t)

			if tt.beforeRun != nil {
				tt.beforeRun(t, pgPool)
			}

			err := r.Update(args.ctx, args.acc)
			assert.ErrorIs(t, err, tt.wantErr)

			if tt.check != nil {
				tt.check(t, pgPool)
			}
		})
	}
}
