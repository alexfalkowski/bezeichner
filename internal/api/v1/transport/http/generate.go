package http

import (
	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/go-service/v2/context"
)

// GenerateIdentifiers for HTTP.
func (s *Server) GenerateIdentifiers(ctx context.Context, req *v1.GenerateIdentifiersRequest) (*v1.GenerateIdentifiersResponse, error) {
	resp, err := s.id.GenerateIdentifiers(ctx, req)
	if err != nil {
		return nil, s.error(err)
	}

	return resp, nil
}
