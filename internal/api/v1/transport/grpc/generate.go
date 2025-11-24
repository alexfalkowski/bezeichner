package grpc

import (
	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/meta"
	"github.com/alexfalkowski/go-service/v2/strings"
)

// GenerateIdentifiers for gRPC.
func (s *Server) GenerateIdentifiers(ctx context.Context, req *v1.GenerateIdentifiersRequest) (*v1.GenerateIdentifiersResponse, error) {
	resp := &v1.GenerateIdentifiersResponse{}
	ids, err := s.id.Generate(ctx, req.GetApplication(), req.GetCount())

	resp.Meta = meta.CamelStrings(ctx, strings.Empty)
	resp.Ids = ids

	return resp, s.error(err)
}
