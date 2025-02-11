package config

import (
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	defaultPort = 8080
)

type Config struct {
	HTTPPort   int    `env:"KS_HTTP_PORT"`
	DBHost     string `env:"KS_DB_HOST"`
	DBPort     int    `env:"KS_DB_PORT"`
	DBUser     string `env:"KS_DB_USER"`
	DBPassword string `env:"KS_DB_PASSWORD"`
	DBName     string `env:"KS_DB_NAME"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(".env", cfg)
	if err != nil {
		_ = cleanenv.ReadEnv(cfg)
		slog.Info(err.Error())
	}

	cfg.validate()

	return cfg, nil
}

func (c *Config) validate() {
	if c.HTTPPort == 0 {
		c.HTTPPort = defaultPort
	}
}
