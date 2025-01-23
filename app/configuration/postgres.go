package configuration

import (
	"fmt"
)

type PostgresConfig struct {
	DatabaseName string `envconfig:"DATABASE_NAME"          default:"dev"`
	User         string `envconfig:"DATABASE_USER"          default:"postgres"`
	Password     string `envconfig:"DATABASE_PASSWORD"      default:"postgres"`
	Host         string `envconfig:"DATABASE_HOST_DIRECT"   default:"localhost"`
	Port         string `envconfig:"DATABASE_PORT_DIRECT"   default:"5432"`
	PoolMaxSize  string `envconfig:"DATABASE_POOL_MAX_SIZE" default:"5"`
	PoolMinSize  string `envconfig:"DATABASE_POOL_MIN_SIZE" default:"1"`
	Hostname     string `envconfig:"HOSTNAME"`
}

func (p PostgresConfig) URL() string {
	connConfig := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		p.User,
		p.Password,
		p.Host,
		p.Port,
		p.DatabaseName,
	)

	connConfig = fmt.Sprintf("%s pool_max_conns=%s", connConfig, p.PoolMaxSize)
	connConfig = fmt.Sprintf("%s pool_min_conns=%s", connConfig, p.PoolMinSize)

	if p.Hostname != "" {
		connConfig = fmt.Sprintf("%s&application_name=%s", connConfig, p.Hostname)
	}

	return connConfig
}
