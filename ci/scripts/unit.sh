#!/bin/bash -eux

export GOPATH=$(pwd)/go

pushd dp-recipe-api
  make test
popd