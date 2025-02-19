package cache

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/gateway/db/cache/cachetest"
)

func TestCache_SetIdempotency(t *testing.T) {
	t.Parallel()

	type args struct {
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
			name: "should set idempotency",
			args: args{
				key: "test_key",
				response: &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader("test_body")),
				},
				ttl: time.Minute,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cli, err := cachetest.NewClient()
			require.NoError(t, err)

			c := &Cache{
				C: cli.Client,
			}

			t.Cleanup(func() {
				require.NoError(t, cli.Teardown())
			})

			ctx := context.Background()

			err = c.SetIdempotency(
				ctx,
				tt.args.key,
				tt.args.response,
				tt.args.ttl,
			)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}
