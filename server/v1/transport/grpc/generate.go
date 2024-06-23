package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/go-service/meta"
)

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
