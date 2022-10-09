package grpc

import (
	"context"
	"fmt"

	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/bezeichner/generator"
	"github.com/alexfalkowski/bezeichner/mapper"
	"github.com/alexfalkowski/go-service/cache/redis/client"
	"github.com/linxGnu/mssqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ServerParams for gRPC.
type ServerParams struct {
	GeneratorConfig *generator.Config
	MapperConfig    *mapper.Config
	DB              *mssqlx.DBs
	Client          client.Client
}

// NewServer for gRPC.
func NewServer(params ServerParams) v1.ServiceServer {
	return &Server{generatorConfig: params.GeneratorConfig, mapperConfig: params.MapperConfig, db: params.DB, client: params.Client}
}

// Server for gRPC.
type Server struct {
	generatorConfig *generator.Config
	mapperConfig    *mapper.Config
	db              *mssqlx.DBs
	client          client.Client

	v1.UnimplementedServiceServer
}

// GetIdentifiers for gRPC.
func (s *Server) GenerateIdentifiers(ctx context.Context, req *v1.GenerateIdentifiersRequest) (*v1.GenerateIdentifiersResponse, error) {
	if req.Count == 0 {
		req.Count = 1
	}

	resp := &v1.GenerateIdentifiersResponse{}

	app := s.generatorConfig.Application(req.Application)
	if app == nil {
		return resp, status.Error(codes.NotFound, fmt.Sprintf("%s: not found", req.Application))
	}

	g, err := generator.NewGenerator(app.Name, app.Kind, s.db, s.client)
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

// MapIdentifiers for gRPC.
func (s *Server) MapIdentifiers(ctx context.Context, req *v1.MapIdentifiersRequest) (*v1.MapIdentifiersResponse, error) {
	resp := &v1.MapIdentifiersResponse{
		Ids: make([]string, len(req.Ids)),
	}

	for i, id := range req.Ids {
		mid, ok := s.mapperConfig.Identifiers[id]
		if !ok {
			return resp, status.Error(codes.NotFound, fmt.Sprintf("%s: not found", id))
		}

		resp.Ids[i] = mid
	}

	return resp, nil
}
