#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t sunhok2/zipserver .
docker push sunhok2/zipserver
go clean
