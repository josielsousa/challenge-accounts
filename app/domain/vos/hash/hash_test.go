package vos

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHash_CompareHashAndSecret(t *testing.T) {
	t.Parallel()
	type fields struct {
		secret string
	}

	type args struct {
		secret string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "should compare with sucessfully",
			args: args{
				secret: "!th3@Secret#",
			},
			fields: fields{
				secret: "!th3@Secret#",
			},
			wantErr: nil,
		},
		{
			name: "should return an error when secret is a empty string",
			args: args{
				secret: "",
			},
			fields: fields{
				secret: "th3Secret#",
			},
			wantErr: ErrInvalidSecret,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			h, err := NewHash(tt.fields.secret)
			require.NoError(t, err)

			err = h.CompareHashAndSecret(tt.args.secret)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestHash_NewHash(t *testing.T) {
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
	}
	for _, tt := range tests {
		tt := tt // capture range variable

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := NewHash(tt.args.secret)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
