package generator

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/strings"
	"github.com/google/uuid"
)

// UUID generator.
type UUID struct{}

// Generate a UUID.
func (g *UUID) Generate(_ context.Context, app *Application) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return strings.Empty, err
	}

	return app.ID(id.String()), nil
}
