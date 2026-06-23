package grpc

import (
	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/meta"
	"github.com/alexfalkowski/go-service/v2/strings"
)

// MapIdentifiers for gRPC.
func (s *Server) MapIdentifiers(ctx context.Context, req *v1.MapIdentifiersRequest) (*v1.MapIdentifiersResponse, error) {
	resp := &v1.MapIdentifiersResponse{}
	mapped, unmapped, err := s.id.Map(req.GetApplication(), req.GetIds())

	resp.Meta = meta.CamelStrings(ctx, strings.Empty)
	resp.Mapped = mapped
	resp.Unmapped = unmapped

	return resp, s.error(err)
}
