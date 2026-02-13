// Package mapper defines configuration for mapping identifiers.
//
// Mapping is used by the domain layer (internal/api/ids) to translate a set of
// input identifiers to their configured replacements.
//
// # Configuration
//
// The mapping table is represented by Config.Identifiers, which is a map from
// input identifier to mapped identifier:
//
//	input -> output
//
// For example, a configuration might map legacy IDs to new canonical IDs.
//
// # Semantics
//
// The domain operation that performs mapping preserves order: it returns outputs
// in the same order as the input slice.
//
// Mapping is strict: if any input identifier does not exist in the table, the
// operation fails with a "not found" error from the domain layer. This prevents
// silently returning partial results.
//
// # Size limits
//
// Request-size limits are enforced by the domain layer (not this package) to
// provide basic DoS protection.
package mapper
