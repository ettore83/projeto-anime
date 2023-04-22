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
	AWS struct {
		Region string `envconfig:"AWS_REGION" default:":us-west-2"`
	}
	DynamoDB struct {
		Host   string `envconfig:"AWS_DYNAMODB_HOST" default:"http://anime-db:8000"`
		Tables struct {
			Anime string `envconfig:"AWS_DYNAMODB_ANIME_TABLE" default:"anime"`
		}
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
