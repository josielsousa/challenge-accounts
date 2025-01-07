package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
)

func TestAccount_GetAccountBalance(t *testing.T) {
	t.Parallel()

	errUnexpected := errors.New("unexpected error")

	type args struct {
		ctx       context.Context
		accountID string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
		setupUC func(t *testing.T) *Account
	}{
		{
			name: "should return balance by account id",
			args: args{
				ctx:       context.Background(),
				accountID: "d079de7d-b3d2-47fa-b1d6-b5c7d7cf5389",
			},
			wantErr: nil,
			want:    50,
			setupUC: func(t *testing.T) *Account {
				mockAccRepo := &accounts.RepositoryMock{
					GetByIDFunc: func(ctx context.Context, id string) (accounts.Account, error) {
						return accounts.Account{
							Balance: 50,
						}, nil
					},
				}

				return NewUsecase(mockAccRepo)
			},
		},
		{
			name: "should return an error when account not found",
			args: args{
				ctx:       context.Background(),
				accountID: "d079de7d-b3d2-47fa-b1d6-b5c7d7cf5389",
			},
			want: 0,
			setupUC: func(t *testing.T) *Account {
				mockAccRepo := &accounts.RepositoryMock{
					GetByIDFunc: func(ctx context.Context, id string) (accounts.Account, error) {
						return accounts.Account{}, accounts.ErrAccountNotFound
					},
				}

				return NewUsecase(mockAccRepo)
			},
			wantErr: accounts.ErrAccountNotFound,
		},
		{
			name: "should return an unexpected error",
			args: args{
				ctx:       context.Background(),
				accountID: "d079de7d-b3d2-47fa-b1d6-b5c7d7cf5389",
			},
			want: 0,
			setupUC: func(t *testing.T) *Account {
				mockAccRepo := &accounts.RepositoryMock{
					GetByIDFunc: func(ctx context.Context, id string) (accounts.Account, error) {
						return accounts.Account{}, errUnexpected
					},
				}

				return NewUsecase(mockAccRepo)
			},
			wantErr: errUnexpected,
		},
	}
	for _, tt := range tests {
		tt := tt // capture range variable

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := tt.setupUC(t)

			got, err := a.GetAccountBalance(tt.args.ctx, tt.args.accountID)
			assert.ErrorIs(t, err, tt.wantErr)

			assert.Equal(t, tt.want, got)
		})
	}
}
