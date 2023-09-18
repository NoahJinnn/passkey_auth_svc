#!/bin/bash
set -e -o pipefail
export PATH="$PWD/.gobincache:$PATH"
go generate
set -x

scriptsdir=$( dirname -- "$0"; )
cd $scriptsdir/../../

PROFILE=$1

set +x # disable debug output
# Check if DOPPLER_TOKEN is set, if not, create a new one and export it
# This is to avoid creating a new token on every run

# if [ -z "$DOPPLER_TOKEN" ]; then
#     echo "DOPPLER_TOKEN is empty."
#     export DOPPLER_TOKEN=$(doppler configs tokens create your-token-name-here -p hqservice -c dev --max-age 1m --plain)
# fi
set -x
doppler configure set config=$PROFILE project=hqservice
doppler run -- go test -race -covermode atomic -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
