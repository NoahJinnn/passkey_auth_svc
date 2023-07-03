#!/bin/bash
scriptsdir=$( dirname -- "$0"; )
cd $scriptsdir/../../

if ! command -v atlas &> /dev/null; then
    echo "Atlas Go is not installed or not in PATH, please check atlasgo.io for installation"
fi

atlas migrate apply \
  --dir "file://migrate/migrations/pgsql" \
  --url "postgres://hqservice:h3ll0HQ@localhost:5432/hqservice?sslmode=require"

echo "Migrate PgSQL done"

atlas migrate apply \
  --dir "file://migrate/migrations/sqlite" \
  --url "sqlite://file.db?_fk=1"

echo "Migrate SQLite done"