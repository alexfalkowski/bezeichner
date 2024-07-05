package http

import (
	"github.com/alexfalkowski/bezeichner/server/ids"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/net/http/rpc"
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

func (h *generateHandler) Handle(ctx rpc.Context, req *GenerateIdentifiersRequest) (*GenerateIdentifiersResponse, error) {
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
