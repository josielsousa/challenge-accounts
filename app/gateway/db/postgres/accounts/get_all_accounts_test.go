package accounts

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres/pgtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestRepository_GetAll(t *testing.T) {
	t.Parallel()

	fakeData := []struct {
		numCPF string
		id     string
	}{
		{
			id:     "77f9a4a4-d81d-4db2-9d82-0f7d98464c1b",
			numCPF: "88350057017",
		},
		{
			id:     "fdebbfaf-3626-4247-af25-09565218441f",
			numCPF: "71970232030",
		},
		{
			id:     "4751e332-d7de-4df1-998d-417b41880076",
			numCPF: "21603280065",
		},
	}

	tests := []struct {
		name      string
		wantErr   error
		beforeRun func(t *testing.T, db *pgxpool.Pool)
		check     func(t *testing.T, accs []accounts.Account)
	}{
		{
			name:    "should got a accounts list with successfully",
			wantErr: nil,
			beforeRun: func(t *testing.T, db *pgxpool.Pool) {
				{
					for i := 0; i < len(fakeData); i++ {
						newCpf, err := cpf.NewCPF(fakeData[i].numCPF)
						require.NoError(t, err)

						acc := accounts.Account{
							ID:        fakeData[i].id,
							Balance:   350_00 * i,
							CPF:       newCpf,
							Name:      fmt.Sprintf("User name %d", i),
							CreatedAt: time.Date(2022, time.January, 4, i, 0, 0, 0, time.Local),
							UpdatedAt: time.Date(2022, time.January, 4, i, 0, 0, 0, time.Local),
						}

						acc.SetSecret(fmt.Sprintf("secret_string_%d", i))

						err = pgtest.AccountsInsert(t, db, acc)
						require.NoError(t, err)
					}
				}
			},
			check: func(t *testing.T, accs []accounts.Account) {
				{
					assert.Len(t, accs, len(fakeData))

					for i, got := range accs {
						hs, err := got.GetSecretHashed()
						require.NoError(t, err)

						err = bcrypt.CompareHashAndPassword([]byte(hs), []byte(fmt.Sprintf("secret_string_%d", i)))
						require.NoError(t, err)

						newCpf, err := cpf.NewCPF(fakeData[i].numCPF)
						require.NoError(t, err)

						acc := accounts.Account{
							ID:        fakeData[i].id,
							Balance:   350_00 * i,
							CPF:       newCpf,
							Name:      fmt.Sprintf("User name %d", i),
							CreatedAt: time.Date(2022, time.January, 4, i, 0, 0, 0, time.Local),
							UpdatedAt: time.Date(2022, time.January, 4, i, 0, 0, 0, time.Local),
						}

						acc.SetSecret("")
						got.SetSecret("")

						acc.SetHashedSecret("")
						got.SetHashedSecret("")

						assert.Equal(t, acc, got)
					}
				}
			},
		},
	}
	for _, tt := range tests {
		tt := tt // capture range variable

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			pgPool := pgtest.NewDB(t)

			tt.beforeRun(t, pgPool)

			r := NewRepository(pgPool)
			got, err := r.GetAll(context.Background())
			assert.ErrorIs(t, err, tt.wantErr)

			tt.check(t, got)
		})
	}
}
