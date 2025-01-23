package transfers

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/hash"
	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres/pgtest"
)

func TestRepository_Insert(t *testing.T) {
	t.Parallel()

	trfID := uuid.NewString()
	accOriginID := uuid.NewString()
	accDestinationID := uuid.NewString()

	secretHash, err := hash.NewHash("the#$%PassWoRd")
	require.NoError(t, err)

	newCpf, err := cpf.NewCPF("88350057017")
	require.NoError(t, err)

	newCpf02, err := cpf.NewCPF("71970232030")
	require.NoError(t, err)

	type args struct {
		trf entities.TransferData
	}

	tests := []struct {
		name      string
		args      args
		checkErr  func(t *testing.T, err error)
		beforeRun func(t *testing.T, db *pgxpool.Pool)
		check     func(t *testing.T, pgPool *pgxpool.Pool)
	}{
		{
			name: "should insert a transfer with successfully",
			args: args{
				trf: entities.TransferData{
					Transfer: entities.Transfer{
						ID:                   trfID,
						AccountOriginID:      accOriginID,
						AccountDestinationID: accDestinationID,
						Amount:               50,
						CreatedAt:            time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
						UpdatedAt:            time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
					},
					AccountOrigin: entities.AccountData{
						ID:      accOriginID,
						Balance: 350_00,
					},
					AccountDestination: entities.AccountData{
						ID:      accDestinationID,
						Balance: 50_00,
					},
				},
			},
			beforeRun: func(t *testing.T, db *pgxpool.Pool) {
				t.Helper()

				{
					accs := []entities.Account{
						{
							ID:        accOriginID,
							Name:      "Teste 01",
							Balance:   350_00,
							CPF:       newCpf,
							Secret:    secretHash,
							CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
							UpdatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
						},
						{
							ID:        accDestinationID,
							Name:      "Teste 02",
							Balance:   50_00,
							CPF:       newCpf02,
							Secret:    secretHash,
							CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
							UpdatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
						},
					}

					for _, acc := range accs {
						err = pgtest.AccountsInsert(t, db, acc)
						require.NoError(t, err)
					}
				}
			},
			check: func(t *testing.T, pgPool *pgxpool.Pool) {
				t.Helper()

				{
					got, err := pgtest.GetTransfer(t, pgPool, trfID)
					require.NoError(t, err)

					expected := entities.Transfer{
						ID:                   trfID,
						AccountOriginID:      accOriginID,
						AccountDestinationID: accDestinationID,
						Amount:               50,
						CreatedAt:            time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
						UpdatedAt:            time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
					}

					assert.Equal(t, expected, got)
				}
				{
					got01, err := pgtest.GetAccount(t, pgPool, accOriginID)
					require.NoError(t, err)

					expected01 := entities.Account{
						ID:        accOriginID,
						Name:      "Teste 01",
						Balance:   350_00,
						CPF:       newCpf,
						Secret:    secretHash,
						CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
						UpdatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
					}

					got01.UpdatedAt = time.Time{}
					expected01.UpdatedAt = time.Time{}

					assert.Equal(t, expected01, got01)
				}
				{
					got02, err := pgtest.GetAccount(t, pgPool, accDestinationID)
					require.NoError(t, err)

					expected02 := entities.Account{
						ID:        accDestinationID,
						Name:      "Teste 02",
						Balance:   50_00,
						CPF:       newCpf02,
						Secret:    secretHash,
						CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
						UpdatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
					}

					got02.UpdatedAt = time.Time{}
					expected02.UpdatedAt = time.Time{}

					assert.Equal(t, expected02, got02)
				}
			},
			checkErr: func(t *testing.T, err error) {
				t.Helper()

				require.NoError(t, err)
			},
		},
		{
			name: "should return an error when insert a transfer",
			args: args{
				trf: entities.TransferData{
					Transfer: entities.Transfer{
						ID:                   trfID,
						AccountOriginID:      accOriginID,
						AccountDestinationID: accDestinationID,
						Amount:               30,
						CreatedAt:            time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
						UpdatedAt:            time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
					},
				},
			},
			beforeRun: func(t *testing.T, db *pgxpool.Pool) {
				t.Helper()

				{
					_, err := db.Exec(context.Background(), "ALTER TABLE transfers RENAME COLUMN amount TO wrong_amount")
					require.NoError(t, err)
				}
			},
			checkErr: func(t *testing.T, err error) {
				t.Helper()

				var pgErr *pgconn.PgError
				require.ErrorAs(t, err, &pgErr)
				assert.Equal(t, `column "amount" of relation "transfers" does not exist`, pgErr.Message)
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

			err := repo.Insert(context.Background(), tt.args.trf)
			tt.checkErr(t, err)

			if tt.check != nil {
				tt.check(t, pgPool)
			}
		})
	}
}
