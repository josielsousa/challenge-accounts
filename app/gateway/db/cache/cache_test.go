package cache

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httputil"
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

			ctx := context.Background()

			err = c.SetIdempotency(
				ctx,
				tt.args.key,
				tt.args.response,
				tt.args.ttl,
			)
			require.ErrorIs(t, err, tt.wantErr)

			res, err := cli.Client.Get(ctx, tt.args.key).Result()
			require.NoError(t, err)

			bs, err := httputil.DumpResponse(tt.args.response, true)
			require.NoError(t, err)

			require.Equal(t, string(bs), res)

			t.Cleanup(func() {
				require.NoError(t, cli.Teardown())
				tt.args.response.Body.Close()
			})
		})
	}
}

func TestCache_GetIdempotency(t *testing.T) {
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
			name: "should get idempotency",
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
		{
			name: "should get idempotency json body",
			args: args{
				key: "test_key",
				response: &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(
						`{"test": "test"}`,
					)),
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

			ctx := context.Background()

			dbs, err := httputil.DumpResponse(tt.args.response, true)
			require.NoError(t, err)

			_, err = cli.Client.Set(
				ctx,
				tt.args.key,
				string(dbs),
				tt.args.ttl,
			).Result()
			require.NoError(t, err)

			out, err := c.GetIdempotency(ctx, tt.args.key)
			require.ErrorIs(t, err, tt.wantErr)

			checkResponse(t, dbs, out)

			t.Cleanup(func() {
				out.Body.Close()
				require.NoError(t, cli.Teardown())
			})
		})
	}
}

func checkResponse(t *testing.T, dbs []byte, out *http.Response) {
	t.Helper()

	outbs, err := httputil.DumpResponse(out, true)
	require.NoError(t, err)

	resp, err := http.ReadResponse(bufio.NewReader(bytes.NewBuffer(dbs)), nil)
	require.NoError(t, err)

	outDf, err := httputil.DumpResponse(resp, true)
	require.NoError(t, err)

	require.Equal(t, string(outDf), string(outbs))

	t.Cleanup(func() {
		resp.Body.Close()
	})
}
