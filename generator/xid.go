package generator

import (
	"context"

	"github.com/rs/xid"
)

// XID generator.
type XID struct{}

// Generate an XID.
func (x *XID) Generate(_ context.Context, _ string) (string, error) {
	id := xid.New()

	return id.String(), nil
}
