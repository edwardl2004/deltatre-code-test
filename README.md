# deltatre-code-test

The system contains two services, backend server, which is the gRPC server and implements the functions exposed via gRPC; 
and api server, which implments Rest APIs for the functions of backend server. The codes of two services are in `beserver/`
and `api/` folders, respectively.

## Install Go tools and Python requirements
The build process requires a number of Go tools. To install them, find the tools.go, and install all the tools listed in the file:
`go installgithub.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway`
`go installgithub.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2`
`go installgoogle.golang.org/genproto/googleapis/api`
`go installgoogle.golang.org/grpc/cmd/protoc-gen-go-grpc`
`go installgoogle.golang.org/protobuf/cmd/protoc-gen-go`

The integration testing code requires some python libraries. To install them, run the command:
`pip install -r ./integration_tests/requirements.txt`

## Build and test the code
run the script to build and test the code:
`./build_test.sh`

## Start the backend server and api server
run the script to start the backend server and api server:
`./run.sh`
