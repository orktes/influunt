build: build-shared-lib
.PHONY: build

build-shared-lib:
	./scripts/build_shared_lib.sh
.PHONY: build-shared-lib

test: test-go test-e2e
.PHONY: test

test-go:
	go test ./...
.PHONY: test-go

test-e2e:
	./tests/run_tests.sh
.PHONY: test-e2e

clean:
	rm -Rf build
.PHONY: clean 