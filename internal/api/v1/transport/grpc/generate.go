package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/go-service/v2/meta"
)

// GetIdentifiers for gRPC.
func (s *Server) GenerateIdentifiers(ctx context.Context, req *v1.GenerateIdentifiersRequest) (*v1.GenerateIdentifiersResponse, error) {
	resp := &v1.GenerateIdentifiersResponse{}
	ids, err := s.id.Generate(ctx, req.GetApplication(), req.GetCount())

	resp.Meta = meta.CamelStrings(ctx, "")
	resp.Ids = ids

	return resp, s.error(err)
}
