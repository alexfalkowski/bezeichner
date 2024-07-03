package http

import (
	"net/http"

	"github.com/alexfalkowski/bezeichner/server/ids"
	nh "github.com/alexfalkowski/go-service/net/http"
)

// Register for HTTP.
func Register(service *ids.Identifier) {
	nh.Handle("/v1/generate", &generateHandler{service: service})
	nh.Handle("/v1/map", &mapHandler{service: service})
}

func handleError(err error) error {
	if ids.IsNotFound(err) {
		return nh.Error(http.StatusNotFound, err.Error())
	}

	return err
}
