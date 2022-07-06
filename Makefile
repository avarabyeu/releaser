.DEFAULT_GOAL := build

COMMIT_HASH = `git rev-parse --short HEAD 2>/dev/null`
BUILD_DATE = `date +%FT%T%z`

GO = go
BINARY_DIR=bin
GOFILES_NOVENDOR = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
GODIRS_NOVENDOR = $(shell go list ./... | grep -v /vendor/)

BUILD_DEPS:= github.com/alecthomas/gometalinter

.PHONY: vendor test build

help:
	@echo "build      - go build"
	@echo "test       - go test"
	@echo "checkstyle - gofmt+golint+misspell"


get-build-deps: vendor
	$(GO) get $(BUILD_DEPS)
	gometalinter --install

test:
	$(GO) test $(GODIRS_NOVENDOR)

checkstyle: get-build-deps
	gometalinter --vendor ./... --fast --disable=gas --disable=errcheck --disable=gotype --deadline 10m

fmt:
	gofmt -l -w -s ${GOFILES_NOVENDOR}

# Builds the project
build: checkstyle test
	$(GO) build ${BUILD_INFO_LDFLAGS} -o ${BINARY_DIR}/service-index ./

install: checkstyle test
	$(GO) install

clean:
	if [ -d ${BINARY_DIR} ] ; then rm -r ${BINARY_DIR} ; fi
