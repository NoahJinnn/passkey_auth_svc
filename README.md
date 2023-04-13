<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [HQ Service modulith](#hq-service-modulith)
  - [Overview](#overview)
    - [Requirements](#requirements)
    - [Setup](#setup)
      - [Environment management](#environment-management)
      - [Naming convention](#naming-convention)
      - [HTTPS](#https)
    - [Usage](#usage)
      - [Cheatsheet](#cheatsheet)
  - [Run](#run)
    - [Docker](#docker)
      - [Run local PostgresSQL DB](#run-local-postgressql-db)
      - [Remove container storage](#remove-container-storage)
    - [Source](#source)
      - [Run directly, without building](#run-directly-without-building)
      - [Run with command args](#run-with-command-args)
      - [Build first, then run](#build-first-then-run)
  - [TODO](#todo)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# HQ Service modulith

## Overview
### Requirements

- Go 1.19
- [Docker](https://docs.docker.com/install/) 19.03+
- [Docker Compose](https://docs.docker.com/compose/install/) 1.25+

### Setup
#### Environment management

Setup Doppler:

```bash
task scripts:install:doppler
```

#### Naming convention

Variables required to run and test project.
Should be kept in sorted order.
Avoid referencing one variable from another if their order may change,
use lower-case variables defined above for such a shared values.
Naming convention:

```
<PROJECT>_<VAR>         - global vars, not specific for some embedded microservice (e.g. domain)
<PROJECT>_X_<SVC>_<VAR> - vars related to external services (e.g. databases)
<PROJECT>_<MS>_<VAR>    - vars related to embedded microservice (e.g. addr)
<PROJECT>__<MS>_<VAR>   - private vars for embedded microservice
```

#### HTTPS

1. This project requires https:// and will send HSTS and CSP HTTP headers,
   and also it uses gRPC with authentication which also require TLS certs,
   so you'll need to create certificate to run it on localhost - follow
   instructions in [Create local CA to issue localhost HTTPS
   certificates](https://gist.github.com/powerman/2fc4b1a5aee62dd9491cee7f75ead0b4).
2. Or you can just use certificates in `configs/dev-pki`, which
   was created this way:

```
$ /path/to/easyrsa init-pki
$ echo Dev CA $(go list -m) | /path/to/easyrsa build-ca nopass
$ /path/to/easyrsa --days=3650 "--subject-alt-name=DNS:postgres" build-server-full postgres nopass
$ /path/to/easyrsa --days=3650 "--subject-alt-name=DNS:localhost" build-server-full ms-hq nopass
$ /path/to/easyrsa --days=3650 "--subject-alt-name=IP:127.0.0.1" build-server-full ms-hq-int nopass
```

### Usage

To develop this project you'll need only standard tools: `go generate`,
`go test`, `go build`, `docker build` with `doppler run`

- `go generate ./...` - do not forget to run after making changes related
  to auto-generated code
- `doppler run -- go test ./...` - test project (excluding integration tests), fast
- `./scripts/test` - thoroughly test project, slow
- `./scripts/test-ci-circle` - run tests locally like CircleCI will do
- `./scripts/cover` - analyse and show coverage
- `./scripts/build` - build docker image and binaries in `bin/`
  - Then use mentioned above `dc` (or `docker-compose`) to run and control
    the project.
  - Access project at host/port(s) defined in `doppler`.

#### Cheatsheet

```sh
doppler run -- dc up -d --remove-orphans               # (re)start all project's services
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

#### Run local PostgresSQL DB

```bash
doppler run -- docker-compose up -d --remove-orphans
```

#### Remove container storage

```bash
docker-compose stop && docker-compose rm -f
docker volume rm hqservice_postgres
```

### Source

#### Run directly, without building

```bash
# cmd/hq/main.go is the entry point with the `main` function
task scripts:run
```

#### Run with command args
```bash
# cmd/hq/main.go is the entry point with the `main` function
task scripts:run -- --port 17002 --wa.id example --wa.origins https://example.com  # Specific auth service running on port `17002` with webauthn ID equals `example`; webauthn origns equals `https://example.com`  
```
#### Build first, then run

In this example below, we demonstrate using the `Taskfile` command to build our binary, then, run our built `hq` binary.

```bash
# build binary only
# our binary gets installed into the ./bin/ folder, as `hq`.
$ task scripts:build:binary

# so now, we can just run the built `hq` binary.
$ ./bin/hq -h
Example monolith with embedded microservices

Usage:
  hq [flags]
  hq [command]

Available Commands:
  help        Help about any command
  ms          Run given embedded microservice's command
  serve       Starts embedded microservices

Flags:
  -h, --help                    help for hq
      --log.level OneOfString   log level [debug|info|warn|err] (default debug)
  -v, --version                 version for hq

Use "hq [command] --help" for more information about a command.

$ ./bin/hq serve -h
Starts embedded microservices

Usage:
  hq serve [flags]

Flags:
  -h, --help                        help for serve
      --host-int NotEmptyString     internal host to serve (default Trans-MacBook-Pro-2.local)
      --hq.port Port                port to serve monolith introspection (default 17000)
      --timeout.shutdown Duration   must be less than 10s used by 'docker stop' between SIGTERM and SIGKILL (default 9s)
      --timeout.startup Duration    must be less than swarm's deploy.update_config.monitor (default 3s)

Global Flags:
      --log.level OneOfString   log level [debug|info|warn|err] (default debug)

$ ./bin/hq -v
hq version v0.2.0 7562a1e 2020-10-22_03:19:37 go1.15.3

$ ./bin/hq serve
         hq: inf      main: `started` version f/design-task-command-to-run-hqservice 51adc59-dirty 2023-02-15_09:36:06
```

## TODO

Functionality Group 1: add/connect assets and debts

- [x] Plaid aggregator with dev env
- [ ] Plaid aggregator with stg, prd env
- [ ] Implement [webauthn](https://github.com/go-webauthn/webauthn) API
- [ ] Implement [Lago](https://www.getlago.com/resources/compare/lago-vs-stripe) for billing service
- [ ] Implement authorization with `casbin`
- [x] Create `User` table
- [ ] Create CRUD REST API for `User` model
- [ ] Create asset tables based on Kubera features
- [ ] Create CRUD REST API for asset models
- [ ] Integration test for `auth` svc APIs

Functionality Group 2: Recap feature (‘reflections’)

- [ ] Create DB models: 1. Asset 2. Cashflow 3. Indices 4. IRR 5. Reflections
- [ ] Create CRUD REST API for all types of model
- [ ] Integration test APIs

Functionality Group 3: Insurance

- [ ] Create DB model for Insurance to store Insurance providers information, link to static assets
- [ ] Create CRUD REST API for Insurance model
- [ ] Integration test APIs
      Functionality Group 4: Safety Deposit Box
- Need to discuss
  Functionality Group 5: Beneficiary
- Need to discuss
