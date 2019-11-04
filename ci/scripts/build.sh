#!/bin/bash -eux

cwd=$(pwd)

export GOPATH=$cwd/go

pushd dp-recipe-api
  make build && mv build/$(go env GOOS)-$(go env GOARCH)/* $cwd/build
  cp Dockerfile.concourse $cwd/build
popd
