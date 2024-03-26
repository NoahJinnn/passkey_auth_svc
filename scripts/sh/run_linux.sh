#!/bin/sh
scriptsdir=$(dirname -- "$0")
cd $scriptsdir/../../
PROFILE=$1

if
    ! command -v doppler &
    >/dev/null
then
    echo "Doppler is not installed or not in PATH, please run 'task scripts:install:doppler' and try again"

elif
    ! command -v air &
    >/dev/null
then
    echo "Air is not installed or not in GOPATH, please run 'task scripts:install:gotools' and try again. Make sure your shell can lookup for GOBIN path by 'export PATH=$PATH:$(go env GOPATH)/bin'"
else
    echo "Passkey Auth service started by exec built app"
    doppler configure set config=$PROFILE project=passkey-auth-service # doppler profile to run passkey-auth-service
    AUTH_DOMAIN=$(doppler secrets get AUTH_DOMAIN --plain)  # get AUTH_DOMAIN value from doppler

    doppler run -- docker compose -f docker/docker-compose.pgsql.yml down &&
        doppler run -- docker compose -f docker/docker-compose.pgsql.yml up -d --remove-orphans &&
        doppler run -- air serve "--wa.id $AUTH_DOMAIN --wa.origins "https://${AUTH_DOMAIN}""
fi
