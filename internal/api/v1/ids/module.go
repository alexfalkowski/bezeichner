package ids

import (
	"github.com/alexfalkowski/bezeichner/internal/api/ids"
	"github.com/alexfalkowski/go-service/v2/di"
)

// Module for fx.
var Module = di.Module(
	ids.Module,
	di.Constructor(NewIdentifier),
)
