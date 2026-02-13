// Package cmd wires the service CLI commands for Bezeichner.
//
// This package is responsible for defining user-facing commands (via the
// go-service CLI framework) and connecting them to the application's dependency
// injection module graph.
//
// # Server command
//
// The primary command is "server", which starts the Bezeichner service and its
// transports. The command is registered by RegisterServer and is backed by the
// package-level Module, which composes all required sub-modules.
//
// The server is intended to be run by providing configuration via the standard
// go-service mechanisms (for example, using the "-i" flag to point at a config
// file). See internal/config for the configuration schema.
//
// This package does not implement business logic. Domain behavior lives in
// internal/api/ids, and transport bindings are wired under internal/api/v1.
package cmd
