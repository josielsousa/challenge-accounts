package configuration

import (
	"fmt"
)

type PostgresConfig struct {
	DatabaseName string `envconfig:"DATABASE_NAME" default:"dev"`
	User         string `envconfig:"DATABASE_USER" default:"postgres"`
	Password     string `envconfig:"DATABASE_PASSWORD" default:"postgres"`
	Host         string `envconfig:"DATABASE_HOST_DIRECT" default:"localhost"`
	Port         string `envconfig:"DATABASE_PORT_DIRECT" default:"5432"`
	PoolMaxSize  string `envconfig:"DATABASE_POOL_MAX_SIZE" default:"10"`
	PoolMinSize  string `envconfig:"DATABASE_POOL_MIN_SIZE" default:"2"`
	Hostname     string `envconfig:"HOSTNAME"`
}

func (p PostgresConfig) DSN() string {
	dbConn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s pool_min_conns=%s pool_max_conns=%s",
		p.User, p.Password, p.Host, p.Port, p.DatabaseName, p.PoolMinSize, p.PoolMaxSize)

	if p.Hostname != "" {
		dbConn = fmt.Sprintf("%s application_name=%s", dbConn, p.Hostname)
	}

	return dbConn
}

func (p PostgresConfig) URL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.User, p.Password, p.Host, p.Port, p.DatabaseName)
}
