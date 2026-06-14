package generator

import (
	"strconv"

	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/runtime"
	"github.com/sony/sonyflake"
)

// Snowflake generates Sonyflake IDs using Sonyflake default settings.
//
// The defaults derive the machine ID from a private IPv4 address.
type Snowflake struct {
	sf *sonyflake.Sonyflake
}

// NewSnowflake constructs a Snowflake generator with Sonyflake default settings.
//
// Sonyflake's constructor can return nil when the default machine ID cannot be
// derived, such as in environments without a private IPv4 address. In that
// case, Generate will panic when it attempts to use the nil Sonyflake instance.
func NewSnowflake() *Snowflake {
	return &Snowflake{sf: sonyflake.NewSonyflake(sonyflake.Settings{})}
}

// Generate a Snowflake ID.
//
// Generate panics if Sonyflake construction failed or if Sonyflake's time range
// is exhausted.
func (s *Snowflake) Generate(_ context.Context, _ *Application) string {
	id, err := s.sf.NextID()
	runtime.Must(err)

	return strconv.FormatUint(id, 10)
}
