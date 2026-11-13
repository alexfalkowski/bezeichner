package ids

import "github.com/alexfalkowski/go-signal/v2"

// Module provides the transport-independent identifier service.
var Module = signal.Module(
	signal.Constructor(NewIdentifier),
)
