package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/client"
	"github.com/alexfalkowski/go-service/security"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
	g "google.golang.org/grpc"
)

// ClientOpts for gRPC.
type ClientOpts struct {
	Lifecycle    fx.Lifecycle
	ClientConfig *client.Config
	Logger       *zap.Logger
	Tracer       trace.Tracer
	Meter        metric.Meter
}

// NewClient for gRPC.
func NewClient(options ClientOpts) (*g.ClientConn, error) {
	opts := []grpc.ClientOption{
		grpc.WithClientLogger(options.Logger), grpc.WithClientTracer(options.Tracer),
		grpc.WithClientMetrics(options.Meter), grpc.WithClientRetry(options.ClientConfig.Retry),
		grpc.WithClientUserAgent(options.ClientConfig.UserAgent),
	}

	if security.IsEnabled(options.ClientConfig.Security) {
		sec, err := grpc.WithClientSecure(options.ClientConfig.Security)
		if err != nil {
			return nil, err
		}

		opts = append(opts, sec)
	}

	conn, err := grpc.NewClient(options.ClientConfig.Host, opts...)
	if err != nil {
		return nil, err
	}

	options.Lifecycle.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			return conn.Close()
		},
	})

	return conn, nil
}
