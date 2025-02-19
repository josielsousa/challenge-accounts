package cache

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/josielsousa/challenge-accounts/app/configuration"
)

type Cache struct {
	C *redis.Client
}

func NewCache(ctx context.Context, cfg configuration.CacheConfig) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("on connecting to cache: %w", err)
	}

	return &Cache{C: client}, nil
}

func (c *Cache) Close() error {
	if err := c.C.Close(); err != nil {
		return fmt.Errorf("closing cache: %w", err)
	}

	return nil
}

func (c *Cache) SetIdempotency(
	ctx context.Context,
	key string,
	response *http.Response,
	ttl time.Duration,
) error {
	bs, err := httputil.DumpResponse(response, true)
	if err != nil {
		return fmt.Errorf("dumping response: %w", err)
	}

	if err := c.C.Set(ctx, key, string(bs), ttl).Err(); err != nil {
		return fmt.Errorf("setting cache: %w", err)
	}

	return nil
}

func (c *Cache) GetIdempotency(ctx context.Context, key string) (*http.Response, error) {
	cac, err := c.C.Get(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("getting cache: %w", err)
	}

	bsReader := bytes.NewReader([]byte(cac))

	resp, err := http.ReadResponse(bufio.NewReader(bsReader), nil)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	return resp, nil
}
