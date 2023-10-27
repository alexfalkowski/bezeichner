package grpc

import (
	"context"
	"fmt"

	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc/telemetry/tracer"
	"github.com/alexfalkowski/go-service/transport/http"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/fx"
	"go.uber.org/zap"
	g "google.golang.org/grpc"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	Lifecycle  fx.Lifecycle
	GRPCServer *grpc.Server
	HTTPServer *http.Server
	GRPCConfig *grpc.Config
	Logger     *zap.Logger
	Tracer     tracer.Tracer
	Meter      metric.Meter
	Server     v1.ServiceServer
}

// Register server.
func Register(params RegisterParams) {
	v1.RegisterServiceServer(params.GRPCServer.Server, params.Server)

	var conn *g.ClientConn

	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			c, err := grpc.NewClient(ctx, fmt.Sprintf("127.0.0.1:%s", params.GRPCConfig.Port), params.GRPCConfig,
				grpc.WithClientLogger(params.Logger), grpc.WithClientTracer(params.Tracer), grpc.WithClientMetrics(params.Meter),
			)
			if err != nil {
				return err
			}

			conn = c

			return v1.RegisterServiceHandler(ctx, params.HTTPServer.Mux, c)
		},
		OnStop: func(ctx context.Context) error {
			if conn == nil {
				return nil
			}

			return conn.Close()
		},
	})
}
