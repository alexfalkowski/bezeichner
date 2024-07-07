package http

import (
	"net/http"

	"github.com/alexfalkowski/bezeichner/server/ids"
	"github.com/alexfalkowski/go-service/net/http/rpc"
)

// Register for HTTP.
func Register(service *ids.Identifier) {
	rpc.Unary("/v1/generate", &generateHandler{service: service})
	rpc.Unary("/v1/map", &mapHandler{service: service})
}

func handleError(err error) error {
	if ids.IsNotFound(err) {
		return rpc.Error(http.StatusNotFound, err.Error())
	}

	return err
}
