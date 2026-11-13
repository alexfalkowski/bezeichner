package generator

import "github.com/alexfalkowski/go-signal/v2"

// Module provides the configured identifier generators.
var Module = signal.Module(signal.Constructor(NewGenerators))
