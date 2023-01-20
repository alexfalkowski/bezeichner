package client

import (
	v1 "github.com/alexfalkowski/bezeichner/client/v1"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		v1.Module,
		fx.Provide(NewClient),
	)
)
