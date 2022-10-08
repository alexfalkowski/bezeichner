package generator

import (
	"context"
	"crypto/rand"
	"time"

	"github.com/oklog/ulid"
)

// ULID generator.
type ULID struct{}

// Generate a ULID.
func (k *ULID) Generate(ctx context.Context) (string, error) {
	ms := ulid.Timestamp(time.Now())

	id, err := ulid.New(ms, rand.Reader)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
