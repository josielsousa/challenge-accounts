package cachetest

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	t.Parallel()

	cli, err := NewClient()
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, cli.Teardown())
	})

	_, err = cli.Client.Set(
		context.Background(),
		"test_key",
		"test_value",
		time.Minute,
	).Result()
	require.NoError(t, err)

	val, err := cli.Client.Get(context.Background(), "test_key").Result()
	require.NoError(t, err)

	require.Equal(t, "test_value", val)
}
