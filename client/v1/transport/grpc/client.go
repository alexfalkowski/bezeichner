package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	v1c "github.com/alexfalkowski/bezeichner/client/v1/config"
	"github.com/alexfalkowski/bezeichner/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc/telemetry/tracer"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ServiceClientParams for gRPC.
type ServiceClientParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	ClientConfig *v1c.Config
	Logger       *zap.Logger
	Tracer       tracer.Tracer
	Meter        metric.Meter
}

// NewServiceClient for gRPC.
func NewServiceClient(params ServiceClientParams) (v1.ServiceClient, error) {
	opts := grpc.ClientOpts{
		Lifecycle:    params.Lifecycle,
		ClientConfig: params.ClientConfig.Config,
		Logger:       params.Logger,
		Tracer:       params.Tracer,
		Meter:        params.Meter,
	}

	conn, err := grpc.NewClient(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	return v1.NewServiceClient(conn), nil
}
