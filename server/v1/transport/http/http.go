package http

import (
	"net/http"

	"github.com/alexfalkowski/bezeichner/server/service"
	nh "github.com/alexfalkowski/go-service/net/http"
)

// Register for HTTP.
func Register(service *service.Service) {
	nh.Handle("/v1/generate", &generateHandler{service: service})
	nh.Handle("/v1/map", &mapHandler{service: service})
}

func handleError(err error) error {
	if service.IsNotFound(err) {
		return nh.Error(http.StatusNotFound, err.Error())
	}

	return err
}
