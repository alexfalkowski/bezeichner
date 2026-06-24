// Package grpc provides the gRPC transport registration for the Bezeichner v1 API.
//
// This package wraps the shared protobuf-facing identifier from
// internal/api/v1/ids and registers the wrapper with a gRPC service registrar.
//
// # Error mapping
//
// The transport maps domain errors to gRPC codes as follows:
//   - ids.ErrInvalidArgument -> codes.InvalidArgument
//   - ids.ErrNotFound        -> codes.NotFound
//   - all other non-nil domain errors currently fall back to codes.NotFound
//
// Callers should rely on these status codes (rather than parsing message text) for
// programmatic behavior.
//
// # Responses and metadata
//
// RPC handlers attach response metadata derived from the incoming context via
// the shared meta helpers in internal/api/v1/ids.
//
// # Registration
//
// Use Register to register the gRPC server with a gRPC service registrar.
package grpc
