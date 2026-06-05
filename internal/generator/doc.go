// Package generator provides identifier generator implementations and a registry
// for selecting them by kind.
//
// The Bezeichner domain layer (internal/api/ids) uses this package to generate
// identifiers for a configured "application". Each application selects a
// generator kind (for example "uuid" or "ulid"). At runtime, the kind is
// resolved through a Generators registry and the selected Generator is invoked.
//
// # Configuration model
//
// The generator configuration is represented by Config and Application:
//
//   - Config contains a list of Applications.
//   - Application is addressed by Name and selects a generator Kind.
//
// The domain service resolves an Application by name using (*Config).Application.
//
// # Registry
//
// NewGenerators constructs the default registry (Generators), which maps a kind
// string to an implementation of the Generator interface. The registry is used
// to resolve generators by kind via (Generators).Generator.
//
// # Generator interface
//
// Generator is the common interface implemented by all generators:
//
//	Generate(ctx, app) string
//
// The app parameter provides access to application configuration. Current
// built-in generators ignore it.
//
// # Built-in kinds
//
// The default registry includes (at the time of writing):
//
//   - "uuid":      UUIDv7 string
//   - "ksuid":     KSUID string
//   - "ulid":      ULID string
//   - "xid":       XID string
//   - "snowflake": Sonyflake-based numeric ID (decimal string)
//   - "nanoid":    NanoID string
//   - "typeid":    TypeID string
//
// # Errors
//
// When resolving a kind, (Generators).Generator returns ErrNotFound if the kind
// is missing from the registry.
package generator
