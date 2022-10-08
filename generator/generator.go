package generator

import (
	"context"
	"errors"

	"github.com/linxGnu/mssqlx"
)

// ErrNotFound for generator.
var ErrNotFound = errors.New("generator not found")

// Generator to generate an identifier.
type Generator interface {
	// Generate an identifier.
	Generate(ctx context.Context) (string, error)
}

// NewGenerator from kind.
func NewGenerator(name, kind string, db *mssqlx.DBs) (Generator, error) {
	switch kind {
	case "uuid":
		return &UUID{}, nil
	case "ksuid":
		return &KSUID{}, nil
	case "ulid":
		return &ULID{}, nil
	case "pg":
		return &PG{name: name, db: db}, nil
	}

	return nil, ErrNotFound
}
