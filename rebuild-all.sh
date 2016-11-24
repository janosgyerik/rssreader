#!/usr/bin/env bash

cd "$(dirname "$0")"

v=$1
test "$v" || v=v0

arch=amd64
cli=cmd/rssreader/main.go

mkdir -p target

build() {
    GOARCH=$arch go build -o target/rssreader-$v-${GOOS}_$arch $cli
}

set -x

GOOS=darwin build
GOOS=linux build
GOOS=windows build
