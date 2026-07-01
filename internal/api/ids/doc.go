// Package ids provides the domain service for generating and mapping identifiers.
//
// This package contains the core business logic that backs the public Bezeichner API
// (exposed via gRPC and an HTTP RPC gateway).
//
// # Overview
//
// The main entry point is Identifier, which offers two operations:
//
//   - Generate: produce one or more identifiers for a named application.
//   - Map: classify a list of existing identifiers as mapped or unmapped.
//
// Both operations enforce request-size limits as a simple DoS-protection
// mechanism. When limits are exceeded, Generate/Map return ErrInvalidArgument.
//
// # Configuration
//
// Generate depends on generator configuration (see internal/generator):
//
// An "application" is selected by name from generator configuration. The application
// specifies a generator kind (for example: "uuid", "ulid", ...). The kind is
// resolved through a Generators registry, and the selected Generator is asked to
// generate each identifier.
//
// Map depends on mapper configuration (see internal/mapper):
//
// The optional mapper configuration provides a lookup table from input
// identifier to mapped identifier. If mapper configuration is omitted, or if the
// requested application is missing, the operation returns ErrNotFound. Missing
// input identifiers are returned in Map's unmapped list instead of failing the
// whole operation.
//
// # Errors
//
// The domain layer defines two primary error categories:
//
//   - ErrInvalidArgument: returned when request limits are exceeded.
//   - ErrNotFound: returned when a requested application, generator kind, or
//     mapper configuration cannot be found.
//
// Use IsInvalidArgument to classify ErrInvalidArgument without relying on error
// strings. Operation errors may wrap ErrNotFound with request-specific context,
// so classify not-found errors with errors.Is rather than direct equality or
// message parsing. Transports are expected to map these error categories to
// their protocol-specific equivalents (e.g., gRPC codes InvalidArgument and
// NotFound).
package ids
