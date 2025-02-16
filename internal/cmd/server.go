package cmd

import (
	v1 "github.com/alexfalkowski/bezeichner/internal/api/v1"
	"github.com/alexfalkowski/bezeichner/internal/config"
	"github.com/alexfalkowski/bezeichner/internal/generator"
	"github.com/alexfalkowski/bezeichner/internal/health"
	"github.com/alexfalkowski/go-service/cache"
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/database/sql"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/module"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/transport"
)

// RegisterServer for cmd.
func RegisterServer(command *cmd.Command) {
	flags := flags.NewFlagSet("server")
	flags.AddInput("env:BEZEICHNER_CONFIG_FILE")

	command.AddServer("server", "Start bezeichner server", flags,
		module.Module, feature.Module, debug.Module,
		transport.Module, telemetry.Module,
		cache.Module, sql.Module, health.Module,
		generator.Module, v1.Module, config.Module, cmd.Module,
	)
}
