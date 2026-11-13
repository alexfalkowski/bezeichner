package grpc

import (
	v1 "github.com/alexfalkowski/bezeichner/v2/api/bezeichner/v1"
	"github.com/alexfalkowski/go-service/v3/context"
)

// ListApplications for gRPC.
func (s *Server) ListApplications(ctx context.Context, req *v1.ListApplicationsRequest) (*v1.ListApplicationsResponse, error) {
	resp, err := s.id.ListApplications(ctx, req)
	if err != nil {
		return nil, s.error(err)
	}

	return resp, nil
}
