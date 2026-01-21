# AGENTS.md

This repository contains **Bezeichner**, a Go service that generates and maps identifiers, exposed via **gRPC** and **HTTP** (HTTP routes map to RPC method names).

## Quick orientation

- Language: **Go** (module `github.com/alexfalkowski/bezeichner`, `go 1.25.0`; see `go.mod:1-3`).
- API contract: `api/bezeichner/v1/service.proto` (+ generated Go and Ruby stubs).
- Build tooling: `make` targets are largely defined in the **`bin/` submodule** (see `.gitmodules:1-3` and `Makefile:1-3`).

## First steps

### 1) Ensure submodules are present

The repo depends on the `bin/` git submodule for build scripts.

```sh
make submodule
# (equivalent to: git submodule sync && git submodule update --init)
```

### 2) Install dependencies

```sh
make dep
```

This runs Go module download/tidy/vendor and installs Ruby deps for `test/` via bundler (see `bin/build/make/go.mak` and `bin/build/make/grpc.mak`).

## Essential commands

All targets below are from `make help` (root).

### Build

```sh
make build        # builds a release binary named after the repo directory
make build-test   # builds a test binary with build tag `features`
```

Notes:
- `build-test` uses `-tags features` and `-mod vendor` (see `bin/build/make/grpc.mak:212-215`).

### Test

```sh
make specs        # Go test suite via gotestsum (race, coverage)
make features     # end-to-end feature tests under test/ (Ruby)
make benchmarks   # runs benchmarks (Ruby harness)
make coverage     # produces HTML + func coverage reports
```

### Lint / format

```sh
make lint
make fix-lint
make format
make sec          # govulncheck
```

Go linting is via `golangci-lint` with formatters enabled (gci/gofmt/gofumpt/goimports) and generated protobuf files excluded (`.golangci.yml:33-43`).

### Protobuf / Buf

```sh
make proto-lint
make proto-format
make proto-generate
make proto-breaking
```

Buf config lives under `api/`:
- `api/buf.yaml`
- `api/buf.gen.yaml` (generates Go stubs into `api/` and Ruby stubs into `test/lib`)

### Local dev

```sh
make dev
```

`dev` runs `air` to rebuild and run:

```sh
./bezeichner server -i file:test/.config/server.yml
```

(see `bin/build/make/grpc.mak:216-219`).

### Environment helpers

```sh
make start
make stop
```

These call scripts under `bin/build/docker/env` (see `bin/build/make/grpc.mak:261-267`).

## Repository structure

High-level layout (observed):

- `main.go` / `internal/cmd/*`: CLI entrypoint. The server command is registered as `bezeichner server` (see `internal/cmd/server.go:9-12`).
- `internal/config/*`: service configuration wrapper; embeds `go-service/v2/config.Config` and exposes feature configs (see `internal/config/config.go:10-16`).
- `internal/generator/*`: identifier generators (uuid/ksuid/ulid/xid/snowflake/nanoid/typeid/pg).
- `internal/mapper/*`: request→response identifier mapping config.
- `internal/api/*`:
  - `internal/api/ids`: domain logic for Generate/Map.
  - `internal/api/v1/transport/grpc`: gRPC server implementation and error mapping.
  - `internal/api/v1/transport/http`: HTTP routing for RPC endpoints.
- `api/bezeichner/v1/*`: protobuf + generated Go code (`*.pb.go`, `*_grpc.pb.go`).
- `test/`: Ruby-based feature test harness + generated Ruby protobuf stubs.

### Dependency injection / modules

The service uses `go-service/v2` DI helpers (`di.Module`, `di.Constructor`, `di.Register`). Modules are composed in `internal/cmd/module.go`:

- `module.Server` (from `go-service/v2`)
- `config.Module`, `health.Module`, `generator.Module`, `v1.Module`

When adding a new capability:
- add configuration structs in `internal/*/config.go` as needed,
- wire via a `Module` variable using `di.Module(...)`,
- register transport handlers under `internal/api/v1`.

## Configuration

A representative server config used by dev/test is `test/.config/server.yml`.

Notable keys observed:
- `generator.applications[]`: configured generators by name/kind/prefix/suffix/separator.
- `mapper.identifiers`: mapping table for `MapIdentifiers`.
- `sql.pg.masters[].url`: points to `file:secrets/pg` under `test/`.
- `transport.http.address` defaults to `tcp://:11000` and `transport.grpc.address` to `tcp://:12000`.

### Postgres generator gotcha

The `pg` generator expects a **sequence named after the application** (see README `README.md:84-88` and implementation `internal/generator/pg.go:17-27`). The service does not create sequences.

## Testing details

### Go tests

- Unit/spec tests run via `make specs` (gotestsum + `-race`, `-mod vendor`).
- `main_test.go` is guarded by the build tag `features` (`main_test.go:1`). Some targets/builds use `-tags features`.

Outputs are written under `test/reports/` (e.g., `test/reports/specs.xml`, coverage files).

### Feature tests (Ruby)

Feature specs are in `test/features/**` (e.g., gRPC/HTTP API feature files). The harness config is `test/nonnative.yml`, which starts the built Go binary (`../bezeichner`) with:

```yaml
parameters: ["-i file:.config/server.yml"]
```

It also defines a Postgres service proxy used for fault injection (see `test/nonnative.yml:17-28`).

## CI notes (CircleCI)

CircleCI runs (see `.circleci/config.yml`):

- `make dep`, `make lint`, `make proto-breaking`, `make sec`, `make trivy-repo`
- `make features`, `make benchmarks`, `make analyse`, `make coverage`, `make codecov-upload`

The build job uses:
- `postgres:18-trixie` (DB `test` / user `test` / password `test`)
- `grafana/mimir:latest` for metrics-related dependencies

## Style / formatting

- Go files use tabs (per `.editorconfig:16-18`).
- Lint runs `golangci-lint` and excludes generated protobuf code paths (`.golangci.yml:33-43`).

## Common pitfalls

- **Submodule required**: `make` targets include files from `bin/`; update/init it before running targets.
- **Generated protobuf code**: avoid hand-editing `api/**/*.pb.go` / `*_grpc.pb.go`; regenerate via `make proto-generate`.
- **Build tag `features`**: some tests/targets require `-tags features` (e.g., `make build-test`).
