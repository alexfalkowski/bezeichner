package grpc

import (
	v1 "github.com/alexfalkowski/bezeichner/v2/api/bezeichner/v1"
	"github.com/alexfalkowski/go-service/v3/context"
)

// GenerateIdentifiers for gRPC.
func (s *Server) GenerateIdentifiers(ctx context.Context, req *v1.GenerateIdentifiersRequest) (*v1.GenerateIdentifiersResponse, error) {
	resp, err := s.id.GenerateIdentifiers(ctx, req)
	if err != nil {
		return nil, s.error(err)
	}

	return resp, nil
}
