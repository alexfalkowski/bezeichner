package http

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/bezeichner/server/service"
	"github.com/alexfalkowski/go-service/meta"
)

type (
	// MapIdentifiersRequest for some identifiers.
	MapIdentifiersRequest struct {
		IDs []string `json:"ids,omitempty"`
	}

	// MapIdentifiersResponse for some identifiers.
	MapIdentifiersResponse struct {
		Meta  map[string]string `json:"meta,omitempty"`
		Error *Error            `json:"error,omitempty"`
		IDs   []string          `json:"ids,omitempty"`
	}

	mapErrorer struct{}
)

// MapIdentifiers for HTTP.
func (s *Server) MapIdentifiers(ctx context.Context, req *MapIdentifiersRequest) (*MapIdentifiersResponse, error) {
	resp := &MapIdentifiersResponse{}

	ids, err := s.service.MapIdentifiers(req.IDs)
	if err != nil {
		return resp, err
	}

	resp.IDs = ids
	resp.Meta = meta.CamelStrings(ctx, "")

	return resp, nil
}

func (*mapErrorer) Error(ctx context.Context, err error) *MapIdentifiersResponse {
	return &MapIdentifiersResponse{Meta: meta.CamelStrings(ctx, ""), Error: &Error{Message: err.Error()}}
}

func (*mapErrorer) Status(err error) int {
	if service.IsNotFoundError(err) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
