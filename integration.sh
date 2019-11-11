#!/bin/bash -xe

ENDPOINT=${ENDPOINT:-http://localhost:3000}

# install dependencies
go mod tidy

# run integration tests
go test ./integration --endpoint=${ENDPOINT}