package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/bezeichner/generator"
	"github.com/alexfalkowski/bezeichner/mapper"
	"github.com/alexfalkowski/go-service/meta"
	"go.uber.org/fx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ServerParams for gRPC.
type ServerParams struct {
	fx.In

	GeneratorConfig *generator.Config
	MapperConfig    *mapper.Config
	Generators      generator.Generators
}

// NewServer for gRPC.
func NewServer(params ServerParams) v1.ServiceServer {
	return &Server{generatorConfig: params.GeneratorConfig, mapperConfig: params.MapperConfig, generators: params.Generators}
}

// Server for gRPC.
type Server struct {
	generatorConfig *generator.Config
	mapperConfig    *mapper.Config
	generators      generator.Generators

	v1.UnimplementedServiceServer
}

// GetIdentifiers for gRPC.
func (s *Server) GenerateIdentifiers(ctx context.Context, req *v1.GenerateIdentifiersRequest) (*v1.GenerateIdentifiersResponse, error) {
	if req.GetCount() == 0 {
		req.Count = 1
	}

	resp := &v1.GenerateIdentifiersResponse{}

	app := s.generatorConfig.Application(req.GetApplication())
	if app == nil {
		return resp, status.Error(codes.NotFound, req.GetApplication()+": not found")
	}

	g, err := s.generators.Generator(app.Kind)
	if err != nil {
		return resp, status.Error(codes.NotFound, err.Error())
	}

	ids := make([]string, req.GetCount())
	for i := 0; i < len(ids); i++ {
		id, err := g.Generate(ctx, app)
		if err != nil {
			return resp, status.Error(codes.Internal, err.Error())
		}

		ids[i] = id
	}

	resp.Ids = ids
	resp.Meta = meta.Strings(ctx)

	return resp, nil
}

// MapIdentifiers for gRPC.
func (s *Server) MapIdentifiers(ctx context.Context, req *v1.MapIdentifiersRequest) (*v1.MapIdentifiersResponse, error) {
	ids := req.GetIds()
	resp := &v1.MapIdentifiersResponse{
		Ids:  make([]string, len(ids)),
		Meta: meta.Strings(ctx),
	}

	for i, id := range ids {
		mid, ok := s.mapperConfig.Identifiers[id]
		if !ok {
			return resp, status.Error(codes.NotFound, id+": not found")
		}

		resp.Ids[i] = mid
	}

	return resp, nil
}
