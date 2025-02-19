package cache

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCache_SetIdempotency(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx      context.Context
		key      string
		response *http.Response
		ttl      time.Duration
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name:    "test case name here",
			args:    args{},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &Cache{
				C: tt.fields.C,
			}

			err := c.SetIdempotency(tt.args.ctx, tt.args.key, tt.args.response)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}
