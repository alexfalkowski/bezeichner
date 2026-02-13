![Gopher](assets/gopher.png)
[![CircleCI](https://circleci.com/gh/alexfalkowski/bezeichner.svg?style=shield)](https://circleci.com/gh/alexfalkowski/bezeichner)
[![codecov](https://codecov.io/gh/alexfalkowski/bezeichner/graph/badge.svg?token=TDRSV3MGSM)](https://codecov.io/gh/alexfalkowski/bezeichner)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexfalkowski/bezeichner)](https://goreportcard.com/report/github.com/alexfalkowski/bezeichner)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/bezeichner.svg)](https://pkg.go.dev/github.com/alexfalkowski/bezeichner)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# Bezeichner

Bezeichner is a small Go service that **generates** and **maps** identifiers, exposed via **gRPC** and **HTTP**.

- gRPC is the primary API surface.
- HTTP is implemented as an RPC gateway that routes by **gRPC full method name** (so both transports share the same contract).

The API contract lives in:
- `api/bezeichner/v1/service.proto`

## Why a service?

Distributed systems often need globally unique identifiers across multiple languages and runtimes. Bezeichner centralizes identifier generation so:
- you don't re-implement ID generation logic per service/language,
- you can standardize generator choices per domain/application,
- you can migrate/translate legacy identifiers via mapping.

## API Overview (v1)

The v1 service supports:
- `GenerateIdentifiers`: generate `count` identifiers for a configured `application`
- `MapIdentifiers`: map a list of identifiers using a configured mapping table

Both endpoints enforce request-size limits in the domain layer for basic DoS protection (limits are currently 1000 items for both generate count and map list).

## Configuration

Bezeichner uses the `go-service` configuration conventions. A representative configuration used by development and feature tests is:

- `test/.config/server.yml`

### Generator configuration

Generator configuration selects **applications**, each of which has:
- a `name` (the public application key you pass on requests),
- a `kind` (the generator implementation to use).

Supported built-in kinds (at time of writing):

- `uuid`
- `ksuid`
- `ulid`
- `xid`
- `snowflake`
- `nanoid`
- `typeid`
- `pg` (Postgres sequence-backed)

Example:

```/dev/null/generator.yml#L1-21
generator:
  applications:
    - name: public-uuid
      kind: uuid
    - name: internal-ulid
      kind: ulid
    - name: order-seq
      kind: pg
```

#### Postgres generator (`kind: pg`)

The `pg` generator reads the next value from a Postgres sequence using:

- `SELECT nextval($1::regclass)`

The sequence name is the configured `application.name` (i.e., `Application.Name`).

Important:
- The service **does not create sequences**. You must provision them separately (migrations/DB tooling/etc.).

### Mapper configuration

Mapper configuration defines a lookup table for identifier translation (useful for legacy migrations):

```/dev/null/mapper.yml#L1-6
mapper:
  identifiers:
    legacy-1: canonical-1
    legacy-2: canonical-2
```

Semantics:
- Mapping is strict: if any input ID is missing from the table, the operation fails.
- Output order matches input order.

### Health configuration

Health checks are provided via `go-health` integration. Timing is configured as durations:

```/dev/null/health.yml#L1-6
health:
  duration: 1s   # how often to run checks
  timeout:  1s   # max time a single check may take
```

The service registers:
- `noop` and `online` checks always,
- a `pg` check only if a DB handle is configured/available.

## Running

### Local dev (hot reload)

```/dev/null/commands.sh#L1-3
make submodule
make dep
make dev
```

`make dev` runs the server using `air` and a config file like:

- `./bezeichner server -i file:test/.config/server.yml`

### Build

```/dev/null/commands.sh#L1-4
make build        # builds ./bezeichner (release)
make build-test   # builds ./bezeichner test binary (features, race, coverage)
```

## Usage examples

Below are examples for both transports. Exact request/response schemas are defined in `api/bezeichner/v1/service.proto`.

### gRPC (grpcurl)

Assuming the service is listening on `localhost:12000` (default in the sample config):

Generate 3 IDs for application `public-uuid`:

```/dev/null/grpcurl.sh#L1-7
grpcurl -plaintext \
  -d '{"application":"public-uuid","count":"3"}' \
  localhost:12000 \
  bezeichner.v1.Service/GenerateIdentifiers
```

Map identifiers:

```/dev/null/grpcurl.sh#L1-7
grpcurl -plaintext \
  -d '{"ids":["legacy-1","legacy-2"]}' \
  localhost:12000 \
  bezeichner.v1.Service/MapIdentifiers
```

### HTTP RPC gateway (curl)

HTTP routes are keyed by the **gRPC full method name**. That means your HTTP client calls the same method identifiers as gRPC.

Assuming the service is listening on `localhost:11000` (default in the sample config):

Generate identifiers:

```/dev/null/curl.sh#L1-7
curl -sS \
  -X POST \
  -H 'content-type: application/json' \
  --data '{"application":"public-uuid","count":"3"}' \
  http://localhost:11000/bezeichner.v1.Service/GenerateIdentifiers
```

Map identifiers:

```/dev/null/curl.sh#L1-7
curl -sS \
  -X POST \
  -H 'content-type: application/json' \
  --data '{"ids":["legacy-1","legacy-2"]}' \
  http://localhost:11000/bezeichner.v1.Service/MapIdentifiers
```

Note:
- The exact HTTP path shape is defined by the underlying `go-service` HTTP RPC router; the important part is that routing is done by gRPC full method name.

## Deployment guidance

Bezeichner is typically deployed as a shared internal service. Depending on your scale and domain boundaries, you can:
- run a single global instance,
- shard by bounded context,
- run per region/cluster.

If you use the `pg` generator, ensure database connectivity and sequence provisioning are handled as part of your infrastructure/migrations.

## Design & dependencies

Bezeichner builds on established ID generation libraries:
- https://github.com/google/uuid
- https://github.com/segmentio/ksuid
- https://github.com/oklog/ulid
- https://github.com/rs/xid
- https://github.com/sony/sonyflake
- https://go.jetify.com/typeid

Service scaffolding and transport/DI patterns:
- https://github.com/alexfalkowski/go-service/v2

## Development

### Repository structure

The project follows:
- https://github.com/golang-standards/project-layout

### Requirements

- Go (see `go.mod` for version)
- Ruby (used for end-to-end feature tests; see `.ruby-version`)

### Setup

Most `make` targets come from the `bin/` git submodule:

```/dev/null/setup.sh#L1-3
make submodule
make dep
make setup
```

### Tests

Go unit/spec tests:

```/dev/null/tests.sh#L1-2
make specs
make lint
```

End-to-end feature tests:

```/dev/null/tests.sh#L1-1
make features
```

### Changelog

See `CHANGELOG.md`.