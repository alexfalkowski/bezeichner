package health

import (
	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/go-health/v2/checker"
	"github.com/alexfalkowski/go-health/v2/server"
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/go-service/v2/env"
	"github.com/alexfalkowski/go-service/v2/health"
	hc "github.com/alexfalkowski/go-service/v2/health/checker"
	"github.com/alexfalkowski/go-service/v2/time"
	"github.com/linxGnu/mssqlx"
)

// RegisterParams for health.
type RegisterParams struct {
	di.In
	Name   env.Name
	Server *server.Server
	DB     *mssqlx.DBs
	Config *Config
}

// Register for health.
func Register(params RegisterParams) {
	t := time.MustParseDuration(params.Config.Timeout)
	d := time.MustParseDuration(params.Config.Duration)
	regs := health.Registrations{
		server.NewRegistration("noop", d, checker.NewNoopChecker()),
		server.NewOnlineRegistration(d, d),
	}

	if params.DB != nil {
		regs = append(regs, server.NewRegistration("pg", d, hc.NewDBChecker(params.DB, t)))
	}

	params.Server.Register(params.Name.String(), regs...)
	params.Server.Register(v1.Service_ServiceDesc.ServiceName, regs[0])
}

func httpHealthObserver(name env.Name, server *server.Server) error {
	return server.Observe(name.String(), "healthz", "pg", "online")
}

func httpLivenessObserver(name env.Name, server *server.Server) error {
	return server.Observe(name.String(), "livez", "noop")
}

func httpReadinessObserver(name env.Name, server *server.Server) error {
	return server.Observe(name.String(), "readyz", "noop")
}

func grpcObserver(server *server.Server) error {
	return server.Observe(v1.Service_ServiceDesc.ServiceName, "grpc", "noop")
}
