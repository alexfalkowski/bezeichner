package generator

import (
	"context"

	"github.com/google/uuid"
)

// UUID generator.
type UUID struct{}

// Generate a UUID.
func (g *UUID) Generate(ctx context.Context, name string) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
