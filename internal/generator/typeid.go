package generator

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"go.jetify.com/typeid"
)

// TypeID generator.
type TypeID struct{}

// Generate a prefixless TypeID.
func (t *TypeID) Generate(_ context.Context, _ *Application) string {
	return typeid.Must(typeid.WithPrefix("")).String()
}
