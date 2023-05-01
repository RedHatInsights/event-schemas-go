#!/usr/bin/env sh

find apps core -type f -name "*.go" -print0 | xargs -0L1 go build
