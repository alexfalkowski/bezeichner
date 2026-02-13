// Package http provides the HTTP transport for the Bezeichner v1 API.
//
// HTTP is implemented as an RPC gateway over the gRPC service contract.
// Rather than defining a separate REST surface, this gateway routes HTTP
// requests by gRPC full method name and invokes the corresponding gRPC
// handler method on the in-process server implementation.
//
// In other words, both transports (gRPC and HTTP) expose the same API
// defined in api/bezeichner/v1/service.proto; this package is only a thin
// routing adapter that binds HTTP endpoints to the gRPC handlers.
//
// # Routing
//
// Register installs routes using the protobuf-generated full method names,
// for example:
//
//   - v1.Service_GenerateIdentifiers_FullMethodName
//   - v1.Service_MapIdentifiers_FullMethodName
//
// Those routes are wired to methods on internal/api/v1/transport/grpc.Server.
//
// # Error semantics
//
// Since the HTTP gateway delegates directly to the gRPC handlers, error
// semantics and classification originate from the gRPC transport's mapping
// of domain errors to status codes.
package http
