package generator

import (
	"context"

	"github.com/segmentio/ksuid"
)

// KSUID generator.
type KSUID struct{}

// Generate a KSUID.
func (k *KSUID) Generate(ctx context.Context) (string, error) {
	id, err := ksuid.NewRandom()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
