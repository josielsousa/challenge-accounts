package configuration

type Cache struct {
	Host string `envconfig:"CACHE_HOST" default:"localhost"`
	Port string `envconfig:"CACHE_PORT" default:"6379"`
}
