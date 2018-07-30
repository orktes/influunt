#!/bin/bash

#set -v
set -e

cat ./scripts/Dockerfile
docker build -t influunt:latest - < ./scripts/Dockerfile
docker run --rm -it -v "$GOPATH":/go -w /go/src/github.com/orktes/influunt influunt:latest sh -c '
cd go/python/
for GOOS in darwin linux windows; do
  for GOARCH in 386 amd64; do
    mkdir -p build/sharedlib/$GOOS/$GOARCH
    echo Building $GOOS for $GOARCH
    go build -buildmode=c-shared -o /go/src/github.com/orktes/influunt/build/sharedlib/$GOOS/$GOARCH/influunt_core.so
  done
done
'