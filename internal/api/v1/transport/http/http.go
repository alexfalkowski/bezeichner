package http

import (
	"github.com/alexfalkowski/bezeichner/internal/api/v1/transport/grpc"
	"github.com/alexfalkowski/go-service/net/http/rpc"
)

// Register for HTTP.
func Register(server *grpc.Server) {
	rpc.Route("/v1/generate", server.GenerateIdentifiers)
	rpc.Route("/v1/map", server.MapIdentifiers)
}
