.DEFAULT_GOAL := build

COMMIT_HASH = `git rev-parse --short HEAD 2>/dev/null`
BUILD_DATE = `date +%FT%T%z`

GO = go
BINARY_DIR=bin
GOFILES_NOVENDOR = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

BUILD_DEPS:= github.com/alecthomas/gometalinter

.PHONY: vendor test build

help:
	@echo "build      - go build"
	@echo "test       - go test"
	@echo "checkstyle - gofmt+golint+misspell"

vendor: ## Install govendor and sync vendored dependencies
	go get -u github.com/kardianos/govendor
	govendor sync

get-build-deps: vendor
	$(GO) get $(BUILD_DEPS)
	gometalinter --install

test: vendor
	govendor test +local

checkstyle: get-build-deps
	gometalinter --vendor ./... --fast --disable=gas --disable=errcheck --disable=gotype --deadline 10m

fmt:
	govendor fmt +local

# Builds the project
build: checkstyle test
	govendor build +local

install: checkstyle test
	govendor install +local

clean:
	if [ -d ${BINARY_DIR} ] ; then rm -r ${BINARY_DIR} ; fi
