package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	App     App
	Log     Log
	GRPC    GRPC
	Swagger Swagger
	Metrics Metrics
}

type App struct {
	Name    string `env:"APP_NAME, required"`
	Version string `env:"APP_VERSION, required"`
}

type Log struct {
	Level string `env:"LOG_LEVEL, required"`
}

type GRPC struct {
	Port string `env:"GRPC_PORT, required"`
}

type Swagger struct {
	Enabled bool `env:"SWAGGER_ENABLED" envDefault:"false"`
}

type Metrics struct {
	Enabled bool `env:"METRICS_ENABLED" envDefault:"false"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
