package http

import (
	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/bezeichner/internal/api/v1/transport/grpc"
	"github.com/alexfalkowski/go-service/v2/net/http/rpc"
)

// Register for HTTP.
func Register(server *grpc.Server) {
	rpc.Route(v1.Service_GenerateIdentifiers_FullMethodName, server.GenerateIdentifiers)
	rpc.Route(v1.Service_MapIdentifiers_FullMethodName, server.MapIdentifiers)
}
