package accounts

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/hash"
	"github.com/josielsousa/challenge-accounts/app/gateway/db/postgres/pgtest"
)

func TestRepository_Insert(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		acc accounts.Account
	}
	tests := []struct {
		name    string
		args    func(t *testing.T) args
		wantErr error
	}{
		{
			name: "should save account with successfully",
			args: func(t *testing.T) args {
				secret, err := hash.NewHash("the#$%PassWoRd")
				require.NoError(t, err)

				newCpf, err := cpf.NewCPF("88350057017")
				require.NoError(t, err)

				return args{
					acc: accounts.Account{
						Name:    "Teste",
						Balance: 350_00,
						Secret:  secret,
						CPF:     newCpf,
					},
					ctx: context.Background(),
				}
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt := tt // capture range variable

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			pgPool := pgtest.NewDB(t)
			r := NewRepository(pgPool)

			args := tt.args(t)

			err := r.Insert(args.ctx, args.acc)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
