package generator

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/strings"
	nanoid "github.com/matoous/go-nanoid"
)

// NanoID generator.
type NanoID struct{}

// Generate a NanoID.
func (n *NanoID) Generate(_ context.Context, app *Application) (string, error) {
	id, err := nanoid.Nanoid()
	if err != nil {
		return strings.Empty, err
	}

	return app.ID(id), nil
}
