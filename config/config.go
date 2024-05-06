package config

import (
	"github.com/alexfalkowski/bezeichner/client"
	v1c "github.com/alexfalkowski/bezeichner/client/v1/config"
	"github.com/alexfalkowski/bezeichner/generator"
	"github.com/alexfalkowski/bezeichner/health"
	"github.com/alexfalkowski/bezeichner/mapper"
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/config"
)

// NewConfig for config.
func NewConfig(i *cmd.InputConfig) (*Config, error) {
	c := &Config{}

	return c, i.Unmarshal(c)
}

// Config for the service.
type Config struct {
	Client         *client.Config    `yaml:"client,omitempty" json:"client,omitempty" toml:"client,omitempty"`
	Health         *health.Config    `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty"`
	Generator      *generator.Config `yaml:"generator,omitempty" json:"generator,omitempty" toml:"generator,omitempty"`
	Mapper         *mapper.Config    `yaml:"mapper,omitempty" json:"mapper,omitempty" toml:"mapper,omitempty"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
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

func v1ClientConfig(cfg *Config) *v1c.Config {
	if !client.IsEnabled(cfg.Client) {
		return nil
	}

	return cfg.Client.V1
}
