#!/bin/bash
scriptsdir=$( dirname -- "$0"; )
cd $scriptsdir/../../

if ! command -v atlas &> /dev/null; then
    echo "entgo cli is not installed or not in PATH, please run 'go get -d entgo.io/ent/cmd/ent'"
fi

ent generate ./ent/schema --target ./ent
ent generate ./internal/db/sqlite/ent/schema
