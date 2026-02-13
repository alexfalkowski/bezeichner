// Package health integrates service health checks and observers for Bezeichner.
//
// This package wires the go-health server into the application and registers a
// set of standard health checks for both HTTP and gRPC surfaces.
//
// # What gets registered
//
// Register installs a set of health registrations under two service names:
//
//   - The process/service name (from env.Name), which is used for HTTP observers.
//   - The gRPC service name (from the generated protobuf descriptor), which is
//     used for gRPC observers.
//
// The registrations include:
//
//   - "noop":   a checker that always reports healthy.
//   - "online": an "online" registration that reflects the server's online state.
//   - "pg":     a database health check when a database handle is available.
//
// Database checking is conditional: if the DB dependency is not present, the
// "pg" registration is not added.
//
// # Timing configuration
//
// Health timings are configured using Config:
//
//   - Timeout:  maximum duration a single check may take.
//   - Duration: interval/frequency used when scheduling checks.
//
// Both values are parsed as durations (for example "250ms", "1s", "5m").
//
// # Observers
//
// In addition to check registrations, this package also installs observers that
// expose the health state via:
//
//   - HTTP endpoints: "healthz", "livez", and "readyz" (under the env.Name scope)
//   - gRPC observer: "grpc" (under the protobuf service name scope)
//
// The exact endpoint routing is handled by the underlying health/server package;
// this package only declares what to observe and which checks each observer
// depends on.
package health
