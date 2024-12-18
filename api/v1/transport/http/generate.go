package http

import (
	"context"

	"github.com/alexfalkowski/bezeichner/api/ids"
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
		Meta map[string]string `json:"meta,omitempty"`
		IDs  []string          `json:"ids,omitempty"`
	}

	generateHandler struct {
		service *ids.Identifier
	}
)

func (h *generateHandler) Generate(ctx context.Context, req *GenerateIdentifiersRequest) (*GenerateIdentifiersResponse, error) {
	resp := &GenerateIdentifiersResponse{}

	ids, err := h.service.Generate(ctx, req.Application, req.Count)
	if err != nil {
		resp.Meta = meta.CamelStrings(ctx, "")

		return resp, handleError(err)
	}

	resp.IDs = ids
	resp.Meta = meta.CamelStrings(ctx, "")

	return resp, nil
}
