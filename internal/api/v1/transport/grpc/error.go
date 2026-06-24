package grpc

import (
	"github.com/alexfalkowski/bezeichner/internal/api/ids"
	"github.com/alexfalkowski/go-service/v2/net/grpc/codes"
	"github.com/alexfalkowski/go-service/v2/net/grpc/status"
)

func (s *Server) error(err error) error {
	if ids.IsInvalidArgument(err) {
		return status.SafeError(codes.InvalidArgument, err)
	}

	return status.SafeError(codes.NotFound, err)
}
