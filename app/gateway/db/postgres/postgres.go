package postgres

import (
	"context"

	"github.com/golang-migrate/migrate/v4"
	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

func Connect(dbURL string, log *logrus.Logger) (*pgx.Conn, error) {
	config, err := pgx.ParseConfig(dbURL)
	if err != nil {
		return nil, err
	}
	if log != nil {
		logger := logrusadapter.NewLogger(log)
		config.Logger = logger
	}
	db, err := pgx.ConnectConfig(context.Background(), config)
	return db, err
}

func ConnectPool(dbURL string, log *logrus.Logger, logLevel LogLevel) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, err
	}

	if log != nil {
		config.ConnConfig.Logger = logrusadapter.NewLogger(log)
	}

	config.ConnConfig.LogLevel = pgx.LogLevel(logLevel)

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	err = runMigrations(dbURL)
	if err != nil {
		return nil, err
	}

	return db, err
}

func runMigrations(dbUrl string) error {
	m, err := GetMigrationHandler(dbUrl)
	if err != nil {
		return err
	}

	defer m.Close()
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
