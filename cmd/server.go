package cmd

import (
	"github.com/alexfalkowski/bezeichner/config"
	"github.com/alexfalkowski/bezeichner/generator"
	"github.com/alexfalkowski/bezeichner/server/health"
	v1 "github.com/alexfalkowski/bezeichner/server/v1"
	"github.com/alexfalkowski/bezeichner/transport"
	"github.com/alexfalkowski/go-service/cache"
	"github.com/alexfalkowski/go-service/database/sql"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	fx.NopLogger, runtime.Module, debug.Module, feature.Module,
	config.Module, health.Module, Module,
	telemetry.Module, metrics.Module, transport.Module,
	cache.RedisModule, cache.SnappyCompressorModule, cache.ProtoMarshallerModule,
	sql.PostgreSQLModule, generator.Module, v1.Module,
}
