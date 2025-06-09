package health

import (
	"github.com/alexfalkowski/go-health/checker"
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/v2/health"
	hc "github.com/alexfalkowski/go-service/v2/health/checker"
	"github.com/alexfalkowski/go-service/v2/time"
	"github.com/linxGnu/mssqlx"
)

func registrations(db *mssqlx.DBs, cfg *Config) health.Registrations {
	t := time.MustParseDuration(cfg.Timeout)
	d := time.MustParseDuration(cfg.Duration)
	registrations := health.Registrations{
		server.NewRegistration("noop", d, checker.NewNoopChecker()),
		server.NewOnlineRegistration(d, d),
	}

	if db != nil {
		registrations = append(registrations, server.NewRegistration("pg", d, hc.NewDBChecker(db, t)))
	}

	return registrations
}
