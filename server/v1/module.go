package v1

import (
	"github.com/alexfalkowski/bezeichner/server/service"
	"github.com/alexfalkowski/bezeichner/server/v1/transport/grpc"
	"github.com/alexfalkowski/bezeichner/server/v1/transport/http"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	service.Module,
	fx.Provide(grpc.NewServer),
	fx.Invoke(grpc.Register),
	fx.Invoke(http.Register),
)
