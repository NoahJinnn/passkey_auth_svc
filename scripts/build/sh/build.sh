#!/bin/bash
set -x -e -o pipefail
scriptsdir=$( dirname -- "$0"; )
export CGO_ENABLED=0


build() {
	rm -rf $scriptsdir/../../../bin/
	mkdir $scriptsdir/../../../bin/
	GOOS=$1 GOARCH=$2 go build -ldflags="-s -w" -a -o $scriptsdir/../../../bin/hq "$@" $scriptsdir/../../../main.go
}

build_debug() {
	rm -rf $scriptsdir/../../../bindebug/
	mkdir $scriptsdir/../../../bindebug/
	GOOS=$1 GOARCH=$2 go build -gcflags="all=-N -l" -a -o $scriptsdir/../../../bindebug/hq "$@" $scriptsdir/../../../main.go
}

if [ "$3" == "debug" ]; then
  echo "Build binary for debug"
  build_debug
else
  echo "Build binary for production"
  build
fi
