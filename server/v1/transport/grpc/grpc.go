package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/bezeichner/server/service"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Register server.
func Register(gs *grpc.Server, server v1.ServiceServer) {
	v1.RegisterServiceServer(gs.Server(), server)
}

// NewServer for gRPC.
func NewServer(service *service.Service) v1.ServiceServer {
	return &Server{service: service}
}

// Server for gRPC.
type Server struct {
	v1.UnimplementedServiceServer
	service *service.Service
}

// GetIdentifiers for gRPC.
func (s *Server) GenerateIdentifiers(ctx context.Context, req *v1.GenerateIdentifiersRequest) (*v1.GenerateIdentifiersResponse, error) {
	resp := &v1.GenerateIdentifiersResponse{}

	ids, err := s.service.GenerateIdentifiers(ctx, req.GetApplication(), req.GetCount())
	if err != nil {
		return resp, s.error(err)
	}

	resp.Ids = ids
	resp.Meta = meta.CamelStrings(ctx, "")

	return resp, nil
}

// MapIdentifiers for gRPC.
func (s *Server) MapIdentifiers(ctx context.Context, req *v1.MapIdentifiersRequest) (*v1.MapIdentifiersResponse, error) {
	resp := &v1.MapIdentifiersResponse{}

	ids, err := s.service.MapIdentifiers(req.GetIds())
	if err != nil {
		return resp, s.error(err)
	}

	resp.Ids = ids
	resp.Meta = meta.CamelStrings(ctx, "")

	return resp, nil
}

func (s *Server) error(err error) error {
	if service.IsNotFoundError(err) {
		return status.Error(codes.NotFound, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
