package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
)

const (
	defaultPort = 8080
)

type Config struct {
	HttpPort   int    `env:"KS_HTTP_PORT"`
	DbHost     string `env:"KS_DB_HOST"`
	DbPort     int    `env:"KS_DB_PORT"`
	DbUser     string `env:"KS_DB_USER"`
	DbPassword string `env:"KS_DB_PASSWORD"`
	DbName     string `env:"KS_DB_NAME"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(".env", cfg)
	if err != nil {
		_ = cleanenv.ReadEnv(cfg)
		slog.Info(err.Error())
	}

	err = cfg.validate()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validate() error {
	switch {
	case c.HttpPort == 0:
		c.HttpPort = defaultPort
	}

	return nil
}
