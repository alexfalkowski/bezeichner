package generator

import (
	"context"
	"errors"

	"github.com/alexfalkowski/go-service/cache/redis/client"
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
func NewGenerator(name, kind string, db *mssqlx.DBs, client client.Client) (Generator, error) {
	switch kind {
	case "uuid":
		return &UUID{}, nil
	case "ksuid":
		return &KSUID{}, nil
	case "ulid":
		return &ULID{}, nil
	case "pg":
		return &PG{name: name, db: db}, nil
	case "redis":
		return &Redis{name: name, client: client}, nil
	case "snowflake":
		return &Snowflake{}, nil
	}

	return nil, ErrNotFound
}
