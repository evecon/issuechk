#!/usr/bin/env bash
set -e
set -x

VERSION=$(git describe --dirty --all)
godep go build -ldflags="-X main.Version=$VERSION" ./cmd/issuechk
