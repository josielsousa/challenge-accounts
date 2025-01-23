package postgres

import (
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Client struct {
	Pool *pgxpool.Pool
}

func NewClient(dbURL string) (*Client, error) {
	pool, err := ConnectPoolWithMigrations(dbURL)
	if err != nil {
		return nil, fmt.Errorf("on connect pgxpool with migrations: %w", err)
	}

	return &Client{
		Pool: pool,
	}, nil
}
