#!/bin/bash

rm -f wccgen.bin

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o wccgen.bin

upx wccgen.bin

cp wccgen.bin ~/Downloads/
