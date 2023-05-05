#!/bin/bash
set -x -e -o pipefail
scriptsdir=$( dirname -- "$0"; )
export CGO_ENABLED=0

build() {
	rm -rf $scriptsdir/../../../bin/
	mkdir $scriptsdir/../../../bin/
	go build -a -o $scriptsdir/../../../bin/ "$@" $scriptsdir/../../../cmd/*
}

build_debug() {
	rm -rf $scriptsdir/../../../bindebug/
	mkdir $scriptsdir/../../../bindebug/
	go build -gcflags="all=-N -l" -a -o $scriptsdir/../../../bindebug/ "$@" $scriptsdir/../../../cmd/*
}

if [ "$1" == "debug" ]; then
  echo "Build binary for debug"
  build_debug
else
  echo "Build binary for production"
  build
fi
