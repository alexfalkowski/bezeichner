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
	Client        client.Config    `yaml:"client" json:"client" toml:"client"`
	Health        health.Config    `yaml:"health" json:"health" toml:"health"`
	Generator     generator.Config `yaml:"generator" json:"generator" toml:"generator"`
	Mapper        mapper.Config    `yaml:"mapper" json:"mapper" toml:"mapper"`
	config.Config `yaml:",inline" json:",inline" toml:",inline"`
}
