package generator

import (
	"strconv"

	"github.com/alexfalkowski/go-service/v2/context"
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

	return app.ID(strconv.FormatUint(id, 10)), err
}
