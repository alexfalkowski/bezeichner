package generator

import (
	"context"
	"crypto/rand"
	"time"

	"github.com/oklog/ulid"
)

// ULID generator.
type ULID struct{}

// Generate a ULID.
func (k *ULID) Generate(_ context.Context, app *Application) (string, error) {
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, rand.Reader)

	return app.ID(id.String()), err
}
