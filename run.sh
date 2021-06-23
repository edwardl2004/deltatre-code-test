#!/usr/bin/env bash
echo "generate stub code..."
go generate ./...

# Its possible that an existing process is running so clean it up just in case
pkill beapp apiapp

echo > integration_tests_be.log
echo > integration_tests_api.log
go build -ldflags="-s -w" -o beapp ./beserver/main.go && ./beapp > integration_tests_be.log 2> integration_tests_be.log &
go build -ldflags="-s -w" -o apiapp ./api/main.go && ./apiapp > integration_tests_api.log 2> integration_tests_api.log &
