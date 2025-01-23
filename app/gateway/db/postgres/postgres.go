package postgres

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type logger struct{}

func ConnectPoolWithMigrations(dbURL string) (*pgxpool.Pool, error) {
	return connectPool(dbURL, true)
}

func ConnectPoolWithoutMigrations(dbURL string) (*pgxpool.Pool, error) {
	return connectPool(dbURL, false)
}

func connectPool(dbURL string, runMigrations bool) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("on pgx pool parse config: %w", err)
	}

	config.ConnConfig.Logger = logger{}

	config.ConnConfig.LogLevel = pgx.LogLevel(pgx.LogLevelWarn)

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("on pg pool connect config: %w", err)
	}

	if runMigrations {
		err = RunMigrationsConn(dbURL)
		if err != nil {
			return nil, fmt.Errorf("on run migrations: %w", err)
		}
	}

	return db, err
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

func (l logger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	attrs := make([]slog.Attr, 0, len(data))
	for k, v := range data {
		attrs = append(attrs, slog.Any(k, v))
	}

	slog.LogAttrs(ctx, slog.Level(level), msg, attrs...)
}
