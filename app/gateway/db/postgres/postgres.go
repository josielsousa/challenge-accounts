package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool struct {
	connString        string
	executeMigrations bool
}

type PoolOption func(*Pool)

func WithConnString(connString string) PoolOption {
	return func(p *Pool) {
		p.connString = connString
	}
}

func WithoutMigrations() PoolOption {
	return func(p *Pool) {
		p.executeMigrations = false
	}
}

func NewPool(
	ctx context.Context, options ...PoolOption,
) (*pgxpool.Pool, error) {
	pool := &Pool{
		connString:        "",
		executeMigrations: true,
	}

	for _, opt := range options {
		opt(pool)
	}

	if strings.TrimSpace(pool.connString) == "" {
		return nil, errors.New("connection config is required")
	}

	return pool.connectPool(ctx)
}

func (p *Pool) connectPool(ctx context.Context) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(p.connString)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("new pool from config: %w", err)
	}

	if p.executeMigrations {
		if err := runMigrations(config.ConnConfig); err != nil {
			return nil, fmt.Errorf("on run migrations: %w", err)
		}
	}

	return pool, nil
}
