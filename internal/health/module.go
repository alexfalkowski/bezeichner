package health

import "github.com/alexfalkowski/go-signal/v2"

// Module registers health checks and HTTP and gRPC observers.
var Module = signal.Module(
	signal.Register(Register),
	signal.Register(httpHealthObserver),
	signal.Register(httpLivenessObserver),
	signal.Register(httpReadinessObserver),
	signal.Register(grpcObserver),
)
