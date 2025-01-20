package accounts

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
)

func TestAccount_GetAllAccounts(t *testing.T) {
	t.Parallel()

	newCpf01, err := cpf.NewCPF("88350057017")
	require.NoError(t, err)

	newCpf02, err := cpf.NewCPF("88350057017")
	require.NoError(t, err)

	errUnexpected := errors.New("unexpected error")

	tests := []struct {
		name    string
		want    []AccountOutput
		wantErr error
		setupUC func() *Usecase
	}{
		{
			name:    "should return an empty list when db is empty",
			wantErr: nil,
			want:    []AccountOutput{},
			setupUC: func() *Usecase {
				t.Helper()

				mockAccRepo := &RepositoryMock{
					GetAllFunc: func(_ context.Context) ([]entities.Account, error) {
						return nil, nil
					},
				}

				return NewUsecase(mockAccRepo)
			},
		},
		{
			name:    "should return a list with all accounts",
			wantErr: nil,
			want: []AccountOutput{
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
			setupUC: func() *Usecase {
				t.Helper()

				mockAccRepo := &RepositoryMock{
					GetAllFunc: func(_ context.Context) ([]entities.Account, error) {
						return []entities.Account{
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
			name:    "should return an unexpected error",
			wantErr: errUnexpected,
			want:    nil,
			setupUC: func() *Usecase {
				t.Helper()

				mockAccRepo := &RepositoryMock{
					GetAllFunc: func(_ context.Context) ([]entities.Account, error) {
						return nil, errUnexpected
					},
				}

				return NewUsecase(mockAccRepo)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a := tt.setupUC()

			got, err := a.GetAllAccounts(context.Background())
			require.ErrorIs(t, err, tt.wantErr)

			assert.Equal(t, tt.want, got)
		})
	}
}
