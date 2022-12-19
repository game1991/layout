#!/bin/sh

set -o errexit
set -o nounset
#set -o pipefail

# if [ -z "${OS:-}" ]; then
#     echo "OS must be set"
#     exit 1
# fi
# if [ -z "${ARCH:-}" ]; then
#     echo "ARCH must be set"
#     exit 1
# fi
if [ -z "${VERSION:-}" ]; then
    echo "VERSION must be set"
    exit 1
fi
if [ -z "${BIN:-}" ]; then
    echo "BIN must be set"
    exit 1
fi

export CGO_ENABLED=0
export GOARCH="${ARCH}"
export GOOS="${OS}"
export GO111MODULE=on
export GOFLAGS="${GOFLAGS:-} -mod=${MOD}"

go build -o ${BIN}                                        \
    -installsuffix "static"                               \
    -ldflags "-X $(go list -m)/build.Version=${VERSION} -X $(go list -m)/build.Branch=${BRANCH} -X $(go list -m)/build.Tag=${TAG} -X  $(go list -m)/build.LastTime=${LASTTIME} -X  $(go list -m)/build.LastCommit=${LASTCOMMIT} -X  $(go list -m)/build.GitRepoPATH=${GitRepoPATH}"   \
    ./cmd/main.go
