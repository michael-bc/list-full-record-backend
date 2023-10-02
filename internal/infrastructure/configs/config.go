package configs

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	Environment    Environment `env:"ENVIRONMENT" envDefault:"development"`
	ServiceName    string      `env:"SERVICE_NAME" envDefault:"list-full-record"`
	HTTPServerAddr string      `env:"HTTP_SERVER_ADDR" envDefault:"0.0.0.0:8080"`
	LogLevel       string      `env:"LOG_LEVEL" envDefault:"debug"`
	DbURI          string      `env:"DB_URI" envDefault:"postgresql://postgres:postgres@localhost:5432/list-full-record?sslmode=disable"`
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load()

	var config Config
	if err := env.Parse(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
