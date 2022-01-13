package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	accUC "github.com/josielsousa/challenge-accounts/app/domain/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccount_GetAllAccounts(t *testing.T) {
	t.Parallel()

	newCpf01, err := cpf.NewCPF("88350057017")
	require.NoError(t, err)

	newCpf02, err := cpf.NewCPF("88350057017")
	require.NoError(t, err)

	errUnexpected := errors.New("unexpected error")

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []accUC.AccountOutput
		wantErr error
		setupUC func(t *testing.T) *Account
	}{
		{
			name: "should return an empty list when db is empty",
			args: args{
				ctx: context.Background(),
			},
			wantErr: nil,
			want:    []accUC.AccountOutput{},
			setupUC: func(t *testing.T) *Account {
				mockAccRepo := &accounts.RepositoryMock{
					GetAllFunc: func(ctx context.Context) ([]accounts.Account, error) {
						return nil, nil
					},
				}

				return NewUsecase(mockAccRepo)
			},
		},
		{
			name: "should return a list with all accounts",
			args: args{
				ctx: context.Background(),
			},
			wantErr: nil,
			want: []accUC.AccountOutput{
				{
					ID:        "acc-id-01",
					Name:      "name test account 01",
					Balance:   999_00,
					CPF:       newCpf01,
					CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
					UpdatedAt: time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
				},
				{
					ID:        "acc-id-02",
					Name:      "name test account 02",
					Balance:   -988_00,
					CPF:       newCpf02,
					CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
					UpdatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
				},
			},
			setupUC: func(t *testing.T) *Account {
				mockAccRepo := &accounts.RepositoryMock{
					GetAllFunc: func(ctx context.Context) ([]accounts.Account, error) {
						return []accounts.Account{
							{
								ID:        "acc-id-01",
								Name:      "name test account 01",
								Balance:   999_00,
								CPF:       newCpf01,
								CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
								UpdatedAt: time.Date(2022, time.January, 4, 1, 0, 0, 0, time.Local),
							},
							{
								ID:        "acc-id-02",
								Name:      "name test account 02",
								Balance:   -988_00,
								CPF:       newCpf02,
								CreatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
								UpdatedAt: time.Date(2022, time.January, 4, 0, 0, 0, 0, time.Local),
							},
						}, nil
					},
				}

				return NewUsecase(mockAccRepo)
			},
		},
		{
			name: "should return an unexpected error",
			args: args{
				ctx: context.Background(),
			},
			wantErr: errUnexpected,
			want:    nil,
			setupUC: func(t *testing.T) *Account {
				mockAccRepo := &accounts.RepositoryMock{
					GetAllFunc: func(ctx context.Context) ([]accounts.Account, error) {
						return nil, errUnexpected
					},
				}

				return NewUsecase(mockAccRepo)
			},
		},
	}
	for _, tt := range tests {
		tt := tt // capture range variable

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := tt.setupUC(t)

			got, err := a.GetAllAccounts(tt.args.ctx)
			assert.ErrorIs(t, err, tt.wantErr)

			assert.Equal(t, tt.want, got)
		})
	}
}
