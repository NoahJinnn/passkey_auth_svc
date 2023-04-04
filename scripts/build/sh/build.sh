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
	rm -rf $scriptsdir/../../../debugbin/
	mkdir $scriptsdir/../../../debugbin/
	go build -gcflags="all=-N -l" -a -o $scriptsdir/../../../debugbin/ "$@" $scriptsdir/../../../cmd/*
}

if [ "$1" == "debug" ]; then
  echo "Build debug bin"
  build_debug
else
  echo "Build normal bin"
  build
fi
