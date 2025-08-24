#!/bin/bash
cd "$(dirname "$0")"
cd ../../

echo "Doing work in directory $PWD"
source scripts/development/meta.sh

BASE_DIR="$PWD"
BRANCH="$(git rev-parse --abbrev-ref HEAD)"
LATEST_COMMIT="$(git rev-parse --short $BRANCH)"

docker build . --file docker/Dockerfile --no-cache --build-arg TARGETOS=linux --build-arg TARGETARCH=amd64 --tag $BINARY:$LATEST_COMMIT
docker tag $BINARY:$LATEST_COMMIT $BINARY:latest