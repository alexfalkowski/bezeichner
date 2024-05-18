package generator

import (
	"fmt"
)

// Application fto generate identifiers.
type Application struct {
	Name      string `yaml:"name,omitempty" json:"name,omitempty" toml:"name,omitempty"`
	Kind      string `yaml:"kind,omitempty" json:"kind,omitempty" toml:"kind,omitempty"`
	Prefix    string `yaml:"prefix,omitempty" json:"prefix,omitempty" toml:"prefix,omitempty"`
	Suffix    string `yaml:"suffix,omitempty" json:"suffix,omitempty" toml:"suffix,omitempty"`
	Separator string `yaml:"separator,omitempty" json:"separator,omitempty" toml:"separator,omitempty"`
}

// ID for the application.
func (a *Application) ID(id string) string {
	return fmt.Sprintf("%s%s%s%s%s", a.Prefix, a.Separator, id, a.Separator, a.Suffix)
}

// Config for generator.
type Config struct {
	Applications []*Application `yaml:"applications,omitempty" json:"applications,omitempty" toml:"applications,omitempty"`
}

// Application by name.
func (c *Config) Application(name string) *Application {
	for _, d := range c.Applications {
		if d.Name == name {
			return d
		}
	}

	return nil
}
