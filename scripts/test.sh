#!/bin/bash

#set -v
set -e

docker build -t influunt:latest - < ./scripts/Dockerfile
docker run --rm -it -v "$GOPATH":/go -w /go/src/github.com/orktes/influunt influunt:latest sh -c '
cd go/
go test ./...
'