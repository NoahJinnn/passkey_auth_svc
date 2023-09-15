#!/bin/bash
set -x -e -o pipefail
source $( dirname -- "$0"; )/build.sh
ROOT=$scriptsdir/../../../

IMAGE_TAG=$(basename "$(go list -m)")
cd $ROOT
# Check if Docker is installed
if command -v docker > /dev/null; then
  echo "Docker is installed, building using Docker"
  doppler run -- docker build -t "$IMAGE_TAG" .

# If neither Podman nor Docker are installed
else
  echo "Please install Docker, unable to build image."
  exit 1
fi
