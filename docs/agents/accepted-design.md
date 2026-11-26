# Accepted design

This index records Bezeichner's accepted design choices, tradeoffs and
dependency limitations. The area files carry the same weight as
[AGENTS.md](../../AGENTS.md).

Before recording a finding, audit entry or proposed change, read the relevant
area files and account for each decision's scope. Record owner-accepted
decisions in the relevant area file.

- [Generators](accepted-design/generators.md) covers the TypeID application-name
  contract and Sonyflake machine IDs in Kubernetes.
- [Transport](accepted-design/transport.md) covers HTTP RPC routing and the
  external connectivity check used by `healthz`.
