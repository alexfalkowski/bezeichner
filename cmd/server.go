package cmd

import (
	"github.com/alexfalkowski/bezeichner/config"
	"github.com/alexfalkowski/bezeichner/server/health"
	v1 "github.com/alexfalkowski/bezeichner/server/v1"
	"github.com/alexfalkowski/go-service/database/sql"
	"github.com/alexfalkowski/go-service/logger"
	"github.com/alexfalkowski/go-service/metrics"
	"github.com/alexfalkowski/go-service/transport"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	fx.NopLogger, fx.Provide(NewVersion), config.Module, health.Module,
	logger.ZapModule, metrics.PrometheusModule, transport.Module,
	sql.PostgreSQLModule, sql.PostgreSQLOpentracingModule,
	v1.Module,
}
