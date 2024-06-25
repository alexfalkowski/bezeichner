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

	mapHandler struct {
		service *service.Service
	}
)

func (h *mapHandler) Handle(ctx context.Context, req *MapIdentifiersRequest) (*MapIdentifiersResponse, error) {
	resp := &MapIdentifiersResponse{}

	ids, err := h.service.MapIdentifiers(req.IDs)
	if err != nil {
		return resp, err
	}

	resp.IDs = ids
	resp.Meta = meta.CamelStrings(ctx, "")

	return resp, nil
}

func (h *mapHandler) Error(ctx context.Context, err error) *MapIdentifiersResponse {
	return &MapIdentifiersResponse{Meta: meta.CamelStrings(ctx, ""), Error: &Error{Message: err.Error()}}
}

func (h *mapHandler) Status(err error) int {
	if service.IsNotFoundError(err) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}