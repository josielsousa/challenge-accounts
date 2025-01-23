package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/domain/erring"
	"github.com/josielsousa/challenge-accounts/app/types"
)

func TestJwt_Authorize(t *testing.T) {
	t.Parallel()

	//nolint:gosec
	defaultToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InN1cGVydXNlciIsImFjY291bnRfaWQiOiJhY2MtaWQtMDAxIiwiZXhwaXJlc19hdCI6MTczNzI5Mjc2NX0.EA4Thsie8vbapJh5pIvHB9RpdWbcLTodcjFQQIWLUNY"

	type fields struct {
		appKey  []byte
		clocker Clocker
	}

	type args struct {
		token string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    types.Claims
		wantErr error
	}{
		{
			name: "give an empty token",
			args: args{
				token: "",
			},
			fields: fields{
				appKey:  []byte("app-key"),
				clocker: &ClockerMock{},
			},
			want:    types.Claims{},
			wantErr: erring.ErrEmptyToken,
		},
		{
			name: "give an empty token",
			args: args{
				token: "null",
			},
			fields: fields{
				appKey:  []byte("app-key"),
				clocker: &ClockerMock{},
			},
			want:    types.Claims{},
			wantErr: erring.ErrEmptyToken,
		},
		{
			name: "give an error when parse token",
			args: args{
				token: defaultToken,
			},
			fields: fields{
				appKey:  []byte("app-key-wrong"),
				clocker: &ClockerMock{},
			},
			want:    types.Claims{},
			wantErr: erring.ErrParseTokenWithClaims,
		},
		{
			name: "give an expired token",
			args: args{
				token: defaultToken,
			},
			fields: fields{
				appKey: []byte("app-key"),
				clocker: &ClockerMock{
					UntilFunc: func(_ time.Time) time.Duration {
						return -10
					},
				},
			},
			want:    types.Claims{},
			wantErr: erring.ErrExpiredToken,
		},
		{
			name: "authorize token",
			args: args{
				token: defaultToken,
			},
			fields: fields{
				appKey: []byte("app-key"),
				clocker: &ClockerMock{
					UntilFunc: func(_ time.Time) time.Duration {
						return 1
					},
				},
			},
			want: types.Claims{
				Username:  "superuser",
				AccountID: "acc-id-001",
				ExpiresAt: 1737292765,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			signer := Jwt{
				appKey: tt.fields.appKey,
				C:      tt.fields.clocker,
			}

			got, err := signer.Authorize(tt.args.token)
			require.ErrorIs(t, err, tt.wantErr)

			assert.Equal(t, tt.want, got)
		})
	}
}
