package generator

import (
	"context"

	"go.jetpack.io/typeid"
)

// TypeID generator.
type TypeID struct{}

// Generate a TypeID.
func (t *TypeID) Generate(_ context.Context, _ string) (string, error) {
	id, err := typeid.WithPrefix("")
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
