#!/bin/bash
set -x -e -o pipefail
scriptsdir=$( dirname -- "$0"; )
GOOS=$1
GOARCH=$2
# Check if GOOS is empty



if [ -z "$GOOS" ]; then
  echo "GOOS is empty. Set GOOS to linux"
  GOOS=linux
fi

if [ -z "$GOARCH" ]; then
  echo "GOARCH is empty. Set GOARCH to amd64"
  GOARCH=amd64
fi


gitver() {
	local ver branch rev dirty
	ver="$(git tag -l --sort=-version:refname --merged HEAD 'v*' | head -n 1)"
	branch="$(git rev-parse --abbrev-ref HEAD)"
	rev="$(git log -1 --format='%h')"

	echo "${ver:+$ver }${branch:+$branch } $(date -u +"%F_%T")"
}


build() {
	rm -rf $scriptsdir/../../../bin/
	mkdir $scriptsdir/../../../bin/
	GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-X '$(go list -m)/pkg/def.ver=$(gitver)' -s -w" -a -o $scriptsdir/../../../bin/hq "$@" $scriptsdir/../../../main.go
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
