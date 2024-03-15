package mapper

// Identifiers to map.
type Identifiers map[string]string

// Config for mapper.
type Config struct {
	Identifiers Identifiers `yaml:"identifiers,omitempty" json:"identifiers,omitempty" toml:"identifiers,omitempty"`
}
