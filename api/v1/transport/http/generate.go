package http

import (
	"context"

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
)

// GenerateIdentifiers for HTTP.
func (h *Handler) GenerateIdentifiers(ctx context.Context, req *GenerateIdentifiersRequest) (*GenerateIdentifiersResponse, error) {
	resp := &GenerateIdentifiersResponse{}
	ids, err := h.id.Generate(ctx, req.Application, req.Count)

	resp.Meta = meta.CamelStrings(ctx, "")
	resp.IDs = ids

	return resp, h.error(err)
}
