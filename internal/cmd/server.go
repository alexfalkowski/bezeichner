package cmd

import (
	v1 "github.com/alexfalkowski/bezeichner/internal/api/v1"
	"github.com/alexfalkowski/bezeichner/internal/config"
	"github.com/alexfalkowski/bezeichner/internal/generator"
	"github.com/alexfalkowski/bezeichner/internal/health"
	"github.com/alexfalkowski/go-service/v2/cache"
	"github.com/alexfalkowski/go-service/v2/cli"
	"github.com/alexfalkowski/go-service/v2/database/sql"
	"github.com/alexfalkowski/go-service/v2/debug"
	"github.com/alexfalkowski/go-service/v2/feature"
	"github.com/alexfalkowski/go-service/v2/module"
	"github.com/alexfalkowski/go-service/v2/telemetry"
	"github.com/alexfalkowski/go-service/v2/transport"
)

// RegisterServer for cmd.
func RegisterServer(command cli.Commander) {
	cmd := command.AddServer("server", "Start bezeichner server",
		module.Module, feature.Module, debug.Module,
		transport.Module, telemetry.Module,
		cache.Module, sql.Module, health.Module,
		generator.Module, v1.Module, config.Module, cli.Module,
	)
	cmd.AddInput("")
}
