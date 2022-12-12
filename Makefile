.PHONY: all
all: build lint test

.PHONY: build
build: ; go build ./...

.PHONY: test
test: ; go test -cover ./...

.PHONY: lint
lint: ; golangci-lint run
