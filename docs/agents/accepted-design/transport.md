# Accepted design for transport

These entries are part of the [accepted-design index](../accepted-design.md).

## HTTP RPC routing

HTTP exposes the v1 protobuf service using RPC routes keyed by generated gRPC
full method names. Use that shared contract when reviewing HTTP routing and
request and response behavior.

[HTTP route registration](../../../internal/api/v1/transport/http/http.go) ·
[v1 service contract](../../../api/bezeichner/v1/service.proto)

## Health observers

HTTP `healthz` intentionally uses the shared `online` check to report external
connectivity, even though identifier generation and mapping need no outbound
network access. HTTP `livez` and `readyz` and the gRPC health observer use the
`noop` check. Revisit the `healthz` choice when the task concerns changing the
health contract, deployment without public egress, or a documented operator
workflow that requires local-process-only health.

[Health observer registration](../../../internal/health/health.go)
