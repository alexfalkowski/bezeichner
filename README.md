![Gopher](assets/gopher.png)
[![CircleCI](https://circleci.com/gh/alexfalkowski/bezeichner.svg?style=shield)](https://circleci.com/gh/alexfalkowski/bezeichner)
[![codecov](https://codecov.io/gh/alexfalkowski/bezeichner/graph/badge.svg?token=TDRSV3MGSM)](https://codecov.io/gh/alexfalkowski/bezeichner)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexfalkowski/bezeichner)](https://goreportcard.com/report/github.com/alexfalkowski/bezeichner)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/bezeichner.svg)](https://pkg.go.dev/github.com/alexfalkowski/bezeichner)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# 🏷️ Bezeichner

Bezeichner is a small Go service that **generates** and **maps** identifiers, exposed via **gRPC** and **HTTP**.

- gRPC is the primary API surface.
- HTTP is implemented as an RPC gateway that routes by **gRPC full method name** (so both transports share the same contract).

The API contract lives in:

- `api/bezeichner/v1/service.proto`

## 🤔 Why a service?

Distributed systems often need globally unique identifiers across multiple languages and runtimes. Bezeichner centralizes identifier generation so:

- you don't re-implement ID generation logic per service/language,
- you can standardize generator choices per domain/application,
- you can migrate/translate legacy identifiers via mapping.

## 🧭 API Overview (v1)

The v1 service supports:

- `GenerateIdentifiers`: generate `count` identifiers for a configured `application`
- `MapIdentifiers`: classify identifiers as mapped or unmapped for a configured `application`

Responses contain generated `ids` or mapped/unmapped identifiers, plus a `meta`
map reserved for transport/service metadata.

> [!NOTE]
> HTTP is not a separate REST API. It is an RPC gateway over the same protobuf service contract used by gRPC.
> [!WARNING]
> Both endpoints enforce request-size limits in the domain layer for basic DoS protection. By default, `GenerateIdentifiers.count` and the `MapIdentifiers.ids` list length are capped at `1000`; larger requests fail with `InvalidArgument`.

Unknown generator applications, unknown mapper applications, unresolved generator kinds, and omitted mapper configuration fail with `NotFound`.

## ⚙️ Configuration

Bezeichner uses the `go-service` configuration conventions. A representative configuration used by development and feature tests is:

- `test/.config/server.yml`

> [!TIP]
> Use `test/.config/server.yml` as the copy-paste source for local examples. The usage examples below use application and mapping names from that file.

The snippets below document the Bezeichner-owned blocks. A runnable service
configuration also includes shared `go-service` keys, such as `environment`,
`id`, `telemetry`, and `transport.http` / `transport.grpc` addresses. The sample
config binds HTTP to `tcp://:11000` and gRPC to `tcp://:12000`, which is what
the examples below assume.

Startup requires the top-level `health` and `generator` blocks, plus the shared
service configuration embedded by `go-service`. The `mapper` block is optional,
but omitting it makes every `MapIdentifiers` request fail with `NotFound`.

### 🧬 Generator configuration

Generator configuration selects **applications**, each of which has:

- a `name` (the public application key you pass on requests),
- a `kind` (the generator implementation to use).

Application entries must have non-empty `name` and `kind` values, and
application names must be unique. Malformed application lists fail config
validation during startup.

Generated identifiers are prefixed with the application name followed by `_`.
For example, an application named `uuid` returns identifiers like `uuid_<uuid>`.
The `typeid` generator uses the application name as the native TypeID prefix.
For `kind: typeid`, choose application names that are valid TypeID prefixes:
lowercase ASCII letters and underscores only (`[a-z_]`), at most 63 characters,
and not starting or ending with `_`.

Supported built-in kinds (at time of writing):

- `uuid` (application-prefixed UUIDv7 string)
- `ksuid` (application-prefixed KSUID string)
- `ulid` (application-prefixed ULID string)
- `xid` (application-prefixed XID string)
- `snowflake` (application-prefixed Sonyflake-based numeric ID as a decimal string)
- `nanoid` (application-prefixed NanoID string)
- `typeid` (TypeID string using the application name as the TypeID prefix)

Example:

```yaml
generator:
  applications:
    - name: uuid
      kind: uuid
    - name: ulid
      kind: ulid
```

### 🗺️ Mapper configuration

Mapper configuration defines application-scoped identifier translations (useful for legacy migrations):

```yaml
mapper:
  applications:
    - name: uuid
      identifiers:
        req1: resp1
        req2: resp2
```

> [!IMPORTANT]
> Mapping classifies each input ID. Known IDs are returned in `mapped`; missing IDs are returned in `unmapped`. Consumers decide whether unmapped IDs should be ignored, reported, retried, or treated as failures.
> The `mapper` block is optional at startup. If it is omitted, all `MapIdentifiers` requests fail with `NotFound`.

### 🚧 Request limit configuration

Request limits configure Bezeichner-owned per-request item caps:

```yaml
limits:
  generate_count: 1000
  map_ids: 1000
```

Both values are optional and default to `1000` when omitted. Set lower values
for smaller deployments or stricter abuse controls. Requests above the effective
limit fail with `InvalidArgument`.

### ❤️ Health configuration

Health checks are provided via `go-health` integration. Timing is configured as durations:

```yaml
health:
  duration: 1s   # how often to run checks
  timeout:  1s   # max time a single check may take
```

Both values must be positive durations.

The service registers:

- `noop` and `online` checks.

`healthz` intentionally uses the `online` check. Bezeichner follows the shared
service convention that all services should report whether they can reach the
outside world if they need public egress later, even when the current generate
and map paths do not require outbound network access.

The service exposes these health observers:

| Observer       | Check    |
| -------------- | -------- |
| HTTP `healthz` | `online` |
| HTTP `livez`   | `noop`   |
| HTTP `readyz`  | `noop`   |
| gRPC health    | `noop`   |

With the sample config, operational HTTP routes are service-prefixed:

| Endpoint           | Local path            |
| ------------------ | --------------------- |
| Health             | `/bezeichner/healthz` |
| Liveness           | `/bezeichner/livez`   |
| Readiness          | `/bezeichner/readyz`  |
| Prometheus metrics | `/bezeichner/metrics` |

The gRPC health service name is `bezeichner.v1.Service`.

## 🚀 Running

### ♻️ Local dev (hot reload)

After setup:

```sh
make dev
```

`make dev` runs the server using `air` with the test config:

- `cd test && ../bezeichner server -config file:.config/server.yml`

### 🏗️ Build

```sh
make build        # builds ./bezeichner (release)
make build-test   # builds ./bezeichner test binary (features, race, coverage)
```

## 🧪 Usage examples

Below are examples for both transports. Exact request/response schemas are defined in `api/bezeichner/v1/service.proto`.

### 🔌 gRPC (grpcurl)

Assuming the service is listening on `localhost:12000` (default in the sample config):

Generate 3 IDs for application `uuid`:

```sh
grpcurl -plaintext \
  -d '{"application":"uuid","count":"3"}' \
  localhost:12000 \
  bezeichner.v1.Service/GenerateIdentifiers
```

Map identifiers:

```sh
grpcurl -plaintext \
  -d '{"application":"uuid","ids":["req1","req3"]}' \
  localhost:12000 \
  bezeichner.v1.Service/MapIdentifiers
```

The response classifies inputs by original ID:

```json
{
  "meta": {
    "requestId": "...",
    "userAgent": "..."
  },
  "mapped": {
    "req1": "resp1"
  },
  "unmapped": ["req3"]
}
```

### 🌐 HTTP RPC gateway (curl)

HTTP routes are keyed by the **gRPC full method name**. That means your HTTP client calls the same method identifiers as gRPC.

Assuming the service is listening on `localhost:11000` (default in the sample config):

Generate identifiers:

```sh
curl -sS \
  -X POST \
  -H 'content-type: application/json' \
  --data '{"application":"uuid","count":3}' \
  http://localhost:11000/bezeichner.v1.Service/GenerateIdentifiers
```

Map identifiers:

```sh
curl -sS \
  -X POST \
  -H 'content-type: application/json' \
  --data '{"application":"uuid","ids":["req1","req2"]}' \
  http://localhost:11000/bezeichner.v1.Service/MapIdentifiers
```

HTTP errors use the same domain classification as gRPC, rendered as HTTP
statuses with safe `text/error` response bodies:

| gRPC/domain error | HTTP status | Common triggers                                                                                                |
| ----------------- | ----------- | -------------------------------------------------------------------------------------------------------------- |
| `InvalidArgument` | `400`       | `GenerateIdentifiers.count` or `MapIdentifiers.ids` exceeds the configured request limit                       |
| `NotFound`        | `404`       | unknown generator application, unknown mapper application, unresolved generator kind, or omitted mapper config |

> [!NOTE]
> The generated gRPC full method names include a leading slash, for example `/bezeichner.v1.Service/GenerateIdentifiers`. In HTTP URLs, that slash is the path separator after the host; `grpcurl` uses the `service/method` form without the leading slash.

## 🛡️ Deployment guidance

Bezeichner is typically deployed as a shared internal service. Depending on your scale and domain boundaries, you can:

- run a single global instance,
- shard by bounded context,
- run per region/cluster.

Non-master branch CI validates Docker images for both supported platforms. To
reproduce that locally, run:

```sh
make platform=amd64 test-docker
make platform=arm64 test-docker
```

Published Docker images use the `alexfalkowski/bezeichner` repository. Pin a
released version tag such as `alexfalkowski/bezeichner:<version>` for
deployments.

Docker release, manifest publication, and deploy targets are CI-owned workflows
that require release artifacts plus DockerHub or GitHub credentials. Use the
local `test-docker` targets for image validation before pushing changes.

> [!CAUTION]
> The `snowflake` generator uses Sonyflake defaults. The intended deployment assumes normal Kubernetes pod networking where each concurrently running pod has a suitable private IPv4-derived machine ID. Re-evaluate that assumption for local multi-process deployments, `hostNetwork`, overlapping pod CIDRs, multi-cluster shared ID spaces, IPv6-only environments, or environments without private IPv4 addresses.
> [!IMPORTANT]
> Deployments should pin released version tags instead of depending on the moving
> `latest` Docker manifest. The release pipeline may update `latest` after
> versioned images are published, but the versioned image tag is the deployment
> contract.

## 🔗 Design & dependencies

Bezeichner builds on established ID generation libraries:

- <https://github.com/alexfalkowski/go-service/tree/master/id>
- <https://github.com/sony/sonyflake>
- <https://go.jetify.com/typeid>

Service scaffolding and transport/DI patterns:

- <https://github.com/alexfalkowski/go-service/v2>

## 🛠️ Development

### 📁 Repository structure

The project follows:

- <https://github.com/golang-standards/project-layout>

### 📋 Requirements

- Go (see `go.mod` for version)
- Ruby with Bundler (used for end-to-end feature tests; no repository-pinned Ruby version is currently declared)

### 📦 Setup

Initialize the `bin/` submodule and install dependencies:

```sh
git submodule sync
git submodule update --init
make dep
```

### ✅ Tests

Go unit/spec tests:

```sh
make specs
```

Lint checks:

```sh
make lint
```

End-to-end feature tests:

```sh
make features
```

`make features` starts its own test server from `test/nonnative.yml`, so the
sample HTTP and gRPC ports (`11000` and `12000`) must be free. Harness and server
logs are written under `test/reports/`.

End-to-end benchmark scenarios:

```sh
make benchmarks
```

### 🧬 Protobuf workflow

The protobuf contract in `api/bezeichner/v1/service.proto` owns the service API
schema. After changing it, regenerate and check the generated Go and Ruby stubs:

```sh
make proto-generate
make proto-stale
```

Before pushing API contract changes, also run:

```sh
make proto-breaking
```

### 🧰 Local CI checks

For a broader local pass before pushing, run the repository-owned checks that
mirror the main CI gates without publishing artifacts:

```sh
make lint
make proto-stale
make sec
make features
make benchmarks
make analyse
make coverage
```

`make proto-breaking` is part of the API-change workflow above because it checks
the protobuf contract against the remote master baseline. Codecov upload,
release, manifest, deploy, and push targets are CI or credential-gated.
