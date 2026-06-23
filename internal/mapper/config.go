package mapper

// Identifiers to map.
type Identifiers map[string]string

// Application describes how to map identifiers for a named application.
type Application struct {
	Identifiers Identifiers `yaml:"identifiers,omitempty" json:"identifiers,omitempty" toml:"identifiers,omitempty"`
	Name        string      `yaml:"name" json:"name" toml:"name" validate:"required"`
}

// Config for mapper.
type Config struct {
	Applications []*Application `yaml:"applications" json:"applications" toml:"applications" validate:"unique=Name,dive,required"`
}

// Application returns the configured Application with the given name.
//
// It returns nil if no application exists with that name.
func (c *Config) Application(name string) *Application {
	for _, d := range c.Applications {
		if d.Name == name {
			return d
		}
	}

	return nil
}
