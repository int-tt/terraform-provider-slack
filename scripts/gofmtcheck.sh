#!/usr/bin/env bash

gofmt_files=$(find . -name "*.go" | grep -v vendor | xargs gofmt -l -s -d)
if [[ -n ${gofmt_files} ]]; then
    echo "${gofmt_files}"
    exit 1
fi

exit 0