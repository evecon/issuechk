#!/usr/bin/env bash
set -e
set -x

if [[ $1 == "xc" ]] ; then
    BUILDTOOL="gox"
else
    BUILDTOOL="godep go build"
fi

VERSION=$(git describe --dirty --all)
$BUILDTOOL -ldflags="-X main.Version=$VERSION" ./cmd/issuechk
