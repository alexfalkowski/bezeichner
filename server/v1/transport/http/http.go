package http

import (
	"net/http"

	"github.com/alexfalkowski/bezeichner/server/ids"
	"github.com/alexfalkowski/go-service/net/http/rpc"
)

// Register for HTTP.
func Register(service *ids.Identifier) {
	gh := &generateHandler{service: service}
	rpc.Unary("/v1/generate", gh.Generate)

	mh := &mapHandler{service: service}
	rpc.Unary("/v1/map", mh.Map)
}

func handleError(err error) error {
	if ids.IsNotFound(err) {
		return rpc.Error(http.StatusNotFound, err.Error())
	}

	return err
}
