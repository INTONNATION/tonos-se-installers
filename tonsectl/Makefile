GO_DIR ?= $(shell pwd)
GO_PKG ?= $(shell go list -e -f "{{ .ImportPath }}")

GOOS ?= $(shell go env GOOS || echo linux)
GOARCH ?= $(shell go env GOARCH || echo amd64)
CGO_ENABLED ?= 0



init: ## init packages
	mkdir -p artifacts ;\
    rm -rf artifacts/*

build: init ## build binary file
	GOOS=${GOOS} CGO_ENABLED=${CGO_ENABLED} GOARCH=${GOARCH} \
	go build -ldflags "-X $(GO_PKG)/cmd/version.appVersion=$(TAG)-$$(date -u +%Y%m%d%H%M) -X $(GO_PKG)/cmd.nodejs_version=14.16.0 -X $(GO_PKG)/cmd.tonosse_version=0.27.1 -X $(GO_PKG)/cmd.arango_version=3.7.9" -o "$(GO_DIR)/artifacts/tonsectl_${GOOS}" main.go

.DEFAULT_GOAL := help
