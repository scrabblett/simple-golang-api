package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	Env     string `env:"ENV" env-default:"development"`
	Storage string `env:"DATABASE_URL" env-default:"postgres://postgres:12345@localhost:5432/books_library?sslmode=disable"`
	Server  HTTPServer
}

type HTTPServer struct {
	Address     string        `env:"HTTP_SERVER_ADDRESS" env-default:"localhost"`
	Port        string        `env:"HTTP_SERVER_PORT" env-default:"8080"`
	Timeout     time.Duration `env:"HTTP_SERVER_TIMEOUT" env-default:"15s"`
	IdleTimeout time.Duration `env:"HTTP_SERVER_IDLE_TIMEOUT" env-default:"60s"`
}

func SetupConfig() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
