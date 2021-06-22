package main

import (
	"log"
	"net"
	"os"

	"github.com/edwardl2004/deltatre-code-test/beserver/proto/wordrepo"
	"github.com/edwardl2004/deltatre-code-test/beserver/service"
	"google.golang.org/grpc"
)

func main() {
	conn, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Printf("uniqueCode: 880d18d1, message: opening tcp connection, error: %v\n", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()

	wordrepo.RegisterWordRepoServer(
		grpcServer,
		service.NewWordRepoService(),
	)

	if err := grpcServer.Serve(conn); err != nil {
		log.Printf("uniqueCode: 02da9d7d, message: running gRPC server, error: %v\n", err)
		os.Exit(1)
	}

}
