package config

import (
	"time"

	"github.com/mnsavag/fast-todo-app.git/internal/lib/log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Service  `envPrefix:"SERVICE_"`
	Database `envPrefix:"DATABASE_"`

	Logger log.Config
}

type Service struct {
	Host                   string        `env:"HOST" envDefault:"127.0.0.1"`
	HttpPort               string        `env:"HTTP_PORT" envDefault:"8080"`
	GrpcPort               string        `env:"GRPC_PORT" envDefault:"8081"`
	ShutdownContextTimeout time.Duration `env:"SHUTDOWN_CONTEXT_TIMEOUT_DURATION" envDefault:"5s"`
}

type Database struct {
	DSN string `env:"DSN"`
}

type Logger struct {
}

func NewConfig() (Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
