package accounts

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
	"github.com/josielsousa/challenge-accounts/app/domain/erring"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/hash"
	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres/pgtest"
)

func TestRepository_GetAccountByCPF(t *testing.T) {
	t.Parallel()

	secretHash, err := hash.NewHash("the#$%PassWoRd")
	require.NoError(t, err)

	newCpf, err := cpf.NewCPF("88350057017")
	require.NoError(t, err)

	type args struct {
		numCPF string
	}

	tests := []struct {
		name      string
		args      args
		beforeRun func(t *testing.T, db *pgxpool.Pool)
		check     func(t *testing.T, got entities.Account)
		wantErr   error
	}{
		{
			name: "should get an account by cpf",
			args: args{
				numCPF: "88350057017",
			},
			check: func(t *testing.T, got entities.Account) {
				t.Helper()

				acc := entities.Account{
					ID:        "cdd3e9ed-b33b-4b18-b5a4-31a791969a30",
					Name:      "Teste",
					Balance:   350_00,
					CPF:       newCpf,
					Secret:    secretHash,
					CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
					UpdatedAt: time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
				}

				assert.Equal(t, acc, got)
			},
			beforeRun: func(t *testing.T, dbTest *pgxpool.Pool) {
				t.Helper()

				{
					acc := entities.Account{
						ID:        "cdd3e9ed-b33b-4b18-b5a4-31a791969a30",
						Name:      "Teste",
						Balance:   350_00,
						CPF:       newCpf,
						Secret:    secretHash,
						CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
						UpdatedAt: time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
					}

					err = pgtest.AccountsInsert(t, dbTest, acc)
					require.NoError(t, err)
				}
			},
			wantErr: nil,
		},
		{
			name: "should return an error when account not found",
			args: args{
				numCPF: "88350057013",
			},
			check: nil,
			beforeRun: func(t *testing.T, dbTest *pgxpool.Pool) {
				t.Helper()
				{
					acc := entities.Account{ID: "cdd3e9ed-b33b-4b18-b5a4-31a791969a30", Name: "Teste", Balance: 350_00, CPF: newCpf, Secret: secretHash, CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local), UpdatedAt: time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local)}
					err = pgtest.AccountsInsert(t, dbTest, acc)
					require.NoError(t, err)
				}
			},
			wantErr: erring.ErrAccountNotFound,
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

			got, err := repo.GetByCPF(context.Background(), tt.args.numCPF)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.check != nil {
				tt.check(t, got)
			}
		})
	}
}

func TestRepository_GetAccountByID(t *testing.T) {
	t.Parallel()

	secretHash, err := hash.NewHash("the#$%PassWoRd")
	require.NoError(t, err)

	newCpf, err := cpf.NewCPF("88350057017")
	require.NoError(t, err)

	type args struct {
		id string
	}

	tests := []struct {
		name      string
		args      args
		beforeRun func(t *testing.T, db *pgxpool.Pool)
		check     func(t *testing.T, got entities.Account)
		wantErr   error
	}{
		{
			name: "should get an account by id",
			args: args{
				id: "cdd3e9ed-b33b-4b18-b5a4-31a791969a30",
			},
			check: func(t *testing.T, got entities.Account) {
				t.Helper()

				acc := entities.Account{
					ID:        "cdd3e9ed-b33b-4b18-b5a4-31a791969a30",
					Name:      "Teste",
					Balance:   350_00,
					CPF:       newCpf,
					Secret:    secretHash,
					CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
					UpdatedAt: time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
				}

				assert.Equal(t, acc, got)
			},
			beforeRun: func(t *testing.T, dbTest *pgxpool.Pool) {
				t.Helper()

				{
					acc := entities.Account{
						ID:        "cdd3e9ed-b33b-4b18-b5a4-31a791969a30",
						Name:      "Teste",
						Balance:   350_00,
						CPF:       newCpf,
						Secret:    secretHash,
						CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
						UpdatedAt: time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
					}

					err = pgtest.AccountsInsert(t, dbTest, acc)
					require.NoError(t, err)
				}
			},
			wantErr: nil,
		},
		{
			name: "should return an error when account not found",
			args: args{
				id: "d079de7d-b3d2-47fa-b1d6-b5c7d7cf5389",
			},
			beforeRun: func(t *testing.T, dbTest *pgxpool.Pool) {
				t.Helper()
				{
					acc := entities.Account{ID: "cdd3e9ed-b33b-4b18-b5a4-31a791969a30", Name: "Teste", Balance: 350_00, CPF: newCpf, Secret: secretHash, CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local), UpdatedAt: time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local)}
					err = pgtest.AccountsInsert(t, dbTest, acc)
					require.NoError(t, err)
				}
			},
			check:   nil,
			wantErr: erring.ErrAccountNotFound,
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

			got, err := repo.GetByID(context.Background(), tt.args.id)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.check != nil {
				tt.check(t, got)
			}
		})
	}
}
