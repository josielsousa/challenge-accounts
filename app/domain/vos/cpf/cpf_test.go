package cpf

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCPF_NewCPF(t *testing.T) {
	t.Parallel()

	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "should create a cpf from string",
			args: args{
				value: "883.500.570-17",
			},
			want:    "88350057017",
			wantErr: nil,
		},
		{
			name: "should create a cpf from string without format",
			args: args{
				value: "80545919002",
			},
			want:    "80545919002",
			wantErr: nil,
		},
		{
			name: "should create a cpf from string with padding",
			args: args{
				value: " 93388834008  ",
			},
			want:    "93388834008",
			wantErr: nil,
		},
		{
			name: "should return an error when create a cpf from empty string",
			args: args{
				value: "",
			},
			want:    "",
			wantErr: ErrInvalid,
		},
		{
			name: "should return an error when create a cpf from string invalid size",
			args: args{
				value: "601",
			},
			want:    "",
			wantErr: ErrInvalid,
		},
		{
			name: "should return an error when create a cpf from string all equals",
			args: args{
				value: "111.111.111-11",
			},
			want:    "",
			wantErr: ErrInvalid,
		},
		{
			name: "should return an error when create a cpf with validator digit invalid",
			args: args{
				value: "883.500.570-20",
			},
			want:    "",
			wantErr: ErrInvalid,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewCPF(tt.args.value)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got.Value())
		})
	}
}

func TestCPF_Mask(t *testing.T) {
	t.Parallel()

	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return a cpf masked from string",
			args: args{
				value: "88350057017",
			},
			want: "883.500.570-17",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c, err := NewCPF(tt.args.value)
			require.NoError(t, err)

			got := c.Mask()
			assert.Equal(t, tt.want, got)
		})
	}
}
