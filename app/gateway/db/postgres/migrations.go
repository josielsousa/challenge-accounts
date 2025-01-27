package postgres

import (
	"embed"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

//go:embed migrations
var MigrationsFS embed.FS

var _ migrate.Logger = &migrateLogger{}

type migrateLogger struct {
	logger  *slog.Logger
	verbose bool
}

func (l *migrateLogger) Printf(msg string, args ...any) {
	msg = fmt.Sprintf(msg, args...)
	msg = strings.ReplaceAll(msg, "/u", "")
	msg = strings.ReplaceAll(msg, "\n", "")

	l.logger.Info(msg)
}

func (l *migrateLogger) Verbose() bool {
	return l.verbose
}

func newMigrateHandler(connConfig *pgx.ConnConfig) (*migrate.Migrate, error) {
	source, err := httpfs.New(http.FS(MigrationsFS), "migrations")
	if err != nil {
		return nil, fmt.Errorf("on create source instance: %w", err)
	}

	driver, err := postgres.WithInstance(
		stdlib.OpenDB(*connConfig), &postgres.Config{
			DatabaseName: connConfig.Database,
		},
	)
	if err != nil {
		return nil, fmt.Errorf(
			"[migrations] failed to get postgres driver: %w",
			err,
		)
	}

	mig, err := migrate.NewWithInstance(
		"httpfs", source, connConfig.Database, driver,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"[migrations] failed to create migrate source instance: %w",
			err,
		)
	}

	return mig, nil
}

func runMigrations(connConfig *pgx.ConnConfig) error {
	migHandler, err := newMigrateHandler(connConfig)
	if err != nil {
		return err
	}

	defer migHandler.Close()

	if err := migHandler.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("on run migrations: %w", err)
		}
	}

	return nil
}
