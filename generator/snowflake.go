package generator

import (
	"context"
	"strconv"

	"github.com/godruoyi/go-snowflake"
)

// Snowflake generator.
type Snowflake struct{}

// Generate a id with snowflake.
func (s *Snowflake) Generate(ctx context.Context) (string, error) {
	id, err := snowflake.NextID()
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(int64(id), 10), nil
}
