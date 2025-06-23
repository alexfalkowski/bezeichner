package generator

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"go.jetify.com/typeid"
)

// TypeID generator.
type TypeID struct{}

// Generate a TypeID.
func (t *TypeID) Generate(_ context.Context, app *Application) (string, error) {
	id, err := typeid.From(app.Prefix, app.Suffix)

	return id.String(), err
}
