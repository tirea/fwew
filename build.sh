#!/bin/bash
echo "formatting..."
for file in *.go util/*.go; do gofmt -w $file; done && \
echo "compiling..."
go build -ldflags "-X github.com/tirea/fwew/util.Build=`git rev-parse HEAD`" -o fwew fwew.go && \
mv fwew bin/ && \
echo "done."
