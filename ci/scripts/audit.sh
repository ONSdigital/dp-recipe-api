#!/bin/bash -eux

export cwd=$(pwd)

pushd $cwd/dp-recipe-api
  make audit
popd 