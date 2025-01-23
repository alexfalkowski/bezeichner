package http

import (
	"net/http"

	"github.com/alexfalkowski/bezeichner/api/ids"
	"github.com/alexfalkowski/go-service/net/http/rpc"
	"github.com/alexfalkowski/go-service/net/http/status"
)

// Register for HTTP.
func Register(handler *Handler) {
	rpc.Route("/v1/generate", handler.GenerateIdentifiers)
	rpc.Route("/v1/map", handler.MapIdentifiers)
}

// NewHandler for HTTP.
func NewHandler(service *ids.Identifier) *Handler {
	return &Handler{service: service}
}

// Handler for HTTP.
type Handler struct {
	service *ids.Identifier
}

func (h *Handler) error(err error) error {
	if err == nil {
		return nil
	}

	if ids.IsNotFound(err) {
		return status.Error(http.StatusNotFound, err.Error())
	}

	return err
}
