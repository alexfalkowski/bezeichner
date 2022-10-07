package grpc

import (
	"context"
	"fmt"

	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/bezeichner/generator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ServerParams for gRPC.
type ServerParams struct {
	Config *generator.Config
}

// NewServer for gRPC.
func NewServer(params ServerParams) v1.ServiceServer {
	return &Server{config: params.Config}
}

// Server for gRPC.
type Server struct {
	config *generator.Config

	v1.UnimplementedServiceServer
}

// GetIdentifiers for gRPC.
func (s *Server) GetIdentifiers(ctx context.Context, req *v1.GetIdentifiersRequest) (*v1.GetIdentifiersResponse, error) {
	if req.Count == 0 {
		req.Count = 1
	}

	resp := &v1.GetIdentifiersResponse{}

	app := s.config.Application(req.Application)
	if app == nil {
		return resp, status.Error(codes.NotFound, fmt.Sprintf("%s: not found", req.Application))
	}

	g, err := generator.NewGenerator(app.Kind)
	if err != nil {
		return resp, status.Error(codes.NotFound, err.Error())
	}

	ids := make([]string, req.Count)
	for i := 0; i < len(ids); i++ {
		id, err := g.Generate(ctx)
		if err != nil {
			return resp, status.Error(codes.Internal, err.Error())
		}

		ids[i] = app.ID(id)
	}

	resp.Ids = ids

	return resp, nil
}
