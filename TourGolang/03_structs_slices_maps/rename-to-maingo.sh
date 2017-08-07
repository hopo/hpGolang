#!/bin/bash

find . -name "*.go" -exec sh -c 'mv -i "$1" "$(dirname $1)/main.go"' _ {} \;
