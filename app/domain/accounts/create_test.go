package accounts

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
	"github.com/josielsousa/challenge-accounts/app/domain/erring"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
)

func TestAccount_Create(t *testing.T) {
	t.Parallel()

	errUnexpected := errors.New("unexpected error")

	type args struct {
		input AccountInput
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		want    AccountOutput
		setupUC func() *Usecase
	}{
		{
			name: "should create a new account from values",
			args: args{
				input: AccountInput{
					Name:    "username test",
					Balance: 50,
					CPF:     "883.500.570-17",
					Secret:  "stringSecret",
				},
			},
			setupUC: func() *Usecase {
				t.Helper()

				mockAccRepo := &RepositoryMock{
					InsertFunc: func(_ context.Context, _ entities.Account) error {
						return nil
					},
				}

				return NewUsecase(mockAccRepo)
			},
			want: AccountOutput{
				ID:        "",
				Name:      "username test",
				Balance:   50,
				CPF:       "88350057017",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			wantErr: nil,
		},
		{
			name: "should return an error when create a new account, account already exists",
			args: args{
				input: AccountInput{
					Name:    "username test",
					Balance: 0,
					CPF:     "88350057017",
					Secret:  "stringSecret",
				},
			},
			setupUC: func() *Usecase {
				t.Helper()

				mockAccRepo := &RepositoryMock{
					InsertFunc: func(_ context.Context, _ entities.Account) error {
						return erring.ErrAccountAlreadyExists
					},
				}

				return NewUsecase(mockAccRepo)
			},
			want:    AccountOutput{},
			wantErr: erring.ErrAccountAlreadyExists,
		},
		{
			name: "should return an error when create a new account, invalid cpf number",
			args: args{
				input: AccountInput{
					Name:    "username test",
					Balance: 0,
					CPF:     "88350057013",
					Secret:  "stringSecret",
				},
			},
			setupUC: func() *Usecase {
				t.Helper()

				mockAccRepo := &RepositoryMock{
					InsertFunc: func(_ context.Context, _ entities.Account) error {
						return erring.ErrAccountAlreadyExists
					},
				}

				return NewUsecase(mockAccRepo)
			},
			want:    AccountOutput{},
			wantErr: cpf.ErrInvalid,
		},
		{
			name: "should return an error when create a new account, unexpected error",
			args: args{
				input: AccountInput{
					Name:    "username test",
					Balance: 0,
					CPF:     "88350057017",
					Secret:  "stringSecret",
				},
			},
			setupUC: func() *Usecase {
				t.Helper()

				mockAccRepo := &RepositoryMock{
					InsertFunc: func(_ context.Context, _ entities.Account) error {
						return errUnexpected
					},
				}

				return NewUsecase(mockAccRepo)
			},
			want:    AccountOutput{},
			wantErr: errUnexpected,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a := tt.setupUC()

			out, err := a.Create(context.Background(), tt.args.input)
			require.ErrorIs(t, err, tt.wantErr)

			out.ID = ""
			out.CreatedAt = time.Time{}
			out.UpdatedAt = time.Time{}
			require.Equal(t, tt.want, out)
		})
	}
}
