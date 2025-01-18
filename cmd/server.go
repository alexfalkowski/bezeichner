package cmd

import (
	v1 "github.com/alexfalkowski/bezeichner/api/v1"
	"github.com/alexfalkowski/bezeichner/config"
	"github.com/alexfalkowski/bezeichner/generator"
	"github.com/alexfalkowski/bezeichner/health"
	"github.com/alexfalkowski/go-service/cache"
	"github.com/alexfalkowski/go-service/database/sql"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/module"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/transport"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	module.Module, feature.Module, debug.Module,
	transport.Module, telemetry.Module,
	cache.Module, sql.Module, health.Module,
	generator.Module, v1.Module, config.Module, Module,
}
