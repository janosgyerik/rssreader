#!/bin/sh

set -x
cd "$(dirname "$0")"
mkdir -p target
go test -coverprofile target/cover.out
go tool cover -html=target/cover.out -o target/cover.html
