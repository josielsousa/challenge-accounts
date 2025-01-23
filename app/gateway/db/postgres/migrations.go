package postgres

import (
	"embed"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/httpfs"

	// needed to run migrations.
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
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

func GetMigrationHandler(dbURL string) (*migrate.Migrate, error) {
	source, err := httpfs.New(http.FS(MigrationsFS), "migrations")
	if err != nil {
		return nil, fmt.Errorf("on create source instance: %w", err)
	}

	mig, err := migrate.NewWithSourceInstance("httpfs", source, dbURL)
	if err != nil {
		return nil, fmt.Errorf("on create migrate instance: %w", err)
	}

	return mig, nil
}

func RunMigrations(dbURL string) error {
	migHandler, err := GetMigrationHandler(dbURL)
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
