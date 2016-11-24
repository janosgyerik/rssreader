#!/usr/bin/env bash

cd "$(dirname "$0")"

set -x

go run cmd/rssreader/main.go "$@"
