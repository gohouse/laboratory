#!/bin/bash

rm -f clearTask.bin

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o clearTask.bin

upx clearTask.bin

basepath=$(cd `dirname $0`; pwd)

expectScpFile=expect
expectSSHFile=expectSsh

ip="180.97.188.201"
port="4321"
password="Admin132"
file="clearTask.bin"
path="/www/sites/go"

expect $basepath/"$expectScpFile" "$ip" "$port" "$password" "$file" "$path"
# expect $basepath/"$expectSSHFile" "$ip" "$port" "$password" "$file"

# sh $basepath/reloadserver.sh