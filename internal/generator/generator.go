package generator

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/alexfalkowski/go-service/v2/id"
)

// ErrNotFound indicates that a generator kind cannot be resolved from a registry.
var ErrNotFound = errors.New("generator not found")

// NewGenerators constructs the default generator registry.
//
// The returned registry maps a generator "kind" string (for example "uuid" or
// "ulid") to a concrete Generator implementation. It is used by the domain layer
// to select an implementation based on configured application kind.
func NewGenerators(ids *id.Map) Generators {
	return Generators{
		"uuid":      NewID(ids, "uuid"),
		"ksuid":     NewID(ids, "ksuid"),
		"ulid":      NewID(ids, "ulid"),
		"xid":       NewID(ids, "xid"),
		"snowflake": NewSnowflake(),
		"nanoid":    NewID(ids, "nanoid"),
		"typeid":    &TypeID{},
	}
}

// Generators is a registry mapping a generator kind string to a Generator.
//
// Treat a registry as read-only after startup. The service may resolve and call
// generators from concurrent request handlers.
type Generators map[string]Generator

// Generator returns the Generator registered for kind.
//
// It returns ErrNotFound if kind does not exist in the registry.
func (gs Generators) Generator(kind string) (Generator, error) {
	if g, ok := gs[kind]; ok {
		return g, nil
	}

	return nil, ErrNotFound
}

// Generator generates identifiers for a configured application.
//
// Implementations may use the provided application configuration or ignore it.
// Generate may be called concurrently by service request handlers; implementations
// with mutable state must synchronize access.
type Generator interface {
	// Generate produces a single identifier for the given application.
	Generate(ctx context.Context, app *Application) string
}
