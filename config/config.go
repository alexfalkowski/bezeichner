package config

import (
	"github.com/alexfalkowski/bezeichner/client"
	"github.com/alexfalkowski/bezeichner/generator"
	"github.com/alexfalkowski/bezeichner/health"
	"github.com/alexfalkowski/bezeichner/mapper"
	"github.com/alexfalkowski/go-service/config"
)

// Config for the service.
type Config struct {
	Client        *client.Config    `yaml:"client,omitempty" json:"client,omitempty" toml:"client,omitempty"`
	Health        *health.Config    `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty"`
	Generator     *generator.Config `yaml:"generator,omitempty" json:"generator,omitempty" toml:"generator,omitempty"`
	Mapper        *mapper.Config    `yaml:"mapper,omitempty" json:"mapper,omitempty" toml:"mapper,omitempty"`
	config.Config `yaml:",inline" json:",inline" toml:",inline"`
}
