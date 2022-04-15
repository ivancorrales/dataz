GO_VERSION := $(shell cat .go-version)
GO  = GOFLAGS=-mod=readonly go


all: clean deps install fmt lint test

.PHONY: setup
setup:
	git add .pre-commit-config.yaml
	pre-commit install

.PHONY: deps
deps:
	$(GO) mod tidy
	cd tools && go mod tidy

.PHONY: fmt
fmt:
	$(GO) fmt ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	$(GO) test -race -covermode=atomic -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html

.PHONY: go-clean
go-clean:
	$(GO) clean -r -i -cache -testcache -modcache

.PHONY: clean
clean:
	rm -rf dist
	rm -f coverage.*

.PHONY: install
install:
	cd tools && $(GO) install $(shell cd tools && go list -f '{{ join .Imports " " }}' -tags=tools)
