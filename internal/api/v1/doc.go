// Package v1 wires the versioned API surface for Bezeichner.
//
// This package is responsible for composing the v1 API module using dependency
// injection. It connects the domain service (internal/api/ids) to the transport
// layers:
//
//   - gRPC transport: internal/api/v1/transport/grpc
//   - HTTP transport: internal/api/v1/transport/http
//
// The HTTP transport is implemented as an RPC gateway that routes requests by
// gRPC full method name, so both transports expose the same API contract defined
// in api/bezeichner/v1/service.proto.
//
// This package does not contain business logic; it only provides the module that
// registers constructors and transport registration hooks with the application's
// DI container.
package v1
