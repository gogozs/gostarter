package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/gogozs/zlib/xlog"
	"github.com/gogozs/zlib/xrpc/client"

	demo "github.com/gogozs/gostarter/internal/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:10091", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	rpcClient, err := client.NewRpcClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		xlog.Fatal("NewRpcClient fail: %+v", err)
	}

	c := demo.NewGreeterClient(rpcClient.Conn())

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &demo.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
