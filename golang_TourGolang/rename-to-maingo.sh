#!/bin/bash

find . -name "*.go" -exec sh -c 'mv "$1" "$(dirname $1)/main.go"' _ {} \;
