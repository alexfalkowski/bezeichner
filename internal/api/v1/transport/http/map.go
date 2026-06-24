package http

import (
	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/go-service/v2/context"
)

// MapIdentifiers for HTTP.
func (s *Server) MapIdentifiers(ctx context.Context, req *v1.MapIdentifiersRequest) (*v1.MapIdentifiersResponse, error) {
	resp, err := s.id.MapIdentifiers(ctx, req)
	if err != nil {
		return nil, s.error(err)
	}

	return resp, nil
}
