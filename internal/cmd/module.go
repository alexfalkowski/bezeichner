package cmd

import (
	v1 "github.com/alexfalkowski/bezeichner/internal/api/v1"
	"github.com/alexfalkowski/bezeichner/internal/config"
	"github.com/alexfalkowski/bezeichner/internal/generator"
	"github.com/alexfalkowski/bezeichner/internal/health"
	"github.com/alexfalkowski/go-service/v3/module"
	"github.com/alexfalkowski/go-signal/v2"
)

// Module composes the server with service configuration, health checks, and the v1 API.
var Module = signal.Module(
	module.Server,
	config.Module,
	health.Module,
	generator.Module,
	v1.Module,
)
