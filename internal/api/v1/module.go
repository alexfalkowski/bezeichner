package v1

import (
	"github.com/alexfalkowski/bezeichner/internal/api/ids"
	"github.com/alexfalkowski/bezeichner/internal/api/v1/transport/grpc"
	"github.com/alexfalkowski/bezeichner/internal/api/v1/transport/http"
	"github.com/alexfalkowski/go-service/v2/di"
)

// Module for fx.
var Module = di.Module(
	ids.Module,
	di.Constructor(grpc.NewServer),
	di.Register(grpc.Register),
	di.Register(http.Register),
)
