package generator

import (
	"context"
	"strconv"

	"github.com/alexfalkowski/go-service/redis"
)

// Redis generator.
type Redis struct {
	client redis.Client
}

// Generate an ID using INCR.
func (r *Redis) Generate(ctx context.Context, name string) (string, error) {
	c := r.client.Incr(ctx, name)
	if err := c.Err(); err != nil {
		return "", err
	}

	res, err := c.Result()
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(res, 10), nil
}
