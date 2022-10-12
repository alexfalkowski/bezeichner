package generator

import (
	"context"

	"github.com/rs/xid"
)

// XID generator.
type XID struct{}

// Generate an XID.
func (x *XID) Generate(ctx context.Context, name string) (string, error) {
	id := xid.New()

	return id.String(), nil
}
