.PHONY: all
all: build lint test

.PHONY: build
build: ; go build ./...

.PHONY: test
test: ; go test ./...

.PHONT: lint
lint: ; golangci-lint run
