package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	App        App
	Log        Log
	GRPC       GRPC
	HTTP       HTTP
	Metrics    Metrics
	PostgreSQL PostgreSQL
	TokenProvider
}

type App struct {
	Name    string `env:"APP_NAME,required"`
	Version string `env:"APP_VERSION,required"`
}

type Log struct {
	Level string `env:"LOG_LEVEL,required"`
}

type GRPC struct {
	Port string `env:"GRPC_PORT,required"`
}

type HTTP struct {
	Port string `env:"HTTP_PORT,required"`
}

type Metrics struct {
	Enabled bool `env:"METRICS_ENABLED,required"`
}

type PostgreSQL struct {
	User       string `env:"DB_USER,required"`
	Password   string `env:"DB_PASSWORD,required"`
	Host       string `env:"DB_HOST,required"`
	Port       string `env:"DB_PORT,required"`
	Name       string `env:"DB_NAME,required"`
	SslEnabled bool   `env:"DB_SSL_ENABLED,required"`
	TxMarker   string `env:"DB_TX_MARKER,required"`
}

type TokenProvider struct {
	SecretKey string `env:"TOKEN_PROVIDER_SECRET_KEY,required"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	return cfg, nil
}
