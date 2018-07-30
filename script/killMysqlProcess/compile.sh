#!/bin/bash

rm -f killMysqlProcess.bin

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o killMysqlProcess.bin

upx killMysqlProcess.bin

basepath=$(cd `dirname $0`; pwd)

expectScpFile=expect

ip="180.97.188.201"
port="4321"
password="Admin132"
file="killMysqlProcess.bin"
path="/www/sites/go"

expect $basepath/"$expectScpFile" "$ip" "$port" "$password" "$file" "$path"
# expect $basepath/"$expectSSHFile" "$ip" "$port" "$password" "$file"

# sh $basepath/reloadserver.sh