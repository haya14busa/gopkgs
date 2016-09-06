#!/bin/bash
ls $GOPATH/src/golang.org/x/tools/imports/*.go | grep --invert-match -e '_test.go$' | grep -e 'fastwalk*' -e 'fix.go' -e 'zstdlib.go' | xargs -i cp {} ./
cp $GOPATH/src/golang.org/x/tools/LICENSE ./
