package v1

import (
	"github.com/alexfalkowski/bezeichner/internal/api/v1/ids"
	"github.com/alexfalkowski/bezeichner/internal/api/v1/transport/grpc"
	"github.com/alexfalkowski/bezeichner/internal/api/v1/transport/http"
	"github.com/alexfalkowski/go-signal/v2"
)

// Module registers the v1 identifier service with the gRPC and HTTP transports.
var Module = signal.Module(
	ids.Module,
	signal.Constructor(grpc.NewServer),
	signal.Constructor(http.NewServer),
	signal.Register(grpc.Register),
	signal.Register(http.Register),
)
