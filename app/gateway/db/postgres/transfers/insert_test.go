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

	"github.com/josielsousa/challenge-accounts/app/domain/entities/transfers"
	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres/pgtest"
)

func TestRepository_Insert(t *testing.T) {
	t.Parallel()

	trfID := uuid.NewString()
	accOriginID := uuid.NewString()
	accDestinationID := uuid.NewString()

	type args struct {
		ctx context.Context
		trf transfers.TransferData
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
				ctx: context.Background(),
				trf: transfers.TransferData{
					Transfer: transfers.Transfer{
						ID:                   trfID,
						AccountOriginID:      accOriginID,
						AccountDestinationID: accDestinationID,
						Amount:               30,
						CreatedAt:            time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
						UpdatedAt:            time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
					},
				},
			},
			check: func(t *testing.T, pgPool *pgxpool.Pool) {
				{
					got, err := pgtest.GetTransfer(t, pgPool, trfID)
					require.NoError(t, err)

					expected := transfers.Transfer{
						ID:                   trfID,
						AccountOriginID:      accOriginID,
						AccountDestinationID: accDestinationID,
						Amount:               30,
						CreatedAt:            time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
						UpdatedAt:            time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
					}

					assert.Equal(t, expected, got)
				}
			},
			checkErr: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "should return an error when insert a transfer",
			args: args{
				ctx: context.Background(),
				trf: transfers.TransferData{
					Transfer: transfers.Transfer{
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
				{
					_, err := db.Exec(context.Background(), "ALTER TABLE transfers RENAME COLUMN amount TO wrong_amount")
					require.NoError(t, err)
				}
			},
			checkErr: func(t *testing.T, err error) {
				var pgErr *pgconn.PgError
				assert.ErrorAs(t, err, &pgErr)
				assert.Equal(t, `column "amount" of relation "transfers" does not exist`, pgErr.Message)
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

			err := r.Insert(tt.args.ctx, tt.args.trf)
			tt.checkErr(t, err)

			if tt.check != nil {
				tt.check(t, pgPool)
			}
		})
	}
}
