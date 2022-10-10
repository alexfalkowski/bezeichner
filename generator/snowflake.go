package generator

import (
	"context"
	"strconv"

	"github.com/sony/sonyflake"
)

// Snowflake generator.
type Snowflake struct{}

// Generate a id with snowflake.
func (s *Snowflake) Generate(ctx context.Context) (string, error) {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})

	id, err := sf.NextID()
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(int64(id), 10), nil
}
