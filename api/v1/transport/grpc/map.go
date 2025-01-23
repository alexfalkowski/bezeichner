package grpc

import (
	"context"

	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/go-service/meta"
)

// MapIdentifiers for gRPC.
func (s *Server) MapIdentifiers(ctx context.Context, req *v1.MapIdentifiersRequest) (*v1.MapIdentifiersResponse, error) {
	resp := &v1.MapIdentifiersResponse{}
	ids, err := s.service.Map(req.GetIds())

	resp.Meta = meta.CamelStrings(ctx, "")
	resp.Ids = ids

	return resp, s.error(err)
}
