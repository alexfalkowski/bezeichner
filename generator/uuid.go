package generator

import (
	"context"

	"github.com/google/uuid"
)

// UUID generator.
type UUID struct{}

// Generate a UUID.
func (g *UUID) Generate(_ context.Context, app *Application) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return app.ID(id.String()), nil
}
