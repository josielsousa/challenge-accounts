package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
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
		return nil, fmt.Errorf("on pgx pool parse config: %w", err)
	}

	// TODO: add trace
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("error connecting to pgx pool: %w", err)
	}

	if runMigrations {
		err = RunMigrationsConn(dbURL)
		if err != nil {
			return nil, fmt.Errorf("on run migrations: %w", err)
		}
	}

	return pool, err
}

func RunMigrationsConn(dbURL string) error {
	migHandler, err := GetMigrationHandler(dbURL)
	if err != nil {
		return fmt.Errorf("on get migration handler: %w", err)
	}

	defer migHandler.Close()

	if err := migHandler.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("on up migration version: %w", err)
		}
	}

	return nil
}
