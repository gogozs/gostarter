package main

import (
	"context"
	"log"
	"net/http"

	demo2 "github.com/gogozs/gostarter/internal/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gogozs/gostarter/internal/server/demo"
)

func main() {
	log.Println("Starting listening on port 10091")
	srv := demo.NewGreetServer(":10091")

	go func() {
		if err := srv.Start(); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	var grpcServerEndpoint = "localhost:10091"
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := demo2.RegisterGreeterHandlerFromEndpoint(context.Background(),
		mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Listening on port 10092")
	port := ":10092"
	http.ListenAndServe(port, mux)
}
