package generator

import (
	"context"
	"errors"

	"github.com/alexfalkowski/go-service/redis"
	"github.com/linxGnu/mssqlx"
)

// ErrNotFound for generator.
var ErrNotFound = errors.New("generator not found")

// NewGenerators of identifiers.
func NewGenerators(db *mssqlx.DBs, client redis.Client) Generators {
	return Generators{
		"uuid":      &UUID{},
		"ksuid":     &KSUID{},
		"ulid":      &ULID{},
		"xid":       &XID{},
		"snowflake": NewSnowflake(),
		"nanoid":    &NanoID{},
		"typeid":    &TypeID{},
		"pg":        &PG{db: db},
		"redis":     &Redis{client: client},
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
	Generate(ctx context.Context, name string) (string, error)
}
