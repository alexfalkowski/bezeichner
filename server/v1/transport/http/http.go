package http

import (
	"github.com/alexfalkowski/bezeichner/server/service"
	"github.com/alexfalkowski/go-service/net/http"
)

// Register for HTTP.
func Register(service *service.Service) {
	http.Handle("/v1/generate", &generateHandler{service: service})
	http.Handle("/v1/map", &mapHandler{service: service})
}
