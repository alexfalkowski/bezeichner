package http

import (
	"github.com/alexfalkowski/bezeichner/server/service"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/net/http"
	"go.uber.org/fx"
)

type (
	// RegisterParams for HTTP.
	RegisterParams struct {
		fx.In

		Marshaller *marshaller.Map
		Mux        http.ServeMux
		Service    *service.Service
	}

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
func Register(params RegisterParams) error {
	s := &Server{service: params.Service}

	gih := http.NewHandler[GenerateIdentifiersRequest](params.Mux, params.Marshaller, &generateErrorer{})
	if err := gih.Handle("POST", "/v1/generate", s.GenerateIdentifiers); err != nil {
		return err
	}

	mih := http.NewHandler[MapIdentifiersRequest](params.Mux, params.Marshaller, &mapErrorer{})
	if err := mih.Handle("POST", "/v1/map", s.MapIdentifiers); err != nil {
		return err
	}

	return nil
}
