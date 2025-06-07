package grpc

import (
	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/bezeichner/internal/api/ids"
	"github.com/alexfalkowski/go-service/v2/net/grpc"
	"github.com/alexfalkowski/go-service/v2/net/grpc/codes"
	"github.com/alexfalkowski/go-service/v2/net/grpc/status"
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

func (s *Server) error(err error) error {
	if err == nil {
		return nil
	}

	if ids.IsNotFound(err) {
		return status.Error(codes.NotFound, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
