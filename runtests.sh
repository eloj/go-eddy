#!/bin/bash
root=$(realpath $(dirname "${BASH_SOURCE[0]}"))
cd "$root"

go test ./... 2>&1 | tee .gotest.log
if [ $? -ne 0 ]; then
	echo "Test failures. Build aborted."
	exit 1
fi
