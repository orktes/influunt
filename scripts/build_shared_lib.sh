#!/bin/bash

set -e

if [[ "$OSTYPE" == "darwin"* ]]; then
  if hash go 2>/dev/null; then
    echo "Building darwin amd64"
    pushd go/python/
    go build -buildmode=c-shared -o ../../build/sharedlib/darwin/amd64/influunt_core.so
    popd
  else
    echo "Go not installed = no darwin lib build"
  fi
else
  echo "Not running OSX = no darwin lib build"
fi

docker build -t influunt:latest - < ./scripts/Dockerfile
docker run --rm -it -v "$GOPATH":/go -w /go/src/github.com/orktes/influunt influunt:latest sh -c '
cd go/python/
echo "Building linux amd64"
go build -buildmode=c-shared -o /go/src/github.com/orktes/influunt/build/sharedlib/linux/amd64/influunt_core.so
'