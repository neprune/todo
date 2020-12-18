#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")/.."

go run cmd/main.go "$@"
