#!/bin/bash
cd "$(dirname "$0")"
cd ../../

echo "Doing work in directory $PWD"
source scripts/development/meta.sh

BASE_DIR="$PWD"
cd "$BASE_DIR" || exit 1

echo "***********************************************"
echo "$(pwd)"
echo "***********************************************"

CGO_ENABLED=0 go build -ldflags '-s -w' || exit 1

mkdir $BASE_DIR/${BINARY}-linux-amd64 || echo "$BASE_DIR/${BINARY}-amd64 already created"
cp -f ${BINARY} $BASE_DIR/${BINARY}-linux-amd64/${BINARY}

cd "$BASE_DIR" || exit 1