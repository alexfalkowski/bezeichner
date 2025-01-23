package http

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
)

type (
	// MapIdentifiersRequest for some identifiers.
	MapIdentifiersRequest struct {
		IDs []string `json:"ids,omitempty"`
	}

	// MapIdentifiersResponse for some identifiers.
	MapIdentifiersResponse struct {
		Meta map[string]string `json:"meta,omitempty"`
		IDs  []string          `json:"ids,omitempty"`
	}
)

// MapIdentifiers for HTTP.
func (h *Handler) MapIdentifiers(ctx context.Context, req *MapIdentifiersRequest) (*MapIdentifiersResponse, error) {
	resp := &MapIdentifiersResponse{}
	ids, err := h.service.Map(req.IDs)

	resp.Meta = meta.CamelStrings(ctx, "")
	resp.IDs = ids

	return resp, h.error(err)
}
