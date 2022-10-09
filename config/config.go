package config

import (
	"github.com/alexfalkowski/bezeichner/generator"
	"github.com/alexfalkowski/bezeichner/health"
	"github.com/alexfalkowski/bezeichner/mapper"
	"github.com/alexfalkowski/go-service/config"
)

// Config for the service.
type Config struct {
	Health        health.Config    `yaml:"health"`
	Generator     generator.Config `yaml:"generator"`
	Mapper        mapper.Config    `yaml:"mapper"`
	config.Config `yaml:",inline"`
}
