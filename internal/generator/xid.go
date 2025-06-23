package generator

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/rs/xid"
)

// XID generator.
type XID struct{}

// Generate an XID.
func (x *XID) Generate(_ context.Context, app *Application) (string, error) {
	id := xid.New()

	return app.ID(id.String()), nil
}
