# Example Go monolith with embedded microservices and The Clean Architecture

[![PkgGoDev](https://pkg.go.dev/badge/github.com/powerman/go-monolith-example)](https://pkg.go.dev/github.com/powerman/go-monolith-example)
[![Go Report Card](https://goreportcard.com/badge/github.com/powerman/go-monolith-example)](https://goreportcard.com/report/github.com/powerman/go-monolith-example)
[![CI/CD](https://github.com/powerman/go-monolith-example/workflows/CI/CD/badge.svg?event=push)](https://github.com/powerman/go-monolith-example/actions?query=workflow%3ACI%2FCD)
[![Coverage Status](https://coveralls.io/repos/github/powerman/go-monolith-example/badge.svg?branch=master)](https://coveralls.io/github/powerman/go-monolith-example?branch=master)
[![Project Layout](https://img.shields.io/badge/Standard%20Go-Project%20Layout-informational)](https://github.com/golang-standards/project-layout)
[![Release](https://img.shields.io/github/v/release/powerman/go-monolith-example)](https://github.com/powerman/go-monolith-example/releases/latest)

This project shows an example of how to implement monolith with embedded
microservices (a.k.a. modular monolith). This way you'll get many upsides
of monorepo without it complexity and at same time most of upsides of
microservice architecture without some of it complexity.

The embedded microservices use Uncle Bob's "Clean Architecture", check
[Example Go microservice](https://github.com/powerman/go-service-example)
for more details.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Overview](#overview)
  - [Structure of Go packages](#structure-of-go-packages)
  - [Features](#features)
- [Development](#development)
  - [Requirements](#requirements)
  - [Setup](#setup)
    - [HTTPS](#https)
  - [Usage](#usage)
    - [Cheatsheet](#cheatsheet)
- [Run](#run)
  - [Docker](#docker)
  - [Source](#source)
- [TODO](#todo)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Overview

### Structure of Go packages

- `api/*` - definitions of own and 3rd-party (in `api/ext-*`)
  APIs/protocols and related auto-generated code
- `cmd/*` - main application(s)
- `internal/*` - packages shared by embedded microservices, e.g.:
  - `internal/config` - configuration (default values, env) shared by
    embedded microservices' subcommands and tests
  - `internal/dom` - domain types shared by microservices (Entities)
- `ms/*` - embedded microservices, with structure:
  - `internal/config` - configuration(s) (default values, env, flags) for
    microservice's subcommands and tests
  - `internal/app` - define interfaces ("ports") for The Clean
    Architecture (or "Ports and Adapters" architecture) and implements
    business-logic
  - `internal/srv/*` - adapters for served APIs/UI
  - `internal/sub` - adapter for incoming events
  - `internal/dal` - adapter for data storage
  - `internal/migrations` - DB migrations (in both SQL and Go)
  - `internal/svc/*` - adapters for accessing external services
- `pkg/*` - helper packages, not related to architecture and
  business-logic (may be later moved to own modules and/or replaced by
  external dependencies), e.g.:
  - `pkg/def/` - project-wide defaults
- `*/old/*` - contains legacy code which shouldn't be modified - this code
  is supposed to be extracted from `old/` directories (and refactored to
  follow Clean Architecture) when it'll need any non-trivial modification
  which require testing

### Features

- [X] Project structure (mostly) follows
  [Standard Go Project Layout](https://github.com/golang-standards/project-layout).
- [X] Strict but convenient golangci-lint configuration.
- [X] Embedded microservices:
  - [X] Well isolated from each other.
  - [X] Can be easily extracted from monolith into separate projects.
  - [X] Share common configuration (both env vars and flags).
  - [X] Each has own CLI subcommands, DB migrations, ports, metrics, …
- [X] Easily testable code (thanks to The Clean Architecture).
- [X] Avoids (and resists to) using global objects (to ensure embedded
  microservices won't conflict on these global objects).
- [X] CLI subcommands support using [cobra](https://github.com/spf13/cobra).
- [X] Graceful shutdown support.
- [X] Configuration defaults can be overwritten by env vars and flags.
- [X] Example JSON-RPC 2.0 over HTTP API, with CORS support.
- [X] Example gRPC API:
  - [X] External and internal APIs on different host/port.
  - [X] gRPC services with and without token-based authentication.
  - [X] API design (mostly) follows
    [Google API Design Guide](https://cloud.google.com/apis/design) and
    [Google API Improvement Proposals](https://google.aip.dev/).
- [X] Example OpenAPI 2.0 using grpc-gateway, with CORS suport:
  - [X] Access to gRPC using HTTP/1 (except bi-directional streaming).
  - [X] Generates `swagger.json` from gRPC `.proto` files.
  - [X] Embedded [Swagger UI](https://swagger.io/tools/swagger-ui/).
- [X] Example DAL (data access layer):
  - [X] MySQL 5.7 (strictest SQL mode).
  - [X] PostgreSQL 11 (secure schema usage pattern).
- [X] Example tests, both unit and integration.
- [X] Production logging using [structlog](https://github.com/powerman/structlog).
- [X] Production metrics using Prometheus.
- [X] Docker and docker-compose support.
- [X] Smart test coverage report, with optional support for coveralls.io.
- [X] Linters for Dockerfile and shell scripts.
- [X] CI/CD setup for GitHub Actions and CircleCI.

## Development

### Requirements

- Go 1.16
- [Docker](https://docs.docker.com/install/) 19.03+
- [Docker Compose](https://docs.docker.com/compose/install/) 1.25+

### Setup

1. After cloning the repo copy `env.sh.dist` to `env.sh`.
2. Review `env.sh` and update for your system as needed.
3. It's recommended to add shell alias `alias dc="if test -f env.sh; then
   source env.sh; fi && docker-compose"` and then run `dc` instead of
   `docker-compose` - this way you won't have to run `source env.sh` after
   changing it.

#### HTTPS

1. This project requires https:// and will send HSTS and CSP HTTP headers,
   and also it uses gRPC with authentication which also require TLS certs,
   so you'll need to create certificate to run it on localhost - follow
   instructions in [Create local CA to issue localhost HTTPS
   certificates](https://gist.github.com/powerman/2fc4b1a5aee62dd9491cee7f75ead0b4).
2. Or you can just use certificates in `configs/insecure-dev-pki`, which
   was created this way:

```
$ . ./env.sh   # Sets $EASYRSA_PKI=configs/insecure-dev-pki.
$ /path/to/easyrsa init-pki
$ echo Dev CA $(go list -m) | /path/to/easyrsa build-ca nopass
$ /path/to/easyrsa --days=3650 "--subject-alt-name=DNS:postgres" build-server-full postgres nopass
$ /path/to/easyrsa --days=3650 "--subject-alt-name=DNS:localhost" build-server-full ms-auth nopass
$ /path/to/easyrsa --days=3650 "--subject-alt-name=IP:127.0.0.1" build-server-full ms-auth-int nopass
```

### Usage

To develop this project you'll need only standard tools: `go generate`,
`go test`, `go build`, `docker build`. Provided scripts are for
convenience only.

- Always load `env.sh` *in every terminal* used to run any project-related
  commands (including `go test`): `source env.sh`.
    - When `env.sh.dist` change (e.g. by `git pull`) next run of `source
    env.sh` will fail and remind you to manually update `env.sh` to match
    current `env.sh.dist`.
- `go generate ./...` - do not forget to run after making changes related
  to auto-generated code
- `go test ./...` - test project (excluding integration tests), fast
- `./scripts/test` - thoroughly test project, slow
- `./scripts/test-ci-circle` - run tests locally like CircleCI will do
- `./scripts/cover` - analyse and show coverage
- `./scripts/build` - build docker image and binaries in `bin/`
  - Then use mentioned above `dc` (or `docker-compose`) to run and control
    the project.
    - Access project at host/port(s) defined in `env.sh`.

#### Cheatsheet

```sh
dc up -d --remove-orphans               # (re)start all project's services
dc logs -f -t                           # view logs of all services
dc logs -f SERVICENAME                  # view logs of some service
dc ps                                   # status of all services
dc restart SERVICENAME
dc exec SERVICENAME COMMAND             # run command in given container
dc stop && dc rm -f                     # stop the project
docker volume rm PROJECT_SERVICENAME    # remove some service's data
```

It's recommended to avoid `docker-compose down` - this command will also
remove docker's network for the project, and next `dc up -d` will create a
new network… repeat this many enough times and docker will exhaust
available networks, then you'll have to restart docker service or reboot.

## Run

### Docker

```
$ docker run -i -t --rm ghcr.io/powerman/go-monolith-example:0.2.0 -v
mono version v0.2.0 7562a1e 2020-10-22_03:12:04 go1.15.3
```

### Source

#### Run directly, without building

```bash
# because golang allows to run a go file directly
# cmd/mono/main.go is the entry point with the `main` function
go run cmd/mono/main.go serve
```

#### Build first, then run

We can either build the binary only, or build the binary and then dockerize it.

In this example below, we demonstrate using the `Taskfile` command to build our binary, then, run our built `mono` binary.

```bash
# build binary only
# our binary gets installed into the ./bin/ folder, as `mono`.
$ task scripts:build:binary

# so now, we can just run the built `mono` binary.
$ ./bin/mono -h
Example monolith with embedded microservices

Usage:
  mono [flags]
  mono [command]

Available Commands:
  help        Help about any command
  ms          Run given embedded microservice's command
  serve       Starts embedded microservices

Flags:
  -h, --help                    help for mono
      --log.level OneOfString   log level [debug|info|warn|err] (default debug)
  -v, --version                 version for mono

Use "mono [command] --help" for more information about a command.

$ ./bin/mono serve -h
Starts embedded microservices

Usage:
  mono serve [flags]

Flags:
  -h, --help                                  help for serve
      --host NotEmptyString                   host to serve (default home)
      --host-int NotEmptyString               internal host to serve (default home)
      --mono.port Port                        port to serve monolith introspection (default 17000)
      --nats.urls NotEmptyString              URLs to connect to NATS (separated by comma) (default nats://localhost:34222)
      --stan.cluster_id NotEmptyString        STAN cluster ID (default local)
      --timeout.shutdown Duration             must be less than 10s used by 'docker stop' between SIGTERM and SIGKILL (default 9s)
      --timeout.startup Duration              must be less than swarm's deploy.update_config.monitor (default 3s)

Global Flags:
      --log.level OneOfString   log level [debug|info|warn|err] (default debug)

$ ./bin/mono -v
mono version v0.2.0 7562a1e 2020-10-22_03:19:37 go1.15.3

$ ./bin/mono serve
         mono: inf      main: `started` version v0.2.0 7562a1e 2020-10-22_03:19:37
         mono: inf     serve: `serve` home:17000 [monolith introspection]
      
^C
         mono: inf     serve: `shutdown` [monolith introspection]
         mono: inf      main: `finished` version v0.2.0 7562a1e 2020-10-22_03:19:37
```

## TODO

- [ ] Add security-related headers for HTTPS endpoints (HSTS, CSP, etc.),
  also move default host from localhost to avoid poisoning it with HSTS.
- [ ] Embed https://github.com/powerman/go-service-example as an example
  of embedding microservices from another repo.
- [ ] Add example of `internal/svc/*` adapters calling some other services.
- [ ] Add LPC (local procedure call API between embedded microservices),
  probably using https://github.com/fullstorydev/grpchan.
- [ ] Add complete CRUD example as per Google API Design Guide (with
  PATCH/FieldMask), probably with generation of models conversion code using
  https://github.com/bold-commerce/protoc-gen-struct-transformer.
- [ ] Add NATS/STAN publish/subscribe example in `internal/sub`
  (or maybe use JetStream instead of STAN?).
- [ ] Switch from github.com/lib/pq to github.com/jackc/pgx.
