// Package http provides the HTTP transport for the Bezeichner v1 API.
//
// HTTP is implemented as an RPC gateway over the gRPC service contract.
// Rather than defining a separate REST surface, this gateway routes HTTP
// requests by gRPC full method name and invokes the corresponding HTTP server
// wrapper method.
//
// In other words, both transports (gRPC and HTTP) expose the same API
// defined in api/bezeichner/v1/service.proto; this package is only a thin
// routing adapter that binds HTTP endpoints to the shared v1 identifier through
// an HTTP-specific wrapper.
//
// # Routing
//
// Register installs routes using the protobuf-generated full method names,
// for example:
//
//   - v1.Service_GenerateIdentifiers_FullMethodName
//   - v1.Service_ListApplications_FullMethodName
//   - v1.Service_MapIdentifiers_FullMethodName
//
// Those routes are wired to methods on Server, which delegates to
// internal/api/v1/ids.Identifier.
//
// # Error semantics
//
// Since the HTTP gateway delegates to the shared v1 identifier through its HTTP
// wrapper, this transport maps domain errors to HTTP statuses; for this API,
// InvalidArgument becomes HTTP 400 and NotFound becomes HTTP 404. Error bodies
// are written as safe plain-text messages using the go-service "text/error"
// media type.
package http
