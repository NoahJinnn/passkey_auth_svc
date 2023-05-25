<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [HQ Service modulith](#hq-service-modulith)
  - [Local development](#local-development)
    - [Requirements](#requirements)
    - [Setup](#setup)
      - [Salt Edge](#salt-edge)
      - [Doppler](#doppler)
      - [HTTPS](#https)
    - [Usage](#usage)
      - [Cheatsheet](#cheatsheet)
  - [Run](#run)
    - [Docker](#docker)
      - [Run local PostgresSQL DB](#run-local-postgressql-db)
      - [Remove container storage](#remove-container-storage)
    - [Source](#source)
      - [Run without build](#run-without-build)
      - [Build first, then run](#build-first-then-run)
  - [Test](#test)
  - [TODO](#todo)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# HQ Service modulith

## Local development
### Requirements

- Go 1.19
- [Docker](https://docs.docker.com/install/) 19.03+
- [Docker Compose](https://docs.docker.com/compose/install/) 1.25+

### Setup

#### Salt Edge

Generate PKI to sign [Signature header](https://docs.saltedge.com/general/#signature_headers)

```bash
# For macOS or Linux
openssl genrsa -out private.pem 2048
openssl rsa -pubout -in private.pem -out public.pem

# For Windows
cd C:\OpenSSL-Win32\bin
openssl genrsa -out private.pem 2048
openssl rsa -pubout -in private.pem -out public.pem
```
#### Doppler

Setup Doppler:

```bash
task scripts:install:doppler
```

```
Naming convention of environment vars required to run and test project:

<PROJECT>_<VAR>         - global vars, not specific for some embedded microservice (e.g. domain)
<PROJECT>_X_<SVC>_<VAR> - vars related to external services (e.g. databases)
<PROJECT>_<MS>_<VAR>    - vars related to embedded microservice (e.g. addr)
<PROJECT>__<MS>_<VAR>   - private vars for embedded microservice
```

#### HTTPS

1. This project requires https:// and will send HSTS and CSP HTTP headers,
  [Create local CA to issue localhost HTTPS 
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
- `./scripts/cover` - analyse and show coverage
- `./scripts/build` - build docker image and binaries in `bin/`
  - Access project at host/port(s) defined in `doppler`.

#### Cheatsheet

```sh
doppler run -- docker-compose up -d --remove-orphans               # (re)start all project's services
docker-compose logs -f -t                           # view logs of all services
docker-compose logs -f SERVICENAME                  # view logs of some service
docker-compose ps                                   # status of all services
docker-compose restart SERVICENAME
docker-compose exec SERVICENAME COMMAND             # run command in given container
docker-compose stop && docker-compose rm -f                     # stop the project
docker volume rm PROJECT_SERVICENAME    # remove some service's data
```
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

#### Run without build
```bash
# main.go is the entry point with the `main` function
task scripts:run
```

```bash
# Run with cmd arguments to override configurations
task scripts:run -- --port 17002 --wa.id example --wa.origins https://example.com,android:apk-key-hash:your_apk_hash  # Specific auth service running on port `17002` with webauthn ID equals `example`; webauthn origns equals `https://example.com,android:apk-key-hash:your_apk_hash`  
```

#### Build first, then run

We use the `Taskfile` command to build our binary, then, run our built `hq` binary.

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

## Test

Run all test
```sh
go test -v ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
go test -count=1 --tags=integration ./... # run integration tests
```

## TODO

Functionality Group 1: add/connect assets and debts

- [x] Plaid aggregator with dev env
- [ ] Plaid aggregator with stg, prd env
- [x] Implement [webauthn](https://github.com/go-webauthn/webauthn) API
- [ ] Implement [Lago](https://www.getlago.com/resources/compare/lago-vs-stripe) for billing service
- [ ] Implement authorization with `casbin`
- [x] Create `User` table
- [x] Create CRUD REST API for `User` model
- [ ] Create asset tables based on Kubera features
- [ ] Create CRUD REST API for asset models
- [x] Integration test for `auth` svc APIs

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
