package http

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/bezeichner/server/service"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/meta"
	nh "github.com/alexfalkowski/go-service/net/http"
	"go.uber.org/fx"
)

// RegisterParams for HTTP.
type RegisterParams struct {
	fx.In

	Marshaller *marshaller.Map
	Mux        nh.ServeMux
	Service    *service.Service
}

// Register for HTTP.
func Register(params RegisterParams) error {
	s := &Server{service: params.Service}

	gih := nh.NewHandler[GenerateIdentifiersRequest, GenerateIdentifiersResponse](params.Mux, params.Marshaller, &GenerateIdentifiersErrorer{})
	if err := gih.Handle("POST", "/v1/generate", s.GenerateIdentifiers); err != nil {
		return err
	}

	mih := nh.NewHandler[MapIdentifiersRequest, MapIdentifiersResponse](params.Mux, params.Marshaller, &MapIdentifiersErrorer{})
	if err := mih.Handle("POST", "/v1/map", s.MapIdentifiers); err != nil {
		return err
	}

	return nil
}

// Server for HTTP.
type Server struct {
	service *service.Service
}

// GenerateIdentifiers for HTTP.
func (s *Server) GenerateIdentifiers(ctx context.Context, req *GenerateIdentifiersRequest) (*GenerateIdentifiersResponse, error) {
	resp := &GenerateIdentifiersResponse{}

	ids, err := s.service.GenerateIdentifiers(ctx, req.Application, req.Count)
	if err != nil {
		return resp, err
	}

	resp.IDs = ids
	resp.Meta = meta.CamelStrings(ctx, "")

	return resp, nil
}

// MapIdentifiers for HTTP.
func (s *Server) MapIdentifiers(ctx context.Context, req *MapIdentifiersRequest) (*MapIdentifiersResponse, error) {
	resp := &MapIdentifiersResponse{}

	ids, err := s.service.MapIdentifiers(req.IDs)
	if err != nil {
		return resp, err
	}

	resp.IDs = ids
	resp.Meta = meta.CamelStrings(ctx, "")

	return resp, nil
}

// GenerateIdentifiersErrorer for HTTP.
type GenerateIdentifiersErrorer struct{}

func (*GenerateIdentifiersErrorer) Error(ctx context.Context, err error) *GenerateIdentifiersResponse {
	return &GenerateIdentifiersResponse{Meta: meta.CamelStrings(ctx, ""), Error: &Error{Message: err.Error()}}
}

func (*GenerateIdentifiersErrorer) Status(err error) int {
	if service.IsNotFoundError(err) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}

// GenerateIdentifiersErrorer for HTTP.
type MapIdentifiersErrorer struct{}

func (*MapIdentifiersErrorer) Error(ctx context.Context, err error) *MapIdentifiersResponse {
	return &MapIdentifiersResponse{Meta: meta.CamelStrings(ctx, ""), Error: &Error{Message: err.Error()}}
}

func (*MapIdentifiersErrorer) Status(err error) int {
	if service.IsNotFoundError(err) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}

// Error for HTTP.
type Error struct {
	Message string `json:"message,omitempty"`
}

// GenerateIdentifiersRequest for a specific application.
type GenerateIdentifiersRequest struct {
	Application string `json:"application,omitempty"`
	Count       uint64 `json:"count,omitempty"`
}

// GenerateIdentifiersResponse for a specific application.
type GenerateIdentifiersResponse struct {
	Meta  map[string]string `json:"meta,omitempty"`
	Error *Error            `json:"error,omitempty"`
	IDs   []string          `json:"ids,omitempty"`
}

// MapIdentifiersRequest for some identifiers.
type MapIdentifiersRequest struct {
	IDs []string `json:"ids,omitempty"`
}

// MapIdentifiersResponse for some identifiers.
type MapIdentifiersResponse struct {
	Meta  map[string]string `json:"meta,omitempty"`
	Error *Error            `json:"error,omitempty"`
	IDs   []string          `json:"ids,omitempty"`
}
