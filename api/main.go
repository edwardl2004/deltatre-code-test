package main

import (
	"context"
	"log"
	"net/http"

	"github.com/edwardl2004/deltatre-code-test/api/proto/wordrepo"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	if err := wordrepo.RegisterWordRepoHandlerFromEndpoint(ctx, mux, "localhost:9090", []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		log.Fatalf("failed registering rpc handler: %v", err)
		return
	}

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatalf("failed running API server: %v", err)
	}
}
