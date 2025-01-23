package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPoolWithMigrations(ctx context.Context, connConfig string) (*pgxpool.Pool, error) {
	return connectPool(ctx, connConfig, true)
}

func ConnectPoolWithoutMigrations(ctx context.Context, connConfig string) (*pgxpool.Pool, error) {
	return connectPool(ctx, connConfig, false)
}

func connectPool(ctx context.Context, connConfig string, runMigrations bool) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connConfig)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

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
