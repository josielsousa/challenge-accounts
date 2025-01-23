package transfers

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres/pgtest"
)

func TestRepository_ListTransfers(t *testing.T) {
	t.Parallel()

	trfID1 := uuid.NewString()
	trfID2 := uuid.NewString()

	accOriginID := uuid.NewString()
	accDestinationID1 := uuid.NewString()
	accDestinationID2 := uuid.NewString()

	type args struct {
		accOriginID string
	}

	tests := []struct {
		name      string
		args      args
		want      []entities.Transfer
		wantErr   error
		beforeRun func(t *testing.T, db *pgxpool.Pool)
	}{
		{
			name: "should return an empty slice when db is empty",
			args: args{
				accOriginID: accOriginID,
			},
			want:    []entities.Transfer{},
			wantErr: nil,
		},
		{
			name: "should return an empty slice when db is empty",
			args: args{
				accOriginID: accOriginID,
			},
			want: []entities.Transfer{
				{
					ID:                   trfID1,
					AccountOriginID:      accOriginID,
					AccountDestinationID: accDestinationID1,
					Amount:               30,
					CreatedAt:            time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
					UpdatedAt:            time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
				},
				{
					ID:                   trfID2,
					AccountOriginID:      accOriginID,
					AccountDestinationID: accDestinationID2,
					Amount:               50,
					CreatedAt:            time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
					UpdatedAt:            time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
				},
			},
			wantErr: nil,
			beforeRun: func(t *testing.T, db *pgxpool.Pool) {
				t.Helper()

				{
					trfs := []entities.Transfer{
						{
							ID:                   trfID1,
							AccountOriginID:      accOriginID,
							AccountDestinationID: accDestinationID1,
							Amount:               30,
							CreatedAt:            time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
							UpdatedAt:            time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
						},
						{
							ID:                   trfID2,
							AccountOriginID:      accOriginID,
							AccountDestinationID: accDestinationID2,
							Amount:               50,
							CreatedAt:            time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
							UpdatedAt:            time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
						},
					}

					for _, trf := range trfs {
						err := pgtest.TransfersInsert(t, db, trf)
						require.NoError(t, err)
					}
				}
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

			got, err := repo.ListTransfers(context.Background(), tt.args.accOriginID)
			require.ErrorIs(t, err, tt.wantErr)

			assert.Equal(t, tt.want, got)
		})
	}
}
