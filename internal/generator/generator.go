package generator

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/crypto/rand"
	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/linxGnu/mssqlx"
)

// ErrNotFound for generator.
var ErrNotFound = errors.New("generator not found")

// NewGenerators of identifiers.
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

// Generators of identifiers.
type Generators map[string]Generator

// Generator from kind.
func (gs Generators) Generator(kind string) (Generator, error) {
	if g, ok := gs[kind]; ok {
		return g, nil
	}

	return nil, ErrNotFound
}

// Generator to generate an identifier.
type Generator interface {
	// Generate an identifier.
	Generate(ctx context.Context, app *Application) (string, error)
}
