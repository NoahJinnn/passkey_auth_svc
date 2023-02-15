#!/bin/bash
set -x -e -o pipefail

# Needed until native support will be implemented: https://github.com/golang/go/issues/37475
### [version] [branch] revision[-dirty] build_date_time
gitver() {
	local ver branch rev dirty
	ver="$(git tag -l --sort=-version:refname --merged HEAD 'v*' | head -n 1)"
	branch="$(git rev-parse --abbrev-ref HEAD)"
	rev="$(git log -1 --format='%h')"
	git update-index -q --refresh --unmerged >/dev/null
	git diff-index --quiet HEAD || dirty="-dirty"
	test "$branch" = "HEAD" || test "$branch" = "main" && branch=

	echo "${ver:+$ver }${branch:+$branch }$rev$dirty $(date -u +"%F_%T")"
}

scriptsdir=$( dirname -- "$0"; )

export CGO_ENABLED=0

build() {
	rm -rf $scriptsdir/../../../bin/
	mkdir $scriptsdir/../../../bin/
	go build -ldflags "-X '$(go list -m)/pkg/def.ver=$(gitver)'" -o $scriptsdir/../../../bin/ "$@" $scriptsdir/../../../cmd/*
}

build_debug() {
	rm -rf $scriptsdir/../../../debugbin/
	mkdir $scriptsdir/../../../debugbin/
	go build -ldflags "-X '$(go list -m)/pkg/def.ver=$(gitver)'" -gcflags="all=-N -l" -o $scriptsdir/../../../debugbin/ "$@" $scriptsdir/../../../cmd/*
}

if [ "$1" == "debug" ]; then
  echo "Build debug bin"
  build_debug
else
  echo "Build product bin"
  build
fi
