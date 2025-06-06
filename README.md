![Gopher](assets/gopher.png)
[![CircleCI](https://circleci.com/gh/alexfalkowski/bezeichner.svg?style=shield)](https://circleci.com/gh/alexfalkowski/bezeichner)
[![codecov](https://codecov.io/gh/alexfalkowski/bezeichner/graph/badge.svg?token=TDRSV3MGSM)](https://codecov.io/gh/alexfalkowski/bezeichner)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexfalkowski/bezeichner)](https://goreportcard.com/report/github.com/alexfalkowski/bezeichner)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/bezeichner.svg)](https://pkg.go.dev/github.com/alexfalkowski/bezeichner)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# Bezeichner

Bezeichner takes care of identifiers used in your services.

## Background

Identifiers are used everywhere and very important. There are many ways to generate one and we take inspiration from the following [design](https://www.linkedin.com/posts/alexxubyte_systemdesign-coding-interviewtips-activity-6976203240094736387-hvMT?utm_source=share&utm_medium=member_ios).

We don't have a preferred method. We just want to provide you with the best option.

### Why a service?

Lot's of distributed systems need global unique IDs. Since you are more than likely going to use microservices we don't need to reinvent the wheel for every language you use. Just use the service!

## Server

The server is defined by the following [proto contract](api/bezeichner/v1/service.proto). So each version of the service will have a new contract.

### Generator

This system allows you to configure an application with a generator.

To configure we just need the have the following configuration:

```yaml
generator:
  applications:
    - name: uuid
      kind: uuid
      prefix: prefix
      suffix: suffix
      separator: "-"
    - name: ksuid
      kind: ksuid
      prefix: prefix
      suffix: suffix
      separator: "-"
    - name: ulid
      kind: ulid
      prefix: prefix
      suffix: suffix
      separator: "-"
    - name: xid
      kind: xid
      prefix: prefix
      suffix: suffix
      separator: "-"
    - name: snowflake
      kind: snowflake
      prefix: prefix
      suffix: suffix
      separator: "-"
    - name: nanoid
      kind: nanoid
      prefix: prefix
      suffix: suffix
      separator: "-"
    - name: typeid
      kind: typeid
      prefix: prefix
      suffix: suffix
      separator: "-"
    - name: pg
      kind: pg
      prefix: prefix
      suffix: suffix
      separator: "-"
```

Each generator has the following properties:
- A distinct name.
- The kind of generator (uuid, ksuid, ulid, xid, snowflake, nanoid, typeid, pg).
- The prefix of the identifier.
- The suffix of the identifier.
- The separator used between the prefix and suffix.

#### Postgres

The postgres kind expects a sequence named after the application. The service does not create one. So you would need to use a [migration](https://github.com/alexfalkowski/migrieren) service.

### Mapper

The system allows you to map to different identifiers. This allows you to deal with legacy identifiers.

To configure we just need the have the following configuration:

```yaml
mapper:
  identifiers:
    req1: resp1
    req2: resp2
```

### Health

The system defines a way to monitor all of it's dependencies.

To configure we just need the have the following configuration:

```yaml
health:
  duration: 1s (how often to check)
  timeout: 1s (when we should timeout the check)
```

### Deployment

Since we are advocating building microservices, you would normally use a [container orchestration system](https://newrelic.com/blog/best-practices/container-orchestration-explained) and have a global service or shard these services per [bounded context](https://martinfowler.com/bliki/BoundedContext.html).

### Design

The service uses the awesome work of others. You can check out:
- https://github.com/segmentio/ksuid
- https://github.com/google/uuid
- https://github.com/oklog/ulid
- https://github.com/rs/xid
- https://github.com/sony/sonyflake
- https://github.com/jetpack-io/typeid-go
- https://github.com/alexfalkowski/go-service/v2

## Client

The client can be used in other projects. This is configured as follows:

```yaml
client:
  v1:
    host: server_host
    timeout: 1s
```

### Dependencies

![Dependencies](./assets/client.png)

## Development

If you would like to contribute, here is how you can get started.

### Structure

The project follows the structure in [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

### Dependencies

Please make sure that you have the following installed:
- [Ruby](.ruby-version)
- Golang

### Style

This project favours the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

### Setup

The get yourself setup, please run the following:

```sh
make setup
```

### Binaries

To make sure everything compiles for the app, please run the following:

```sh
make build-test
```

### Features

To run all the features, please run the following:

```sh
make features
```

### Changes

To see what has changed, please have a look at `CHANGELOG.md`
