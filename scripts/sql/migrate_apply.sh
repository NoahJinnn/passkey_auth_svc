#!/bin/bash
scriptsdir=$( dirname -- "$0"; )
cd $scriptsdir/../../

if ! command -v atlas &> /dev/null; then
    echo "Atlas Go is not installed or not in PATH, please check atlasgo.io for installation"
fi

atlas migrate apply \
  --dir "file://migrate/migrations/pgsql" \
  --url "postgres://passkey-auth-service:h3ll0@localhost:5432/passkey-auth-service?sslmode=require"

echo "Migrate PgSQL done"
