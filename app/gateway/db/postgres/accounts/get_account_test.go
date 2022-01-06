package accounts

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"

	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres/pgtest"
)

func TestRepository_GetAccountByCPF(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx    context.Context
		numCPF string
	}
	tests := []struct {
		name      string
		args      args
		check     func(t *testing.T, got accounts.Account)
		wantErr   error
		beforeRun func(t *testing.T, db *pgxpool.Pool)
	}{
		{
			name: "should get an account by cpf",
			args: args{
				ctx:    context.Background(),
				numCPF: "88350057017",
			},
			check: func(t *testing.T, got accounts.Account) {
				newCpf, err := cpf.NewCPF("88350057017")
				require.NoError(t, err)

				acc := accounts.Account{
					ID:        "cdd3e9ed-b33b-4b18-b5a4-31a791969a30",
					Name:      "Teste",
					Balance:   350_00,
					CPF:       newCpf,
					CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
					UpdatedAt: time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
				}

				hs, err := got.GetSecretHashed()
				require.NoError(t, err)

				err = bcrypt.CompareHashAndPassword([]byte(hs), []byte("123456"))
				require.NoError(t, err)

				acc.SetSecret("")
				got.SetSecret("")

				acc.SetHashedSecret("")
				got.SetHashedSecret("")

				assert.Equal(t, acc, got)
			},
			beforeRun: func(t *testing.T, db *pgxpool.Pool) {
				{
					newCpf, err := cpf.NewCPF("88350057017")
					require.NoError(t, err)

					acc := accounts.Account{
						ID:        "cdd3e9ed-b33b-4b18-b5a4-31a791969a30",
						Name:      "Teste",
						Balance:   350_00,
						CPF:       newCpf,
						CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
						UpdatedAt: time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
					}

					acc.SetSecret("123456")

					err = pgtest.AccountsInsert(t, db, acc)
					require.NoError(t, err)
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

			if tt.beforeRun != nil {
				tt.beforeRun(t, pgPool)
			}

			got, err := r.GetAccountByCPF(tt.args.ctx, tt.args.numCPF)
			assert.ErrorIs(t, err, tt.wantErr)

			tt.check(t, got)
		})
	}
}
