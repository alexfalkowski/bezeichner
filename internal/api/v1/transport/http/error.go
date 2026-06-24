package http

import (
	"github.com/alexfalkowski/bezeichner/internal/api/ids"
	"github.com/alexfalkowski/go-service/v2/net/http"
	"github.com/alexfalkowski/go-service/v2/net/http/status"
)

func (s *Server) error(err error) error {
	if ids.IsInvalidArgument(err) {
		return status.SafeError(http.StatusBadRequest, err)
	}

	return status.SafeError(http.StatusNotFound, err)
}
