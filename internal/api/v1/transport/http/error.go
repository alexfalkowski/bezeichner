package http

import (
	"github.com/alexfalkowski/bezeichner/v2/internal/api/ids"
	"github.com/alexfalkowski/go-service/v3/net/http"
	"github.com/alexfalkowski/go-service/v3/net/http/status"
)

func (s *Server) error(err error) error {
	if ids.IsInvalidArgument(err) {
		return status.SafeError(http.StatusBadRequest, err)
	}

	return status.SafeError(http.StatusNotFound, err)
}
