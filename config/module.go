package config

import (
	"github.com/alexfalkowski/go-service/config"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(NewConfigurator),
	config.ConfigModule,
	fx.Provide(healthConfig),
	fx.Provide(generatorConfig),
	fx.Provide(mapperConfig),
	fx.Provide(v1ClientConfig),
)
