package config

import (
	"github.com/alexfalkowski/bezeichner/generator"
	"github.com/alexfalkowski/bezeichner/health"
	"github.com/alexfalkowski/bezeichner/mapper"
	"github.com/alexfalkowski/go-service/config"
)

// NewConfigurator for config.
func NewConfigurator() config.Configurator {
	cfg := &Config{}

	return cfg
}

func healthConfig(cfg config.Configurator) *health.Config {
	return &cfg.(*Config).Health
}

func generatorConfig(cfg config.Configurator) *generator.Config {
	return &cfg.(*Config).Generator
}

func mapperConfig(cfg config.Configurator) *mapper.Config {
	return &cfg.(*Config).Mapper
}
