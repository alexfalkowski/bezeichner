package generator

// Application describes how to generate identifiers for a named application.
//
// Name is used to select the application configuration from Config.
//
// Kind selects the generator implementation from a Generators registry (see
// NewGenerators). For example: "uuid", "ulid", or "pg".
type Application struct {
	Name string `yaml:"name,omitempty" json:"name,omitempty" toml:"name,omitempty"`
	Kind string `yaml:"kind,omitempty" json:"kind,omitempty" toml:"kind,omitempty"`
}

// Config contains the generator configuration.
//
// Applications is the set of generator applications available to the service.
// Applications are addressed by Application.Name.
type Config struct {
	Applications []*Application `yaml:"applications,omitempty" json:"applications,omitempty" toml:"applications,omitempty"`
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
