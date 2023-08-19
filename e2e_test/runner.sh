#!/usr/bin/env bash

set -xe

trap './e2e_test/cleaner.sh' EXIT

# Build the CLI

go build -o dmgen ./dmgen
go install ./dmgen

# Run the tests

go test ./e2e_test/...
