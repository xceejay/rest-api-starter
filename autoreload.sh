#!/bin/bash

find . -name '*.go' | grep -E -v "test.go$" | entr -r bash -c 'clear && echo "rebuilding..." && make run '
