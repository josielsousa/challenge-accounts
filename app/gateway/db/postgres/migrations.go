package postgres

import (
	"embed"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/httpfs"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

//go:embed migrations
var migrations embed.FS

func GetMigrationHandler(dbUrl string) (*migrate.Migrate, error) {
	source, err := httpfs.New(http.FS(migrations), "migrations")
	if err != nil {
		return nil, err
	}

	return migrate.NewWithSourceInstance("httpfs", source, dbUrl)
}

func RunMigrations(dbUrl string) error {
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
