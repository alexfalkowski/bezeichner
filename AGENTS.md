# AGENTS.md

This repository contains **Bezeichner**, a Go service that generates and maps identifiers, exposed via **gRPC** and **HTTP**.

HTTP is implemented as an RPC gateway that routes by gRPC full method name.

## Shared skills

This repository uses the shared skills from `bin/skills/`. Read
`bin/AGENTS.md` for the canonical shared skill list and use the smallest
matching skill for the task.

## Repository at a glance

- Language: **Go** (`module github.com/alexfalkowski/bezeichner`, `go 1.25.0`; see `go.mod:1-3`).
- API contract: `api/bezeichner/v1/service.proto` (+ generated Go stubs under `api/` and Ruby stubs under `test/lib`).
- Build tooling: the root `Makefile` is a thin wrapper around the **`bin/` git submodule** (see `.gitmodules:1-3` and `Makefile:1-3`).

## First steps

### 1) Ensure submodules are present

Most `make` targets come from the `bin/` submodule.

```sh
make submodule
# expands to: git submodule sync && git submodule update --init
# see: bin/build/make/git.mak:29-31
```

### 2) Install dependencies

```sh
make dep
```

Observed behavior:
- Go deps are vendored (`go mod vendor`) and most Go build/test targets use `-mod vendor` (see `bin/build/make/grpc.mak:214`).
- Ruby deps for feature tests live under `test/` and are installed via bundler into `test/vendor/bundle` (see `bin/build/make/ruby.mak:15-21` and `test/.bundle/config:1-2`).

## Essential commands

### Build

```sh
make build        # builds ./bezeichner (release binary)
make build-test   # builds ./bezeichner test binary (-tags features, -race, coverage enabled)
```

Implementation: `bin/build/make/grpc.mak:208-215`.

### Test

```sh
make specs        # Go test suite (gotestsum + race + coverage)
make features     # end-to-end features under test/ (Ruby + nonnative harness)
make benchmarks   # Ruby benchmark harness
make coverage     # HTML + func coverage reports
```

Notes:
- `main_test.go` is guarded by build tag `features` (`main_test.go:1-10`).
- `make features` depends on `make build-test` (see `bin/build/make/grpc.mak:126-128`).

### Lint / format / security

```sh
make lint
make fix-lint
make format
make sec          # govulncheck
```

- Go linting is via `golangci-lint` plus formatting tools (gci/gofmt/gofumpt/goimports). Generated protobuf code is excluded (`.golangci.yml:33-43`).

### Protobuf (Buf)

```sh
make proto-lint
make proto-format
make proto-generate
make proto-breaking
```

Buf config is in `api/`:
- `api/buf.yaml`, `api/buf.gen.yaml`
- Ruby stubs are generated into `test/lib` (see `api/buf.gen.yaml:11-14`).

### Local dev

```sh
make dev
```

This uses `air` to rebuild/run the server:

```sh
./bezeichner server -i file:test/.config/server.yml
```

See `bin/build/make/grpc.mak:216-219`.

### Environment helpers

```sh
make start
make stop
```

These call scripts under `bin/build/docker/env` (see `bin/build/make/grpc.mak:261-267`).

## Code organization

### Entry points / CLI

- `main.go` constructs a `go-service/v2/cli` application and registers the `server` command (`main.go:9-15`).
- The server command is registered in `internal/cmd/server.go:9-12` and wires DI via `internal/cmd/module.go:13-19`.

### Dependency injection (DI)

The service uses `go-service/v2/di` modules:

- Root module composition: `internal/cmd/module.go:13-19`
  - `module.Server` (from `go-service/v2`)
  - `config.Module`, `health.Module`, `generator.Module`, `v1.Module`

The v1 module wires transports and the domain service:
- `internal/api/v1/module.go:11-16`.

### API layers

- Domain logic: `internal/api/ids/`
  - `Identifier.Generate` and `Identifier.Map` (`internal/api/ids/ids.go`).
  - Request-size limits are enforced here (`internal/api/ids/limits.go`).
- gRPC transport:
  - `internal/api/v1/transport/grpc/*` implements the protobuf service.
  - Errors are mapped to gRPC status codes in `internal/api/v1/transport/grpc/grpc.go:27-41`.
- HTTP transport:
  - `internal/api/v1/transport/http/http.go:9-13` routes HTTP RPC calls by gRPC full method name.

## Configuration

A representative config used by dev/feature tests is `test/.config/server.yml`.

Notable keys observed:
- `generator.applications[]`: generator applications (**name** and **kind**).
- `mapper.identifiers`: mapping table for `MapIdentifiers`.
- `transport.http.address` defaults to `tcp://:11000` and `transport.grpc.address` to `tcp://:12000`.

Notes:
- The generator application model in code currently includes only `name` and `kind` (no `prefix`, `suffix`, or `separator` fields).
- HTTP is an RPC gateway that routes by **gRPC full method name**, so the HTTP surface mirrors the gRPC contract in `api/bezeichner/v1/service.proto`.

## Generators

Generator implementations are in `internal/generator/*` and are selected by `Application.Kind`.
Standard kinds such as `uuid`, `ksuid`, `ulid`, `xid`, and `nanoid` are adapted
from `github.com/alexfalkowski/go-service/v2/id`.

- Registry: `internal/generator/generator.go:16-30`.
- Applications are defined via `internal/generator/config.go:3-25`.

### Snowflake generator deployment assumption

- Do **not** flag `internal/generator/snowflake.go` using Sonyflake default settings as a general bug.
- In the intended Kubernetes deployment, pods use normal pod networking, so Sonyflake's default private-IPv4-derived machine ID is an accepted uniqueness source for concurrently running pods.
- Only raise Snowflake machine-ID collision risk when the task explicitly concerns local multi-process deployments, `hostNetwork`, overlapping pod CIDRs, multi-cluster shared ID spaces, IPv6-only/no-private-IPv4 environments, or changing the deployment topology/ID contract.

## Request size limits (DoS protection)

Limits are enforced in the domain layer:

- `GenerateIdentifiers`: `count` is capped (`internal/api/ids/ids.go:33-36`, `internal/api/ids/limits.go:5-8`).
- `MapIdentifiers`: number of IDs is capped (`internal/api/ids/ids.go:62-65`).

These surface to clients as `InvalidArgument` via the gRPC error mapper (`internal/api/v1/transport/grpc/grpc.go:32-34`).

## Testing

### Go specs

`make specs` runs gotestsum with race and coverage (see `bin/build/make/grpc.mak:134-136`). Outputs land under `test/reports/`.

### Feature tests (Ruby / nonnative)

- Feature specs live under `test/features/**`.
- Harness config: `test/nonnative.yml`.
  - Launches `../bezeichner server -i file:.config/server.yml` (`test/nonnative.yml:6-12`).
- Cucumber report options: `test/.config/cucumber.yml:1`.

If bundler fails loading native gems (e.g., `json` extension), one observed fix is rebuilding the gem:

```sh
cd test && bundle pristine json
```

## Style / formatting

- Go files use tabs (per `.editorconfig:16-18`).
- Avoid hand-editing generated protobuf outputs under `api/**/*.pb.go` / `api/**/*_grpc.pb.go`; regenerate via `make proto-generate`.

## CI notes (CircleCI)

CircleCI runs (see `.circleci/config.yml`):

- `make dep`, `make lint`, `make proto-breaking`, `make sec`, `make trivy-repo`
- `make features`, `make benchmarks`, `make analyse`, `make coverage`, `make codecov-upload`
