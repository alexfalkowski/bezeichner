package client

import (
	v1 "github.com/alexfalkowski/bezeichner/client/v1/config"
)

// Config for client.
type Config struct {
	V1 v1.Config `yaml:"v1" json:"v1" toml:"v1"`
}