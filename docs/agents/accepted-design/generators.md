# Accepted design for generators

These entries are part of the [accepted-design index](../accepted-design.md).

## TypeID application-name contract

Bezeichner owns the `typeid` adapter around `go.jetify.com/typeid`. The adapter
uses the configured application name as the native TypeID prefix. Operators
must supply a valid TypeID prefix as the application name for `kind: typeid`.
The adapter enforces this contract when generating an ID. Raise prefix
validation concerns only when the task explicitly concerns changing the TypeID
naming contract, making invalid configuration fail earlier, or documenting
application-name restrictions.

[TypeID implementation](../../../internal/generator/typeid.go)

## Snowflake machine IDs in Kubernetes

The `snowflake` generator uses Sonyflake defaults. In the intended Kubernetes
deployment, concurrently running pods use normal pod networking, and the
default private IPv4-derived machine ID is an accepted source of uniqueness.
Raise machine-ID collision risk only when the task explicitly concerns local
multi-process deployments, `hostNetwork`, overlapping pod CIDRs, multi-cluster
shared ID spaces, IPv6-only environments, environments without private IPv4
addresses, or changing the deployment topology or ID contract.

[Snowflake implementation](../../../internal/generator/snowflake.go)
