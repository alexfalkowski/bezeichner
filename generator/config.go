package generator

import (
	"fmt"
)

// Application fto generate identifiers.
type Application struct {
	Name      string `yaml:"name" json:"name" toml:"name"`
	Kind      string `yaml:"kind" json:"kind" toml:"kind"`
	Prefix    string `yaml:"prefix" json:"prefix" toml:"prefix"`
	Suffix    string `yaml:"suffix" json:"suffix" toml:"suffix"`
	Separator string `yaml:"separator" json:"separator" toml:"separator"`
}

// ID for the application.
func (a *Application) ID(id string) string {
	return fmt.Sprintf("%s%s%s%s%s", a.Prefix, a.Separator, id, a.Separator, a.Suffix)
}

// Config for generator.
type Config struct {
	Applications []Application `yaml:"applications"`
}

// Application by name.
func (c *Config) Application(name string) *Application {
	for _, d := range c.Applications {
		if d.Name == name {
			return &d
		}
	}

	return nil
}
