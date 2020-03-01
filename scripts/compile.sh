#!/usr/bin/env bash

set -e
REPO_ROOT=$(git rev-parse --show-toplevel)

VERSION=$(git describe --tags "$(git rev-list --tags --max-count=1)")"-local"
COMMIT_HASH=$(git rev-parse --short HEAD 2>/dev/null)
DATE=$(date "+%Y-%m-%d")
GOVERSION=$(go version | awk '{print $3;}')
GO_BUILD_LDFLAGS="-s -w -X 'main.commit=${COMMIT_HASH}' -X 'main.date=${DATE}' -X 'main.version=${VERSION}' -X 'main.goversion=${GOVERSION}'"

APP="aoc-cli"
MODULE="github.com/oleg-balunenko/advent-of-code"

GO_BUILD_PACKAGE="${MODULE}/cmd/aoc-cli/."

BIN_DIR="${REPO_ROOT}/bin"
rm -rf "${BIN_DIR}"
mkdir -p "${BIN_DIR}"

go build -o "${BIN_DIR}"/"${APP}" -a -ldflags "${GO_BUILD_LDFLAGS}" "${GO_BUILD_PACKAGE}"

echo "Executable built and stored at ${BIN_DIR}/${APP}"
