package config

import (
	v1c "github.com/alexfalkowski/bezeichner/client/v1/config"
	"github.com/alexfalkowski/bezeichner/generator"
	"github.com/alexfalkowski/bezeichner/health"
	"github.com/alexfalkowski/bezeichner/mapper"
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/config"
)

// NewConfigurator for config.
func NewConfigurator(i *cmd.InputConfig) (config.Configurator, error) {
	c := &Config{}

	return c, i.Unmarshal(c)
}

func healthConfig(cfg config.Configurator) *health.Config {
	return cfg.(*Config).Health
}

func generatorConfig(cfg config.Configurator) *generator.Config {
	return cfg.(*Config).Generator
}

func mapperConfig(cfg config.Configurator) *mapper.Config {
	return cfg.(*Config).Mapper
}

func v1ClientConfig(cfg config.Configurator) *v1c.Config {
	return cfg.(*Config).Client.V1
}
