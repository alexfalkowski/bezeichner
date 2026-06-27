package generator

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"go.jetify.com/typeid"
)

// TypeID generates native TypeIDs using the application name as the TypeID
// prefix.
//
// Application names used with this generator must satisfy TypeID prefix syntax:
// lowercase ASCII letters and underscores only, at most 63 characters, and no
// leading or trailing underscore.
type TypeID struct{}

// Generate returns a TypeID prefixed with the application name.
//
// It panics if app.Name violates the TypeID prefix contract.
func (t *TypeID) Generate(_ context.Context, app *Application) string {
	return typeid.Must(typeid.WithPrefix(app.Name)).String()
}
