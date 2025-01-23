package v1

import (
	"github.com/alexfalkowski/bezeichner/api/ids"
	"github.com/alexfalkowski/bezeichner/api/v1/transport/grpc"
	"github.com/alexfalkowski/bezeichner/api/v1/transport/http"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	ids.Module,
	fx.Provide(grpc.NewServer),
	fx.Invoke(grpc.Register),
	fx.Provide(http.NewHandler),
	fx.Invoke(http.Register),
)
