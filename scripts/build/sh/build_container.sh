#!/bin/bash
set -x -e -o pipefail

source $( dirname -- "$0"; )/build.sh

# Build binaries for linux-based Docker container.
GOOS=linux GOARCH=amd64 build "$@"

cd docker

# Check if Podman is installed
if command -v podman > /dev/null; then
  echo "Podman is installed, building using Podman"
  podman build -t "$(basename "$(go list -m)")" $( dirname -- "$0"; )/../../../

# Check if Docker is installed
elif command -v docker > /dev/null; then
  echo "Docker is installed, building using Docker"
  docker build -t "$(basename "$(go list -m)")" $( dirname -- "$0"; )/../../../

# If neither Podman nor Docker are installed
else
  echo "Neither Podman nor Docker are installed, unable to build image."
  exit 1
fi

# # Build binaries for host system.
test "$(go env GOOS)" = linux -a "$(go env GOARCH)" = amd64 || build "$@"
