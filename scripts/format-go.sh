#!/bin/bash

set -euxo pipefail

cd "$(dirname "$0")/.."

go fmt github.com/neprune/todo...
