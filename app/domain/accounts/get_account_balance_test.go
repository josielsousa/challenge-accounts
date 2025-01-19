package accounts

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	accE "github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/erring"
)

func TestAccount_GetAccountBalance(t *testing.T) {
	t.Parallel()

	errUnexpected := errors.New("unexpected error")

	type args struct {
		accountID string
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
		setupUC func() *Account
	}{
		{
			name: "should return balance by account id",
			args: args{
				accountID: "d079de7d-b3d2-47fa-b1d6-b5c7d7cf5389",
			},
			wantErr: nil,
			want:    50,
			setupUC: func() *Account {
				t.Helper()

				mockAccRepo := &RepositoryMock{
					GetByIDFunc: func(_ context.Context, _ string) (accE.Account, error) {
						return accE.Account{
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
				accountID: "d079de7d-b3d2-47fa-b1d6-b5c7d7cf5389",
			},
			want: 0,
			setupUC: func() *Account {
				t.Helper()

				mockAccRepo := &RepositoryMock{
					GetByIDFunc: func(_ context.Context, _ string) (accE.Account, error) {
						return accE.Account{}, erring.ErrAccountNotFound
					},
				}

				return NewUsecase(mockAccRepo)
			},
			wantErr: erring.ErrAccountNotFound,
		},
		{
			name: "should return an unexpected error",
			args: args{
				accountID: "d079de7d-b3d2-47fa-b1d6-b5c7d7cf5389",
			},
			want: 0,
			setupUC: func() *Account {
				t.Helper()

				mockAccRepo := &RepositoryMock{
					GetByIDFunc: func(_ context.Context, _ string) (accE.Account, error) {
						return accE.Account{}, errUnexpected
					},
				}

				return NewUsecase(mockAccRepo)
			},
			wantErr: errUnexpected,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a := tt.setupUC()

			got, err := a.GetAccountBalance(context.Background(), tt.args.accountID)
			require.ErrorIs(t, err, tt.wantErr)

			assert.Equal(t, tt.want, got)
		})
	}
}
