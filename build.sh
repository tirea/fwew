#!/bin/bash
echo "formatting..."
for file in *.go util/*.go; do gofmt -w $file; done && \
echo "compiling..."
go build fwew.go && mv fwew bin/ && \
echo "done."

