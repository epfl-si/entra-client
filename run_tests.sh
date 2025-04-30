#!/bin/bash

if ! [ -f env_test ]; then
  echo "No env_test, no tests."
  exit 0
fi

source env_test
go test  -coverprofile=./coverage.out ./...

