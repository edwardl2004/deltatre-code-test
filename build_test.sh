#!/usr/bin/env bash

# install go tools in tools.go

# format go code
echo "Running go fmt..."
gofmt -s -w ./

echo "Running static analysis..."
# https://golangci-lint.run/usage/install/
golangci-lint run ./... || exit

echo "generate stub code..."
go generate ./...

echo "Running unit tests..."
go test ./...

echo "Running flake8 lint..."
# See ./integration_tests/requirements.txt
pushd ./integration_tests/ || exit
flake8 --ignore E501,W503 . || exit
popd || exit

echo "Running integration tests..."
echo "Building application and starting..."

# Its possible that an existing process is running so clean it up just in case
pkill beapp apiapp

echo > integration_tests_be.log
echo > integration_tests_api.log
go build -ldflags="-s -w" -o beapp ./beserver/main.go && ./beapp > integration_tests_be.log 2> integration_tests_be.log &
go build -ldflags="-s -w" -o apiapp ./api/main.go && ./apiapp > integration_tests_api.log 2> integration_tests_api.log &

echo "Waiting for server to startup..."
sleep 2 

echo "Running HTTP integration tests..."
pushd ./integration_tests/ || exit

if py.test ; then
	echo -e "TESTS PASSED"
else
	echo -e "TESTS FAILED"
fi
popd || exit

pkill beapp apiapp

