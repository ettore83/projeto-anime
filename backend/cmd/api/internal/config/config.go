package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	Http struct {
		Address      string `envconfig:"HTTP_ADDRESS" default:":8081"`
		WriteTimeout int    `envconfig:"HTTP__TIMEOUT" default:"5"`
		ReadTimeout  int    `envconfig:"HTTP_READ_TIMEOUT" default:"5"`
		IdleTimeout  int    `envconfig:"HTTP_IDLE_TIMEOUT" default:"5"`
	}
}

// NewConfig creates a new initialised application Config.
func NewConfig() (*Config, error) {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return nil, errors.WithStack(err)
	}

	return &config, nil
}
