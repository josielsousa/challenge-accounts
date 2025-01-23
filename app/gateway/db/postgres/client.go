package postgres

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
)

type Client struct {
	Pool *pgxpool.Pool
}

func NewClient(ctx context.Context, dbURL string) (*Client, error) {
	pool, err := ConnectPoolWithMigrations(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("on connect pgxpool with migrations: %w", err)
	}

	return &Client{
		Pool: pool,
	}, nil
}
