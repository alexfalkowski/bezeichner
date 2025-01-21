package generator

import (
	"context"

	nanoid "github.com/matoous/go-nanoid"
)

// NanoID generator.
type NanoID struct{}

// Generate a NanoID.
func (n *NanoID) Generate(_ context.Context, app *Application) (string, error) {
	id, err := nanoid.Nanoid()

	return app.ID(id), err
}
