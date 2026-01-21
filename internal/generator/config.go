package generator

// Application fto generate identifiers.
type Application struct {
	Name string `yaml:"name,omitempty" json:"name,omitempty" toml:"name,omitempty"`
	Kind string `yaml:"kind,omitempty" json:"kind,omitempty" toml:"kind,omitempty"`
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
