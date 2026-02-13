package generator

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/crypto/rand"
	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/linxGnu/mssqlx"
)

// ErrNotFound indicates that a generator kind cannot be resolved from a registry.
var ErrNotFound = errors.New("generator not found")

// NewGenerators constructs the default generator registry.
//
// The returned registry maps a generator "kind" string (for example "uuid" or
// "ulid") to a concrete Generator implementation. It is used by the domain layer
// to select an implementation based on configured application kind.
func NewGenerators(db *mssqlx.DBs, generator *rand.Generator) Generators {
	return Generators{
		"uuid":      &UUID{},
		"ksuid":     &KSUID{},
		"ulid":      &ULID{generator: generator},
		"xid":       &XID{},
		"snowflake": NewSnowflake(),
		"nanoid":    &NanoID{},
		"typeid":    &TypeID{},
		"pg":        &PG{db: db},
	}
}

// Generators is a registry mapping a generator kind string to a Generator.
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
// Implementations may use the provided application configuration (for example,
// the Postgres generator uses app.Name to select a sequence), or may ignore it.
type Generator interface {
	// Generate produces a single identifier for the given application.
	Generate(ctx context.Context, app *Application) (string, error)
}
