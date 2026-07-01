package http

import (
	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/go-service/v2/context"
)

// ListApplications for HTTP.
func (s *Server) ListApplications(ctx context.Context, req *v1.ListApplicationsRequest) (*v1.ListApplicationsResponse, error) {
	resp, err := s.id.ListApplications(ctx, req)
	if err != nil {
		return nil, s.error(err)
	}

	return resp, nil
}
