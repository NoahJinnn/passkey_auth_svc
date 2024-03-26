**Table of Contents**

- [Passkey Auth Service modulith](#passkey-auth-service-modulith)
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


# Passkey Auth Service modulith
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

<PROJECT>_<VAR>         - shared vars, not specific for any embedded microservice (e.g. jwt token)
<PROJECT>_X_<VAR> - vars related to external services (e.g. databases)
<PROJECT>_<MS>_<VAR>    - vars related to embedded microservice (e.g. addr)
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
         hq: inf      main: `started` version f/design-task-command-to-run-passkey-auth-service 51adc59-dirty 2023-02-15_09:36:06
```

## Test

Running tests

```sh
# run all tests
$ PROFILE=<your-doppler-profile> task scripts:test

# TODO: 
# currently not working, fix this and add integration test
# change this to task scripts:test:integration
$ go test -count=1 --tags=integration ./... # run integration tests
```
## Run app with docker-compose 

1. Build container
```sh
$ task scripts:build:container
```

2. Run docker-compose
```sh
$ DOPPLER_TOKEN=$(doppler configs tokens create docker -p passkey-auth-service -c dev_noah --max-age 30m --plain) \
IMAGE_TAG=passkey-auth-service:latest \
doppler run -- docker  compose -f docker/docker-compose.svc.yml up
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
