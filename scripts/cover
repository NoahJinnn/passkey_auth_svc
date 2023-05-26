#!/bin/bash
set -e -o pipefail
export PATH="$PWD/.gobincache:$PATH"
GOBIN="$PWD/.gobincache" go install gotest.tools/gotestsum

gotestsum -- \
	-coverprofile cover.out \
	-tags=integration "$@" ./...

go tool cover -func=cover.out | tail -n 1 | xargs echo

go tool cover -html=cover.out
