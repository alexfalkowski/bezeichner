package generator

// Application describes how to generate identifiers for a named application.
//
// Name is used to select the application configuration from Config.
//
// Kind selects the generator implementation from a Generators registry (see
// NewGenerators). For example: "uuid" or "ulid".
type Application struct {
	Name string `yaml:"name" json:"name" toml:"name" validate:"required"`
	Kind string `yaml:"kind" json:"kind" toml:"kind" validate:"required"`
}

// Config contains the generator configuration.
//
// Applications is the set of generator applications available to the service.
// Applications are addressed by Application.Name.
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
