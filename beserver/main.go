package main

import (
	"net"

	"github.com/edwardl2004/deltatre-code-test/beserver/proto/wordrepo"
	"github.com/edwardl2004/deltatre-code-test/beserver/service"
	"google.golang.org/grpc"
)

func main() {
	conn, err := net.Listen("tcp", ":9090")
	if err != nil {
		return
	}

	grpcServer := grpc.NewServer()

	wordrepo.RegisterWordRepoServer(
		grpcServer,
		service.NewWordRepoService(),
	)

	if err := grpcServer.Serve(conn); err != nil {
		return
	}

}
