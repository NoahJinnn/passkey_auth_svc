**Table of Contents**

- [HQ Service modulith](#hq-service-modulith)
  - [Requirements](#requirements)
    - [Salt Edge](#salt-edge)
    - [Doppler](#doppler)
    - [HTTPS](#https)
  - [Run](#run)
    - [Docker](#docker)
    - [Taskfile](#taskfile)      
  - [Test](#test)
  - [Migrate](#migrate)
  - [TODO](#todo)


# HQ Service modulith
## Requirements

- Go 1.19
- [Docker](https://docs.docker.com/install/) 19.03+
- [Docker Compose](https://docs.docker.com/compose/install/) 1.25+

### Salt Edge

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
### Doppler

Setup Doppler:

```bash
task scripts:install:doppler
```

### Taskfile
We use the `Taskfile` command to build our binary, then, run our built `hq` binary.
- `./scripts/sh` - test and run project
- `./scripts/cover` - analyse and show coverage
- `./scripts/build` - build docker image and binaries in `bin/`
  - Access project at host/port(s) defined in `doppler`.


```md
Naming convention of environment vars required to run and test project:

<PROJECT>_<VAR>         - global vars, not specific for some embedded microservice (e.g. domain)
<PROJECT>_X_<SVC>_<VAR> - vars related to external services (e.g. databases)
<PROJECT>_<MS>_<VAR>    - vars related to embedded microservice (e.g. addr)
<PROJECT>__<MS>_<VAR>   - private vars for embedded microservice
```

### HTTPS

1. Generate HTTPS certificates for PostgreSQL using [easyrsa](https://github.com/OpenVPN/easy-rsa/blob/master/README.quickstart.md)

```
$ /path/to/easyrsa init-pki
$ echo Dev CA $(go list -m) | /path/to/easyrsa build-ca nopass
$ /path/to/easyrsa --days=3650 "--subject-alt-name=DNS:postgres" build-server-full postgres nopass
$ /path/to/easyrsa --days=3650 "--subject-alt-name=DNS:localhost" build-server-full ms-hq nopass
$ /path/to/easyrsa --days=3650 "--subject-alt-name=IP:127.0.0.1" build-server-full ms-hq-int nopass
```

2. Generate HTTPS certificates for `networth` service using [mkcert](https://github.com/FiloSottile/mkcert#mkcert)
```
cd ./configs && mkdir http-pki && cd http-pki
mkcert localhost 127.0.0.1
```



## Development

Run our app with single command:

```bash
$ PROFILE=<your-doppler-profile> task scripts:run
```

or, more granular scripts as below.

```bash
# Run with cmd arguments to override configurations
# Specific auth service running on:
# port = `17002`
# webauthn ID = `example`;
# webauthn origins = `https://example.com,android:apk-key-hash:your_apk_hash`
$ doppler run -- air serve  "--port 17002 --wa.id example --wa.origins https://example.com,android:apk-key-hash:your_apk_hash"
```

Build binary only

```bash
# our binary gets installed into the ./bin/ folder, as `hq`.
$ task scripts:build:binary
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

Running tests

```sh
# run all tests
task scripts:test

# TODO: 
# currently not working, fix this and add integration test
# change this to task scripts:test:integration
go test -count=1 --tags=integration ./... # run integration tests
```

## Migrate

Running schema diff to generate migration scripts for:

```sh
# PostgreSQL
atlas migrate diff migration_name \
  --dir "file://ent/migrate/migrations" \
  --to "ent://ent/schema" \
  --dev-url "docker://postgres/15/test?search_path=public"
```

```sh
task scripts:migrate
```
### Docker

```sh
# Cheatsheet
doppler run -- docker-compose up -d --remove-orphans               # (re)start all project's services
docker-compose logs -f -t                                          # view logs of all services
docker-compose stop && docker-compose rm -f                        # stop & remove the containers
docker volume rm hqservice_postgres                                # clear all data
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
