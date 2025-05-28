package cmd

import (
	v1 "github.com/alexfalkowski/bezeichner/internal/api/v1"
	"github.com/alexfalkowski/bezeichner/internal/config"
	"github.com/alexfalkowski/bezeichner/internal/generator"
	"github.com/alexfalkowski/bezeichner/internal/health"
	"github.com/alexfalkowski/go-service/v2/module"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	module.Server,
	config.Module,
	health.Module,
	generator.Module,
	v1.Module,
)
