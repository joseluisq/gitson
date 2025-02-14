BUILD_TIME ?= $(shell date -u '+%Y-%m-%dT%H:%m:%S')

install:
	@go version
	@go get -v golang.org/x/lint/golint
.PHONY: install

run:
	@go version
	@go run examples/main.go
.PHONY: run

test:
	@go version
	@golint -set_exit_status ./...
	@go vet ./...
	@go test $$(go list ./... | grep -v /examples) \
		-v -timeout 30s -race -coverprofile=coverage.txt -covermode=atomic
.PHONY: test

build:
	@go version
	@go build -v \
		-ldflags "-s -w -X 'main.version=0.0.0' -X 'main.buildTime=$(BUILD_TIME)'" \
		-a -o bin/gitson ./examples
.PHONY: build

coverage:
	@bash -c "bash <(curl -s https://codecov.io/bash)"
.PHONY: coverage
