package limits

const (
	defaultGenerateCount = 1000
	defaultMapIDs        = 1000
)

// Config configures per-request item-count limits.
type Config struct {
	GenerateCount uint64 `yaml:"generate_count,omitempty" json:"generate_count,omitempty" toml:"generate_count,omitempty"`
	MapIDs        uint64 `yaml:"map_ids,omitempty" json:"map_ids,omitempty" toml:"map_ids,omitempty"`
}

// MaxGenerateCount returns the effective GenerateIdentifiers count limit.
func (c *Config) MaxGenerateCount() uint64 {
	if c == nil || c.GenerateCount == 0 {
		return defaultGenerateCount
	}

	return c.GenerateCount
}

// MaxMapIDs returns the effective MapIdentifiers ids limit.
func (c *Config) MaxMapIDs() uint64 {
	if c == nil || c.MapIDs == 0 {
		return defaultMapIDs
	}

	return c.MapIDs
}
