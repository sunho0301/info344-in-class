#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t sunhok2/testserver .
docker push sunhok2/testserver
go clean
