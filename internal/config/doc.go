// Package config defines the top-level configuration schema for the Bezeichner
// service and provides dependency-injection wiring for configuration.
//
// This package is intentionally small and focused:
//
//   - It declares Config, the root configuration struct that aggregates
//     configuration for sub-systems (health, generator, mapper) and embeds the
//     shared service config type from go-service.
//   - It exposes a DI module that loads configuration from the configured source
//     (file/env/etc. as determined by the runtime) and makes typed sub-configs
//     available to other modules.
//
// # Structure
//
// Config is the service's root configuration. It includes pointers to:
//   - internal/health.Config
//   - internal/generator.Config
//   - internal/mapper.Config
//
// and embeds *go-service/v2/config.Config for common service settings.
//
// # Dependency injection
//
// The package-level Module wires configuration loading and provides constructors
// for the sub-config pointers so other packages can depend on their specific
// configuration type rather than the whole Config.
package config
