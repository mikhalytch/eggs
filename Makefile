.PHONY: all
all: build lint test

.PHONY: build
build: ; go build ./...

.PHONY: test
test: ; go test -race -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: lint
lint: ; golangci-lint run
