package grpc

import (
	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/bezeichner/internal/api/v1/ids"
	"github.com/alexfalkowski/go-service/v2/net/grpc"
)

// Register server.
func Register(registrar grpc.ServiceRegistrar, server *Server) {
	v1.RegisterServiceServer(registrar, server)
}

// NewServer for gRPC.
func NewServer(id *ids.Identifier) *Server {
	return &Server{id: id}
}

// Server for gRPC.
type Server struct {
	v1.UnimplementedServiceServer
	id *ids.Identifier
}
