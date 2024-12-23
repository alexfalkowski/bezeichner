package http

import (
	"context"

	"github.com/alexfalkowski/bezeichner/api/ids"
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

	mapHandler struct {
		service *ids.Identifier
	}
)

func (h *mapHandler) Map(ctx context.Context, req *MapIdentifiersRequest) (*MapIdentifiersResponse, error) {
	resp := &MapIdentifiersResponse{}

	ids, err := h.service.Map(req.IDs)
	if err != nil {
		resp.Meta = meta.CamelStrings(ctx, "")

		return resp, handleError(err)
	}

	resp.IDs = ids
	resp.Meta = meta.CamelStrings(ctx, "")

	return resp, nil
}
