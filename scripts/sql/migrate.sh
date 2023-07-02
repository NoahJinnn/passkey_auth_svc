#!/bin/bash
scriptsdir=$( dirname -- "$0"; )
cd $scriptsdir/../../

if ! command -v atlas &> /dev/null; then
    echo "Atlas Go is not installed or not in PATH, please check atlasgo.io for installation"
fi

atlas migrate diff pgsql \
--dir "file://migrate/migrations/pgsql" \
--to "ent://ent/schema/pgsql" \
--dev-url "docker://postgres/15/test?search_path=public"

atlas migrate lint \
  --dev-url="docker://postgres/15/test?search_path=public" \
  --dir="file://migrate/migrations/pgsql" \
  --latest=1

atlas migrate diff sqlite \
--dir "file://migrate/migrations/sqlite" \
--to "ent://ent/schema/sqlite" \
--dev-url "sqlite://file?mode=memory&_fk=1"

atlas migrate lint \
  --dev-url="sqlite://file?mode=memory" \
  --dir="file://migrate/migrations/sqlite" \
  --latest=1