package health

import (
	"github.com/alexfalkowski/go-health/checker"
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/health"
	hc "github.com/alexfalkowski/go-service/health/checker"
	"github.com/alexfalkowski/go-service/redis"
	"github.com/alexfalkowski/go-service/time"
	"github.com/linxGnu/mssqlx"
	"go.uber.org/fx"
)

// Params for health.
type Params struct {
	fx.In

	Health *Config
	Redis  redis.Client
	DB     *mssqlx.DBs
}

// NewRegistrations for health.
func NewRegistrations(params Params) health.Registrations {
	t := time.MustParseDuration(params.Health.Timeout)
	d := time.MustParseDuration(params.Health.Duration)
	registrations := health.Registrations{
		server.NewRegistration("noop", d, checker.NewNoopChecker()),
		server.NewRegistration("redis", d, hc.NewRedisChecker(params.Redis, t)),
		server.NewRegistration("pg", d, hc.NewDBChecker(params.DB, t)),
	}

	return registrations
}
