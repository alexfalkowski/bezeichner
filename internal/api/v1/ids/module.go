package ids

import (
	"github.com/alexfalkowski/bezeichner/v2/internal/api/ids"
	"github.com/alexfalkowski/go-signal/v2"
)

// Module adapts the identifier service to the v1 API.
var Module = signal.Module(
	ids.Module,
	signal.Constructor(NewIdentifier),
)
