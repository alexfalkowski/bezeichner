package generator

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/id"
)

// NewID adapts the go-service ID generator registered for kind.
func NewID(ids *id.Map, kind string) *ID {
	return &ID{generator: ids.Get(kind)}
}

// ID adapts a go-service ID generator to the Bezeichner generator interface.
type ID struct {
	generator id.Generator
}

// Generate an ID.
func (i *ID) Generate(_ context.Context, _ *Application) string {
	return i.generator.Generate()
}
