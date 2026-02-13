// Package grpc provides the gRPC transport implementation for the Bezeichner v1 API.
//
// This package adapts the protobuf-defined service in api/bezeichner/v1 to the
// domain service in internal/api/ids. It is responsible for:
//   - Implementing the generated v1.ServiceServer interface.
//   - Translating transport requests into domain calls.
//   - Mapping domain error categories to gRPC status codes.
//
// # Error mapping
//
// The transport maps domain errors to gRPC codes as follows:
//   - ids.ErrInvalidArgument -> codes.InvalidArgument
//   - ids.ErrNotFound        -> codes.NotFound
//   - any other error        -> codes.Internal
//
// Callers should rely on these status codes (rather than parsing message text) for
// programmatic behavior.
//
// # Responses and metadata
//
// RPC handlers attach response metadata derived from the incoming context via the
// shared meta helpers.
//
// # Registration
//
// Use Register to register a Server implementation with a gRPC service registrar.
// Construct servers using NewServer, which injects the domain Identifier service.
package grpc
