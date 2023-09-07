#!/bin/bash
set -x -e -o pipefail
source $( dirname -- "$0"; )/build.sh

# Build binaries for linux-based Docker container.
GOOS=$1 GOARCH=$2 build "$@"

cd docker

# Check if Docker is installed
elif command -v docker > /dev/null; then
  echo "Docker is installed, building using Docker"
  docker build -t "$(basename "$(go list -m)")" $( dirname -- "$0"; )/../../../

# If neither Podman nor Docker are installed
else
  echo "Please install Docker, unable to build image."
  exit 1
fi


