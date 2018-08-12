build: build-shared-lib
.PHONY: build

build-shared-lib:
	./scripts/build_shared_lib.sh
.PHONY: build-shared-lib

test: test-go test-e2e
.PHONY: test

install:
	(cd python && python3 setup.py install)
.PHONY: install

go-generate:
	cd go
	go generate ./...
.PHONY: go-generate

test-go:
	./scripts/test.sh
.PHONY: test-go

test-e2e:
	./tests/run_tests.sh
.PHONY: test-e2e

clean:
	rm -Rf build
	rm -Rf python/build
	rm -Rf python/dist
	rm -Rf python/influunt.egg-info
.PHONY: clean 