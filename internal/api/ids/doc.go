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
//   - Map: map a list of existing identifiers to their configured replacements.
//
// Both operations enforce request-size limits as a simple DoS-protection mechanism.
// When limits are exceeded, Generate/Map return ErrInvalidArgument.
//
// # Configuration
//
// Generate depends on generator configuration (see internal/generator):
//
// An "application" is selected by name from generator configuration. The application
// specifies a generator kind (for example: "uuid", "ulid", "pg", ...). The kind is
// resolved through a Generators registry, and the selected Generator is asked to
// generate each identifier.
//
// Map depends on mapper configuration (see internal/mapper):
//
// The mapper configuration provides a lookup table from input identifier to mapped
// identifier. If any input identifier is missing, the operation returns ErrNotFound.
//
// # Errors
//
// The domain layer defines two primary error categories:
//
//   - ErrInvalidArgument: returned when request limits are exceeded.
//   - ErrNotFound: returned when a requested application, generator kind, or mapping
//     entry cannot be found.
//
// Helper predicates IsInvalidArgument and IsNotFound are provided to classify errors.
// Transports are expected to map these error categories to their protocol-specific
// equivalents (e.g., gRPC codes InvalidArgument and NotFound).
package ids
