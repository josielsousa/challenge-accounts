package postgres

import (
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

func Connect(dbURL string, log *logrus.Entry) (*pgx.Conn, error) {
	config, err := pgx.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("on pgx pool parse config: %w", err)
	}

	if log != nil {
		logger := logrusadapter.NewLogger(log)
		config.Logger = logger
	}

	db, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("on pg pool connect config: %w", err)
	}

	err = runMigrations(dbURL)
	if err != nil {
		return nil, fmt.Errorf("on run migrations: %w", err)
	}

	return db, err
}

func ConnectPool(dbURL string, log *logrus.Entry, logLevel LogLevel) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("on pgx pool parse config: %w", err)
	}

	if log != nil {
		config.ConnConfig.Logger = logrusadapter.NewLogger(log)
	}

	config.ConnConfig.LogLevel = pgx.LogLevel(logLevel)

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("on pg pool connect config: %w", err)
	}

	err = runMigrations(dbURL)
	if err != nil {
		return nil, fmt.Errorf("on run migrations: %w", err)
	}

	return db, err
}

func runMigrations(dbUrl string) error {
	m, err := GetMigrationHandler(dbUrl)
	if err != nil {
		return fmt.Errorf("on get migration handler: %w", err)
	}

	defer m.Close()
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("on up migration version: %w", err)
	}

	return nil
}
