package generator

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"go.jetify.com/typeid"
)

// TypeID generator.
type TypeID struct{}

// Generate a TypeID prefixed with the application name.
func (t *TypeID) Generate(_ context.Context, app *Application) string {
	return typeid.Must(typeid.WithPrefix(app.Name)).String()
}
