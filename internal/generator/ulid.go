package generator

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/crypto/rand"
	"github.com/alexfalkowski/go-service/v2/time"
	"github.com/oklog/ulid"
)

// ULID generator.
type ULID struct {
	generator *rand.Generator
}

// Generate a ULID.
func (k *ULID) Generate(_ context.Context, app *Application) (string, error) {
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, k.generator)

	return app.ID(id.String()), err
}
