package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJwt_SignToken(t *testing.T) {
	t.Parallel()

	type fields struct {
		appKey  []byte
		clocker Clocker
	}

	type args struct {
		id       string
		username string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr error
	}{
		{
			name: "sign token",
			args: args{
				id:       "acc-id-001",
				username: "superuser",
			},
			fields: fields{
				appKey: []byte("app-key"),
				clocker: &ClockerMock{
					NowFunc: func() time.Time {
						return time.Date(2025, time.January, 19, 13, 14, 25, 0, time.UTC)
					},
				},
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InN1cGVydXNlciIsImFjY291bnRfaWQiOiJhY2MtaWQtMDAxIiwiZXhwaXJlc19hdCI6MTczNzI5Mjc2NX0.EA4Thsie8vbapJh5pIvHB9RpdWbcLTodcjFQQIWLUNY",
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

			got, err := signer.SignToken(tt.args.id, tt.args.username)
			require.ErrorIs(t, err, tt.wantErr)

			assert.Equal(t, tt.want, got)
		})
	}
}
