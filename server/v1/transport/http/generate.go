package http

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/bezeichner/server/service"
	"github.com/alexfalkowski/go-service/meta"
)

type (
	// GenerateIdentifiersRequest for a specific application.
	GenerateIdentifiersRequest struct {
		Application string `json:"application,omitempty"`
		Count       uint64 `json:"count,omitempty"`
	}

	// GenerateIdentifiersResponse for a specific application.
	GenerateIdentifiersResponse struct {
		Meta  map[string]string `json:"meta,omitempty"`
		Error *Error            `json:"error,omitempty"`
		IDs   []string          `json:"ids,omitempty"`
	}

	generateHandler struct {
		service *service.Service
	}
)

func (h *generateHandler) Handle(ctx context.Context, req *GenerateIdentifiersRequest) (*GenerateIdentifiersResponse, error) {
	resp := &GenerateIdentifiersResponse{}

	ids, err := h.service.GenerateIdentifiers(ctx, req.Application, req.Count)
	if err != nil {
		return resp, err
	}

	resp.IDs = ids
	resp.Meta = meta.CamelStrings(ctx, "")

	return resp, nil
}

func (h *generateHandler) Error(ctx context.Context, err error) *GenerateIdentifiersResponse {
	return &GenerateIdentifiersResponse{Meta: meta.CamelStrings(ctx, ""), Error: &Error{Message: err.Error()}}
}

func (h *generateHandler) Status(err error) int {
	if service.IsNotFoundError(err) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
