package main

import (
	"context"
	"log"
	"net/http"
	"os"

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
		log.Printf("uniqueCode: e941595a, message: registering rpc handler, error: %v\n", err)
		os.Exit(1)
	}

	log.Println("uniqueCode: edf995e6, message: registered rpc handler successfully")

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Printf("uniqueCode: ec674bb5, message: running API server, error: %v\n", err)
		os.Exit(1)
	}
}
