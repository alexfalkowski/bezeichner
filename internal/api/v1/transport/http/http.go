package http

import (
	v1 "github.com/alexfalkowski/bezeichner/api/bezeichner/v1"
	"github.com/alexfalkowski/bezeichner/internal/api/v1/ids"
	"github.com/alexfalkowski/go-service/v2/net/http/rpc"
)

// Register for HTTP.
func Register(server *Server) {
	server.Route(v1.Service_GenerateIdentifiers_FullMethodName, server.GenerateIdentifiers)
	server.Route(v1.Service_ListApplications_FullMethodName, server.ListApplications)
	server.Route(v1.Service_MapIdentifiers_FullMethodName, server.MapIdentifiers)
}

// NewServer for HTTP.
func NewServer(server *rpc.Server, id *ids.Identifier) *Server {
	return &Server{id: id, Server: server}
}

// Server for HTTP.
type Server struct {
	id *ids.Identifier
	*rpc.Server
}
