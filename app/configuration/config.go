package configuration

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	API      APIConfig
	Postgres PostgresConfig
}

func LoadConfig() (*Config, error) {
	var config Config

	noPrefix := ""

	err := envconfig.Process(noPrefix, &config)
	if err != nil {
		return nil, fmt.Errorf("on loading config: %w", err)
	}

	return &config, nil
}

type APIConfig struct {
	Port    string `envconfig:"API_PORT" default:"3000"`
	Host    string `envconfig:"API_HOST" default:"0.0.0.0"`
	AppName string `envconfig:"APP_NAME" default:"challenge_accounts"`
}
