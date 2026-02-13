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
//	Generate(ctx, app) (string, error)
//
// The app parameter provides access to application configuration. Some
// generators ignore it; others use it (for example the Postgres generator uses
// app.Name to select a sequence).
//
// # Built-in kinds
//
// The default registry includes (at the time of writing):
//
//   - "uuid":      random UUID string
//   - "ksuid":     KSUID string
//   - "ulid":      ULID string
//   - "xid":       XID string
//   - "snowflake": Sonyflake-based numeric ID (decimal string)
//   - "nanoid":    NanoID string
//   - "typeid":    TypeID string
//   - "pg":        Postgres sequence value (decimal string)
//
// # Postgres ("pg") generator
//
// The "pg" generator reads the next value from a Postgres sequence using:
//
//	SELECT nextval($1::regclass)
//
// The sequence name is taken from Application.Name. Sequences are not created by
// this service and must be provisioned externally.
//
// # Errors
//
// When resolving a kind, (Generators).Generator returns ErrNotFound if the kind
// is missing from the registry.
package generator
