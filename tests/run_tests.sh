#!/bin/bash

set -e

cat ./scripts/Dockerfile
docker build -t influunt:latest - < ./scripts/Dockerfile
docker run --rm -it -v "$GOPATH":/go -w /go/src/github.com/orktes/influunt influunt:latest sh -c '
set -e

cd python
python3 setup.py install
cd ..

ls -l /usr/local/lib/python3.6/site-packages/influunt-0.1-py3.6.egg/influunt/

cd tests

for filename in *.py; do
echo "Running $filename"
python3 $filename
done

'