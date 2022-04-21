#!/usr/bin/env bash
export GOOS=linux
export CGO_ENABLED=0
go mod tidy
go build  -o service-linux-amd64 .
echo built $(pwd)
echo "go build success!!!"

docker build -t gorpher/service-linux-amd64:latest .
rm -f service-linux-amd64
