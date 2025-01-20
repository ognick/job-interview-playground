package config

import (
	"fmt"

	"github.com/caarlos0/env/v7"

	"github.com/ognick/job-interview-playground/pkg/httpsrv"
	"github.com/ognick/job-interview-playground/pkg/logger"
)

type Config struct {
	HTTPAddress httpsrv.Addr  `env:"CONF_HTTP_ADDRESS" envDefault:":8080"`
	Logger      logger.Config `env:"CONF_LOGGER"`
}

func NewConfig() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("failed to parse env: %w", err)
	}

	return cfg, nil
}
