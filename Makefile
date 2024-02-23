.PHONY: build clean test

# The binary to build (just the basename).
BIN := shizu

# Where to build the binary.
OUTDIR := ./bin

# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe --tags --always --dirty)

build:
	@echo "building ${VERSION}"
	@go build -o ${OUTDIR}/${BIN} -ldflags "-X main.Version=${VERSION}" ./cmd/shizu

clean:
	@echo "cleaning"
	@rm -rf ${OUTDIR}

test:
	@echo "testing"
	@go test ./...
