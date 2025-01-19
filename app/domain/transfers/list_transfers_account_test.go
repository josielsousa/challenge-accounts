package transfers

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	trfE "github.com/josielsousa/challenge-accounts/app/domain/entities/transfers"
)

func TestUsecase_ListTransfersAccount(t *testing.T) {
	t.Parallel()

	errUnknown := errors.New("unknown error")

	type fields struct {
		R trfE.Repository
	}

	type args struct {
		accOriginID string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []TransferOutput
		wantErr error
	}{
		{
			name: "give an error on list transfers",
			args: args{
				accOriginID: "acc-id-001",
			},
			fields: fields{
				R: &trfE.RepositoryMock{
					ListTransfersFunc: func(_ context.Context, id string) ([]trfE.Transfer, error) {
						require.Equal(t, "acc-id-001", id)

						return []trfE.Transfer{}, errUnknown
					},
				},
			},
			want:    nil,
			wantErr: errUnknown,
		},
		{
			name: "give an list of transfers",
			args: args{
				accOriginID: "acc-id-001",
			},
			fields: fields{
				R: &trfE.RepositoryMock{
					ListTransfersFunc: func(_ context.Context, id string) ([]trfE.Transfer, error) {
						require.Equal(t, "acc-id-001", id)

						return []trfE.Transfer{
							{
								ID:                   "trf-id-001",
								AccountOriginID:      "acc-id-002",
								AccountDestinationID: "acc-id-001",
								Amount:               5_00,
								CreatedAt:            time.Time{},
								UpdatedAt:            time.Time{},
							},
							{
								ID:                   "trf-id-002",
								AccountOriginID:      "acc-id-001",
								AccountDestinationID: "acc-id-003",
								Amount:               5_00,
								CreatedAt:            time.Time{},
								UpdatedAt:            time.Time{},
							},
						}, nil
					},
				},
			},
			want: []TransferOutput{
				{
					ID:                   "trf-id-001",
					AccountOriginID:      "acc-id-002",
					AccountDestinationID: "acc-id-001",
					Amount:               5_00,
					CreatedAt:            time.Time{},
				},
				{
					ID:                   "trf-id-002",
					AccountOriginID:      "acc-id-001",
					AccountDestinationID: "acc-id-003",
					Amount:               5_00,
					CreatedAt:            time.Time{},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			usecase := Usecase{
				R: tt.fields.R,
			}

			got, err := usecase.ListTransfersAccount(context.Background(), tt.args.accOriginID)
			require.ErrorIs(t, err, tt.wantErr)

			assert.Equal(t, tt.want, got)
		})
	}
}
