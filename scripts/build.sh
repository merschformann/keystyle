#!/bin/bash

set -e

# Change to the project root directory
cd "$(dirname "$0")/.."

# Build the plugin
go build -buildmode=plugin plugin/keystyle.go

# Build golangci-lint
cd golangci-lint
make build
cd ..

# Copy binaries to dist directory
mkdir -p dist
mv -v golangci-lint/golangci-lint dist/
# mv -v keystyle.so dist/ # TODO: check gotz for a working (but bad) version
