package grpc

import (
	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/bezeichner/api/ids"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Register server.
func Register(gs *grpc.Server, server v1.ServiceServer) {
	v1.RegisterServiceServer(gs.Server(), server)
}

// NewServer for gRPC.
func NewServer(id *ids.Identifier) v1.ServiceServer {
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
