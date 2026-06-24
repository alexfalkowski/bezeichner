// Package ids provides the protobuf-facing identifier service for the
// Bezeichner v1 API.
//
// This package adapts the transport-agnostic domain service in internal/api/ids
// to the generated api/bezeichner/v1 request and response types. The gRPC and
// HTTP transports wrap this identifier, keeping request handling shared while
// leaving transport-specific error classification and extension points in each
// transport package.
package ids
