package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/go-service/meta"
)

// MapIdentifiers for gRPC.
func (s *Server) MapIdentifiers(ctx context.Context, req *v1.MapIdentifiersRequest) (*v1.MapIdentifiersResponse, error) {
	resp := &v1.MapIdentifiersResponse{}

	ids, err := s.service.MapIdentifiers(req.GetIds())
	if err != nil {
		resp.Meta = meta.CamelStrings(ctx, "")

		return resp, s.error(err)
	}

	resp.Ids = ids
	resp.Meta = meta.CamelStrings(ctx, "")

	return resp, nil
}
