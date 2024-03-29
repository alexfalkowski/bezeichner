package generator

import (
	"context"

	"github.com/jaevor/go-nanoid"
)

// NanoID generator.
type NanoID struct{}

// Generate a NanoID.
func (n *NanoID) Generate(_ context.Context, app *Application) (string, error) {
	id, err := nanoid.Standard(21)
	if err != nil {
		return "", err
	}

	return app.ID(id()), nil
}
