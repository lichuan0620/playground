ROOT := github.com/lichuan0620/playground
VERSION ?= $(shell git describe --tags --always --dirty)
COMMIT := $(shell git rev-parse HEAD)
BRANCH := $(shell git branch | grep \* | cut -d ' ' -f2)
BUILD_USER ?= ChuanLi
BUILD_DATE := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

REGISTRY ?= lichuan0620

GO_VERSION ?= 1.15

GOPATH ?= $(shell go env GOPATH)
BIN_DIR := $(GOPATH)/bin
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint

.PHONY: test lint build build-local build-linux container push clean

build: build-local

test:
	@echo ">> running tests"
	@go test $$(go list ./... | grep -v /vendor) -coverprofile=coverage.out
	@go tool cover -func coverage.out | tail -n 1 | awk '{ print "Total coverage: " $$3 }'

lint: $(GOLANGCI_LINT)
	@echo ">> running golangci-lint"
	@$(GOLANGCI_LINT) run

$(GOLANGCI_LINT):
	@curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(BIN_DIR) v1.31.0

build-local: clean
	@echo ">> building local binary"
	@GOOS=$(shell uname -s | tr A-Z a-z) GOARCH=amd64 CGO_ENABLED=0     \
	  go build -mod=vendor -ldflags "-s -w                              \
	    -X $(ROOT)/pkg/version.Version=$(VERSION)                       \
	    -X $(ROOT)/pkg/version.Branch=$(BRANCH)                         \
	    -X $(ROOT)/pkg/version.Commit=$(COMMIT)                         \
	    -X $(ROOT)/pkg/version.BuildUser=$(BUILD_USER)                  \
	    -X $(ROOT)/pkg/version.BuildDate=$(BUILD_DATE)"                 \
	    -o playground

build-linux:
	@echo ">> building linux binary"
	@PWD=$(PWD) sudo docker run --rm -t                                 \
	  -v "$(PWD):/go/src/$(ROOT)" -w /go/src/$(ROOT)                    \
	  -e GOOS=linux	-e GOARCH=amd64 -e GOPATH=/go                       \
	  golang:$(GO_VERSION) /bin/bash -c                                 \
		'go build -mod=vendor -ldflags "-s -w                           \
		  -X $(ROOT)/pkg/version.Version=$(VERSION)                     \
		  -X $(ROOT)/pkg/version.Branch=$(BRANCH)                       \
		  -X $(ROOT)/pkg/version.Commit=$(COMMIT)                       \
		  -X $(ROOT)/pkg/version.BuildUser=$(BUILD_USER)                \
		  -X $(ROOT)/pkg/version.BuildDate=$(BUILD_DATE)"               \
		-o playground'

container: build-linux
	@echo ">> building image"
	@sudo docker build -t $(REGISTRY)/playground:$(VERSION) -f ./Dockerfile .

push: container
	@echo ">> pushing image"
	@sudo docker push $(REGISTRY)/playground:$(VERSION)

clean:
	@echo ">> cleaning up"
	@rm -f playground coverage.out
