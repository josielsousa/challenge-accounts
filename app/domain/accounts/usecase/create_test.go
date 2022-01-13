package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	accUC "github.com/josielsousa/challenge-accounts/app/domain/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
)

func TestAccount_Create(t *testing.T) {
	t.Parallel()

	errUnexpected := errors.New("unexpected error")

	type args struct {
		ctx   context.Context
		input accUC.AccountInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
		setupUC func(t *testing.T) *Account
	}{
		{
			name: "should create a new account from values",
			args: args{
				ctx: context.Background(),
				input: accUC.AccountInput{
					Name:    "username test",
					Balance: 50,
					CPF:     "883.500.570-17",
					Secret:  "stringSecret",
				},
			},
			setupUC: func(_ *testing.T) *Account {
				mockAccRepo := &accounts.RepositoryMock{
					InsertFunc: func(_ context.Context, _ accounts.Account) error {
						return nil
					},
				}

				return NewUsecase(mockAccRepo)
			},
			wantErr: nil,
		},
		{
			name: "should return an error when create a new account, account already exists",
			args: args{
				ctx: context.Background(),
				input: accUC.AccountInput{
					Name:    "username test",
					Balance: 0,
					CPF:     "88350057017",
					Secret:  "stringSecret",
				},
			},
			setupUC: func(_ *testing.T) *Account {
				mockAccRepo := &accounts.RepositoryMock{
					InsertFunc: func(_ context.Context, _ accounts.Account) error {
						return accounts.ErrAccountAlreadyExists
					},
				}

				return NewUsecase(mockAccRepo)
			},
			wantErr: accounts.ErrAccountAlreadyExists,
		},
		{
			name: "should return an error when create a new account, invalid cpf number",
			args: args{
				ctx: context.Background(),
				input: accUC.AccountInput{
					Name:    "username test",
					Balance: 0,
					CPF:     "88350057013",
					Secret:  "stringSecret",
				},
			},
			setupUC: func(_ *testing.T) *Account {
				mockAccRepo := &accounts.RepositoryMock{
					InsertFunc: func(_ context.Context, _ accounts.Account) error {
						return accounts.ErrAccountAlreadyExists
					},
				}

				return NewUsecase(mockAccRepo)
			},
			wantErr: cpf.ErrInvalid,
		},
		{
			name: "should return an error when create a new account, unexpected error",
			args: args{
				ctx: context.Background(),
				input: accUC.AccountInput{
					Name:    "username test",
					Balance: 0,
					CPF:     "88350057017",
					Secret:  "stringSecret",
				},
			},
			setupUC: func(_ *testing.T) *Account {
				mockAccRepo := &accounts.RepositoryMock{
					InsertFunc: func(_ context.Context, _ accounts.Account) error {
						return errUnexpected
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

			err := a.Create(tt.args.ctx, tt.args.input)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
