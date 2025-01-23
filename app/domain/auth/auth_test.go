package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
	"github.com/josielsousa/challenge-accounts/app/domain/erring"
	"github.com/josielsousa/challenge-accounts/app/types"
)

func TestUsecase_Signin(t *testing.T) {
	t.Parallel()

	type fields struct {
		R Repository
		H Hasher
		S Signer
	}

	type args struct {
		credential types.Credentials
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    types.Auth
		wantErr error
	}{
		{
			name: "give a account not found",
			args: args{
				credential: types.Credentials{
					Cpf:    "12345678901",
					Secret: "secr3t",
				},
			},
			fields: fields{
				R: &RepositoryMock{
					GetByCPFFunc: func(_ context.Context, _ string) (entities.Account, error) {
						return entities.Account{}, nil
					},
				},
			},
			want:    types.Auth{},
			wantErr: erring.ErrAccountNotFound,
		},
		{
			name: "give a unauthorized",
			args: args{
				credential: types.Credentials{
					Cpf:    "12345678901",
					Secret: "secr3t",
				},
			},
			fields: fields{
				R: &RepositoryMock{
					GetByCPFFunc: func(_ context.Context, _ string) (entities.Account, error) {
						return entities.Account{
							ID: "acc-id-001",
						}, nil
					},
				},
				H: &HasherMock{
					VerifySecretFunc: func(_, _ string) error {
						return erring.ErrUnauthorized
					},
				},
			},
			want:    types.Auth{},
			wantErr: erring.ErrUnauthorized,
		},
		{
			name: "give a unexpected",
			args: args{
				credential: types.Credentials{
					Cpf:    "12345678901",
					Secret: "secr3t",
				},
			},
			fields: fields{
				R: &RepositoryMock{
					GetByCPFFunc: func(_ context.Context, _ string) (entities.Account, error) {
						return entities.Account{
							ID: "acc-id-001",
						}, nil
					},
				},
				H: &HasherMock{
					VerifySecretFunc: func(_, _ string) error {
						return nil
					},
				},
				S: &SignerMock{
					SignTokenFunc: func(_, _ string) (types.Auth, error) {
						return types.Auth{}, erring.ErrUnexpected
					},
				},
			},
			want:    types.Auth{},
			wantErr: erring.ErrUnexpected,
		},
		{
			name: "singin success",
			args: args{
				credential: types.Credentials{
					Cpf:    "12345678901",
					Secret: "secr3t",
				},
			},
			fields: fields{
				R: &RepositoryMock{
					GetByCPFFunc: func(_ context.Context, _ string) (entities.Account, error) {
						return entities.Account{
							ID: "acc-id-001",
						}, nil
					},
				},
				H: &HasherMock{
					VerifySecretFunc: func(_, _ string) error {
						return nil
					},
				},
				S: &SignerMock{
					SignTokenFunc: func(_, _ string) (types.Auth, error) {
						return types.Auth{
							Token: "t0k3N",
						}, nil
					},
				},
			},
			want: types.Auth{
				Token: "t0k3N",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			usecase := Usecase{
				R: tt.fields.R,
				H: tt.fields.H,
				S: tt.fields.S,
			}

			got, err := usecase.Signin(context.Background(), tt.args.credential)
			require.ErrorIs(t, err, tt.wantErr)

			assert.Equal(t, tt.want, got)
		})
	}
}
