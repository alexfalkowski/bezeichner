package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/go-service/marshaller"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(NewConfigurator),
	marshaller.Module,
	config.ConfigModule,
	fx.Provide(healthConfig),
	fx.Provide(generatorConfig),
	fx.Provide(mapperConfig),
	fx.Provide(v1ClientConfig),
)
