package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestGenHash(t *testing.T) {
	t.Parallel()

	type args struct {
		secret string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "sucess create a new hash",
			args: args{
				secret: "teste",
			},
			wantErr: nil,
		},
		{
			name: "sucess create a new hash with special chars",
			args: args{
				secret: "the#$%PassWoRdok",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt := tt // capture range variable

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := GenHash(tt.args.secret)
			assert.ErrorIs(t, err, tt.wantErr)

			err = bcrypt.CompareHashAndPassword([]byte(got), []byte(tt.args.secret))
			require.NoError(t, err)
		})
	}
}

func TestCompareHashedAndSecret(t *testing.T) {
	t.Parallel()

	type args struct {
		hashedSecret string
		secret       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "should successfully compare hashed string with secret value",
			args: args{
				secret:       "teste",
				hashedSecret: "$2a$10$GUNat4gFNNRKtRZ25DO82.L/XcrcmF8eKpt7/PCGsBvqiAJKx63Au",
			},
			wantErr: nil,
		},
		{
			name: "should successfully compare hashed string with secret value",
			args: args{
				secret:       "the#$%PassWoRd",
				hashedSecret: "$2a$10$oUydrP.MQZq7gvLpCvzGaOKBAqwBAoRgzqz7pLks3C0ulIkrpSEQa",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt := tt // capture range variable

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := CompareHashedAndSecret(tt.args.hashedSecret, tt.args.secret)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
