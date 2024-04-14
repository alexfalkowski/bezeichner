package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	v1c "github.com/alexfalkowski/bezeichner/client/v1/config"
	g "github.com/alexfalkowski/bezeichner/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/http"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	GRPCServer   *grpc.Server
	HTTPServer   *http.Server
	ClientConfig *v1c.Config
	Logger       *zap.Logger
	Tracer       trace.Tracer
	Meter        metric.Meter
	Server       v1.ServiceServer
}

// Register server.
func Register(params RegisterParams) error {
	v1.RegisterServiceServer(params.GRPCServer.Server, params.Server)

	opts := g.ClientOpts{
		Lifecycle:    params.Lifecycle,
		ClientConfig: params.ClientConfig.Config,
		Logger:       params.Logger,
		Tracer:       params.Tracer,
		Meter:        params.Meter,
	}

	conn, err := g.NewClient(opts)
	if err != nil {
		return err
	}

	return v1.RegisterServiceHandler(context.Background(), params.HTTPServer.ServeMux(), conn)
}
