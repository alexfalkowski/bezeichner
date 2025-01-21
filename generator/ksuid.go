package generator

import (
	"context"

	"github.com/segmentio/ksuid"
)

// KSUID generator.
type KSUID struct{}

// Generate a KSUID.
func (k *KSUID) Generate(_ context.Context, app *Application) (string, error) {
	id, err := ksuid.NewRandom()

	return app.ID(id.String()), err
}
