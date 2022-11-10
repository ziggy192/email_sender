PWD = $(shell pwd)

SRC = `go list -f {{.Dir}} ./... | grep -v /vendor/`

install:
	@echo "==> Installing tools..."
	@go install golang.org/x/tools/...
	@go install golang.org/x/lint/golint
	@GO111MODULE=off go install github.com/golang/mock/mockgen
	@GO111MODULE=off go install mvdan.cc/gofumpt/gofumports
	@GO111MODULE=off go install github.com/daixiang0/gci
	@brew install golangci/tap/golangci-lint
	@brew upgrade golangci/tap/golangci-lint

all: fmt lint test

fmt:
	@echo "==> Formatting source code..."
	@go fmt $(SRC)

lint:
	@echo "==> Running lint check..."
	@go vet $(SRC)

test:
	@echo "==> Running tests..."
	@go clean -testcache ./...
	@go test `go list ./... | grep -v cmd` -race --cover

.PHONY: all fmt lint test
