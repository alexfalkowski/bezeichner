package config

import (
	"github.com/alexfalkowski/bezeichner/generator"
	"github.com/alexfalkowski/bezeichner/health"
	"github.com/alexfalkowski/bezeichner/mapper"
	"github.com/alexfalkowski/go-service/config"
)

// Config for the service.
type Config struct {
	Health         *health.Config    `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty"`
	Generator      *generator.Config `yaml:"generator,omitempty" json:"generator,omitempty" toml:"generator,omitempty"`
	Mapper         *mapper.Config    `yaml:"mapper,omitempty" json:"mapper,omitempty" toml:"mapper,omitempty"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

// Valid or error.
func (c Config) Valid() error {
	if c.Generator == nil || c.Config == nil {
		return config.ErrInvalidConfig
	}

	return c.Config.Valid()
}

func decorateConfig(cfg *Config) *config.Config {
	return cfg.Config
}

func healthConfig(cfg *Config) *health.Config {
	return cfg.Health
}

func generatorConfig(cfg *Config) *generator.Config {
	return cfg.Generator
}

func mapperConfig(cfg *Config) *mapper.Config {
	return cfg.Mapper
}
