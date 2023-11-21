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
func (r *Redis) Generate(ctx context.Context, app *Application) (string, error) {
	c := r.client.Incr(ctx, app.Name)
	if err := c.Err(); err != nil {
		return "", err
	}

	res, err := c.Result()
	if err != nil {
		return "", err
	}

	return app.ID(strconv.FormatInt(res, 10)), nil
}
