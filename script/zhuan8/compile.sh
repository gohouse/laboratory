#!/bin/bash

rm -f releaseTimeoutTask.bin

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o releaseTimeoutTask.bin

upx releaseTimeoutTask.bin

basepath=$(cd `dirname $0`; pwd)

expectScpFile=expect
expectSSHFile=expectSsh

ip="180.97.188.201"
port="4321"
password="Admin132"
file="releaseTimeoutTask.bin"
path="/www/sites/go"

expect $basepath/"$expectScpFile" "$ip" "$port" "$password" "$file" "$path"
# expect $basepath/"$expectSSHFile" "$ip" "$port" "$password" "$file"

# sh $basepath/reloadserver.sh