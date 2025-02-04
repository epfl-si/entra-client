#!/bin/bash

if ! [ -f .env ]; then
  echo "No .env, no tests."
  exit 0
fi

source .env
go test -v ./...

