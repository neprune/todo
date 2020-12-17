#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")/.."

cd test
go run ../cmd/main.go "$@"
