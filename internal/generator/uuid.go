package generator

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/google/uuid"
)

// UUID generator.
type UUID struct{}

// Generate a UUID.
func (g *UUID) Generate(_ context.Context, app *Application) (string, error) {
	id, err := uuid.NewRandom()

	return app.ID(id.String()), err
}
