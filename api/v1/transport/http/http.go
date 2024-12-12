package http

import (
	"net/http"

	"github.com/alexfalkowski/bezeichner/api/ids"
	"github.com/alexfalkowski/go-service/net/http/rpc"
	"github.com/alexfalkowski/go-service/net/http/status"
)

// Register for HTTP.
func Register(service *ids.Identifier) {
	gh := &generateHandler{service: service}
	rpc.Route("/v1/generate", gh.Generate)

	mh := &mapHandler{service: service}
	rpc.Route("/v1/map", mh.Map)
}

func handleError(err error) error {
	if ids.IsNotFound(err) {
		return status.Error(http.StatusNotFound, err.Error())
	}

	return err
}
