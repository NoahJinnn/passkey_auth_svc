#!/bin/bash
set -e -o pipefail
export PATH="$PWD/.gobincache:$PATH"
go generate
set -x

scriptsdir=$( dirname -- "$0"; )
cd $scriptsdir/../../

PROFILE=$1

doppler configure set config=$PROFILE project=passkey-auth-service
doppler run -- go test -race -covermode atomic -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
