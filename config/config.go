package config

import (
	"github.com/alexfalkowski/bezeichner/generator"
	"github.com/alexfalkowski/bezeichner/health"
	"github.com/alexfalkowski/go-service/config"
)

// Config for the service.
type Config struct {
	Health        health.Config    `yaml:"health"`
	Generator     generator.Config `yaml:"generator"`
	config.Config `yaml:",inline"`
}
