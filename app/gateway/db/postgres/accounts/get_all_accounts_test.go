package accounts

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/hash"
	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres/pgtest"
)

func TestRepository_GetAll(t *testing.T) {
	t.Parallel()

	secretHash, err := hash.NewHash("the#$%PassWoRd")
	require.NoError(t, err)

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
			beforeRun: func(t *testing.T, dbTest *pgxpool.Pool) {
				t.Helper()

				{
					for i := range fakeData {
						newCpf, err := cpf.NewCPF(fakeData[i].numCPF)
						require.NoError(t, err)

						acc := accounts.Account{
							ID:        fakeData[i].id,
							Balance:   350_00 * i,
							CPF:       newCpf,
							Secret:    secretHash,
							Name:      fmt.Sprintf("User name %d", i),
							CreatedAt: time.Date(2022, time.January, 4, i, 0, 0, 0, time.Local),
							UpdatedAt: time.Date(2022, time.January, 4, i, 0, 0, 0, time.Local),
						}

						err = pgtest.AccountsInsert(t, dbTest, acc)
						require.NoError(t, err)
					}
				}
			},
			check: func(t *testing.T, accs []accounts.Account) {
				t.Helper()

				{
					assert.Len(t, accs, len(fakeData))

					for i, got := range accs {
						newCpf, err := cpf.NewCPF(fakeData[i].numCPF)
						require.NoError(t, err)

						acc := accounts.Account{
							ID:        fakeData[i].id,
							Balance:   350_00 * i,
							CPF:       newCpf,
							Secret:    secretHash,
							Name:      fmt.Sprintf("User name %d", i),
							CreatedAt: time.Date(2022, time.January, 4, i, 0, 0, 0, time.Local),
							UpdatedAt: time.Date(2022, time.January, 4, i, 0, 0, 0, time.Local),
						}

						assert.Equal(t, acc, got)
					}
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			pgPool := pgtest.NewDB(t)

			tt.beforeRun(t, pgPool)

			r := NewRepository(pgPool)
			got, err := r.GetAll(context.Background())
			require.ErrorIs(t, err, tt.wantErr)

			tt.check(t, got)
		})
	}
}
