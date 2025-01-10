package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"

	"github.com/josielsousa/challenge-accounts/app/common"
)

func TestNewHash(t *testing.T) {
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
			name: "success create a new hash",
			args: args{
				secret: "teste",
			},
			wantErr: nil,
		},
		{
			name: "success create a new hash with special chars",
			args: args{
				secret: "the#$%PassWoRdok",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewHash(tt.args.secret)
			require.ErrorIs(t, err, tt.wantErr)

			err = bcrypt.CompareHashAndPassword([]byte(got.Value()), []byte(tt.args.secret))
			require.NoError(t, err)
		})
	}
}

func TestHash_Compare(t *testing.T) {
	t.Parallel()

	type args struct {
		secretToCompare string
		valueToHash     string
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "should compare secrets with success",
			args: args{
				valueToHash:     "teste",
				secretToCompare: "teste",
			},
			wantErr: nil,
		},
		{
			name: "should return an error when compare secret",
			args: args{
				valueToHash:     "teste",
				secretToCompare: "teste2",
			},
			wantErr: ErrInvalidSecret,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			hash, err := bcrypt.GenerateFromPassword([]byte(tt.args.valueToHash), bcrypt.MinCost)
			require.NoError(t, err)

			h := Hash{
				hashedValue: string(hash),
			}

			err = h.Compare(tt.args.secretToCompare)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestHash_Scan(t *testing.T) {
	t.Parallel()

	hashedSecret, err := bcrypt.GenerateFromPassword([]byte("th3Secr3T"), 4)
	require.NoError(t, err)

	type args struct {
		value interface{}
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "should scan the value with successfully",
			args: args{
				value: string(hashedSecret),
			},
			wantErr: nil,
		},
		{
			name: "should return an error when scan value is nil",
			args: args{
				value: nil,
			},
			wantErr: common.ErrScanValueNil,
		},
		{
			name: "should return an error when scan value not is an string",
			args: args{
				value: 123789,
			},
			wantErr: common.ErrScanValueIsNotString,
		},
		{
			name: "should return an error when scan value is invalid hash",
			args: args{
				value: "th3Secr3T",
			},
			wantErr: ErrScanInvalidSecret,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h := &Hash{
				hashedValue: "",
			}

			err := h.Scan(tt.args.value)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
