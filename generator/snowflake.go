package generator

import (
	"context"
	"strconv"

	"github.com/sony/sonyflake"
)

// Snowflake generator.
type Snowflake struct {
	sf *sonyflake.Sonyflake
}

// NewSnowflake generator.
func NewSnowflake() *Snowflake {
	return &Snowflake{sf: sonyflake.NewSonyflake(sonyflake.Settings{})}
}

// Generate a id with snowflake.
func (s *Snowflake) Generate(_ context.Context, app *Application) (string, error) {
	id, err := s.sf.NextID()
	if err != nil {
		return "", err
	}

	return app.ID(strconv.FormatInt(int64(id), 10)), nil //nolint:gosec
}
