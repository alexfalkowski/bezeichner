package http

import (
	"github.com/alexfalkowski/bezeichner/server/service"
	"github.com/alexfalkowski/go-service/net/http"
)

// Error for HTTP.
type Error struct {
	Message string `json:"message,omitempty"`
}

// Register for HTTP.
func Register(service *service.Service) {
	http.Handle("POST /v1/generate", &generateHandler{service: service})
	http.Handle("POST /v1/map", &mapHandler{service: service})
}
