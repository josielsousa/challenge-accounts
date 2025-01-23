package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPoolWithMigrations(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	return connectPool(ctx, dbURL, true)
}

func ConnectPoolWithoutMigrations(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	return connectPool(ctx, dbURL, false)
}

func connectPool(ctx context.Context, dbURL string, runMigrations bool) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	// TODO: add trace
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("error connecting to pgx pool: %w", err)
	}

	if runMigrations {
		if err := RunMigrations(config.ConnConfig); err != nil {
			return nil, fmt.Errorf("on run migrations: %w", err)
		}
	}

	return pool, nil
}
