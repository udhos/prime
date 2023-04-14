#!/bin/bash

version=$(go run ./cmd/prime -version | awk '{ print $2 }' | awk -F= '{ print $2 }')

echo version=$version

docker build --no-cache \
    -t udhos/prime:latest \
    -t udhos/prime:$version \
    -f docker/Dockerfile .