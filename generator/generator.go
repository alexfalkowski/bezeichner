package generator

import (
	"context"
	"errors"
)

// ErrNotFound for generator.
var ErrNotFound = errors.New("generator not found")

// Generator to generate an identifier.
type Generator interface {
	// Generate an identifier.
	Generate(ctx context.Context) (string, error)
}

// NewGenerator from kind.
func NewGenerator(kind string) (Generator, error) {
	switch kind {
	case "uuid":
		return &UUID{}, nil
	case "ksuid":
		return &KSUID{}, nil
	}

	return nil, ErrNotFound
}
