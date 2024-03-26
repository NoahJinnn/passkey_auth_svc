#!/bin/bash
set -x -e -o pipefail
scriptsdir=$( dirname -- "$0"; )
GOOS=$1
GOARCH=$2

if [ -z "$GOOS" ]; then
  echo "GOOS is empty. Set GOOS to linux"
  GOOS=linux
fi

if [ -z "$GOARCH" ]; then
  echo "GOARCH is empty. Set GOARCH to amd64"
  GOARCH=amd64
fi

build() {
	rm -rf $scriptsdir/../../../bin/
	mkdir $scriptsdir/../../../bin/
	GOOS=$GOOS GOARCH=$GOARCH go build -a -o $scriptsdir/../../../bin/hq "$@" $scriptsdir/../../../main.go
}

build_debug() {
	rm -rf $scriptsdir/../../../bindebug/
	mkdir $scriptsdir/../../../bindebug/
	GOOS=$GOOS GOARCH=$GOARCH go build -gcflags="all=-N -l" -a -o $scriptsdir/../../../bindebug/hq "$@" $scriptsdir/../../../main.go
}

if [ "$3" == "debug" ]; then
  echo "Build binary for debug"
  build_debug
else
  echo "Build binary for production"
  build
fi
