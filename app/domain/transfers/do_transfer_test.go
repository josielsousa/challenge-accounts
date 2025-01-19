package transfers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	accE "github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	trfE "github.com/josielsousa/challenge-accounts/app/domain/entities/transfers"
	"github.com/josielsousa/challenge-accounts/app/domain/erring"
)

func TestUsecase_DoTransfer(t *testing.T) {
	t.Parallel()

	type fields struct {
		AR AccountRepository
		R  Repository
	}

	type args struct {
		input TransferInput
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "give an invalid amount error",
			args: args{
				input: TransferInput{
					AccountOriginID:      "acc-id-002",
					AccountDestinationID: "acc-id-001",
					Amount:               0,
				},
			},
			wantErr: erring.ErrInvalidAmount,
		},
		{
			name: "give an insufficient funds error",
			args: args{
				input: TransferInput{
					AccountOriginID:      "acc-id-002",
					AccountDestinationID: "acc-id-001",
					Amount:               10_00,
				},
			},
			fields: fields{
				AR: &AccountRepositoryMock{
					GetByIDFunc: func(_ context.Context, id string) (accE.Account, error) {
						accounts := map[string]accE.Account{
							"acc-id-001": {Balance: 15_00},
							"acc-id-002": {Balance: 5_00},
						}

						return accounts[id], nil
					},
				},
			},
			wantErr: erring.ErrInsufficientFunds,
		},
		{
			name: "give an account origin not found error",
			args: args{
				input: TransferInput{
					AccountOriginID:      "acc-id-002",
					AccountDestinationID: "acc-id-001",
					Amount:               10_00,
				},
			},
			fields: fields{
				AR: &AccountRepositoryMock{
					GetByIDFunc: func(_ context.Context, id string) (accE.Account, error) {
						accounts := map[string]accE.Account{
							"acc-id-001": {Balance: 15_00},
						}

						acc, ok := accounts[id]
						if !ok {
							return accE.Account{}, erring.ErrAccountOriginNotFound
						}

						return acc, nil
					},
				},
			},
			wantErr: erring.ErrAccountOriginNotFound,
		},
		{
			name: "give an account destination not found error",
			args: args{
				input: TransferInput{
					AccountOriginID:      "acc-id-002",
					AccountDestinationID: "acc-id-001",
					Amount:               10_00,
				},
			},
			fields: fields{
				AR: &AccountRepositoryMock{
					GetByIDFunc: func(_ context.Context, id string) (accE.Account, error) {
						accounts := map[string]accE.Account{
							"acc-id-002": {Balance: 15_00},
						}

						acc, ok := accounts[id]
						if !ok {
							return accE.Account{}, erring.ErrAccountDestinationNotFound
						}

						return acc, nil
					},
				},
			},
			wantErr: erring.ErrAccountDestinationNotFound,
		},
		{
			name: "should do a transfer",
			args: args{
				input: TransferInput{
					AccountOriginID:      "acc-id-002",
					AccountDestinationID: "acc-id-001",
					Amount:               10_00,
				},
			},
			fields: fields{
				AR: &AccountRepositoryMock{
					GetByIDFunc: func(_ context.Context, id string) (accE.Account, error) {
						accounts := map[string]accE.Account{
							"acc-id-002": {ID: "acc-id-002", Balance: 15_00},
							"acc-id-001": {ID: "acc-id-001", Balance: 15_00},
						}

						acc, ok := accounts[id]
						if !ok {
							return accE.Account{}, erring.ErrAccountDestinationNotFound
						}

						return acc, nil
					},
				},
				R: &RepositoryMock{
					InsertFunc: func(_ context.Context, data trfE.TransferData) error {
						require.Equal(t, 10_00, data.Amount)
						require.Equal(t, 5_00, data.AccountOrigin.Balance)
						require.Equal(t, 25_00, data.AccountDestination.Balance)

						return nil
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			usecase := Usecase{
				AR: tt.fields.AR,
				R:  tt.fields.R,
			}

			err := usecase.DoTransfer(context.Background(), tt.args.input)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}
