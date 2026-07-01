// Package mapper defines configuration for mapping identifiers.
//
// Mapping is used by the domain layer (internal/api/ids) to translate a set of
// input identifiers to their configured replacements for a named application.
//
// # Configuration
//
// Mapping configuration is represented by Config.Applications. Each application
// has a name and an Identifiers map from input identifier to mapped identifier:
//
//	application:
//	  input -> output
//
// For example, a configuration might map legacy IDs to new canonical IDs.
//
// # Semantics
//
// The domain operation that performs mapping returns one result for each input
// identifier. Known identifiers include a mapped value, while missing
// identifiers omit it.
//
// Mapper configuration is optional at service startup, but mapping still
// requires it. If mapper configuration is omitted, the requested application is
// not configured, the operation fails with a "not found" error from the domain
// layer.
//
// # Size limits
//
// Request-size limits are enforced by the domain layer (not this package) to
// provide basic DoS protection.
package mapper
