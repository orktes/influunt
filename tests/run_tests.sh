#!/bin/bash

set -e

cat ./scripts/Dockerfile
docker build -t influunt:latest - < ./scripts/Dockerfile
docker run --rm -it -v "$GOPATH":/go -w /go/src/github.com/orktes/influunt influunt:latest sh -c '
cd tests
python3 simple.py
'