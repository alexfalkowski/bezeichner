package cmd

import (
	"github.com/alexfalkowski/bezeichner/config"
	"github.com/alexfalkowski/bezeichner/generator"
	"github.com/alexfalkowski/bezeichner/server/health"
	v1 "github.com/alexfalkowski/bezeichner/server/v1"
	"github.com/alexfalkowski/go-service/cache"
	"github.com/alexfalkowski/go-service/compress"
	"github.com/alexfalkowski/go-service/database/sql"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/encoding"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/sync"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/transport"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	sync.Module, compress.Module, encoding.Module,
	runtime.Module, debug.Module, feature.Module,
	transport.Module, health.Module,
	telemetry.Module, cache.Module, sql.Module,
	generator.Module, v1.Module,
	config.Module, Module,
}
