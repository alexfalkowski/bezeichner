package http

import (
	"github.com/alexfalkowski/bezeichner/server/service"
	"github.com/alexfalkowski/go-service/net/http"
)

type (
	// Server for HTTP.
	Server struct {
		service *service.Service
	}

	// Error for HTTP.
	Error struct {
		Message string `json:"message,omitempty"`
	}
)

// Register for HTTP.
func Register(service *service.Service) {
	s := &Server{service: service}

	http.Handler("POST /v1/generate", &generateErrorer{}, s.GenerateIdentifiers)
	http.Handler("POST /v1/map", &mapErrorer{}, s.MapIdentifiers)
}
