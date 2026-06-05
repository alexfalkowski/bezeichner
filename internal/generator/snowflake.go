package generator

import (
	"strconv"

	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/runtime"
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

// Generate a Snowflake ID.
//
// If Sonyflake's time range is exhausted, this method panics via [runtime.Must].
func (s *Snowflake) Generate(_ context.Context, _ *Application) string {
	id, err := s.sf.NextID()
	runtime.Must(err)

	return strconv.FormatUint(id, 10)
}
