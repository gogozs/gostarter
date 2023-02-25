package demo

import (
	"context"

	rpc "github.com/gogozs/zlib/xrpc/server"

	"google.golang.org/grpc"

	demo "github.com/gogozs/gostarter/internal/api"
)

type GreeterService struct {
	demo.UnimplementedGreeterServer
}

var _ demo.GreeterServer = (*GreeterService)(nil)

func (s GreeterService) SayHello(ctx context.Context, request *demo.HelloRequest) (*demo.HelloReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s GreeterService) SayHelloAgain(ctx context.Context, request *demo.HelloRequest) (*demo.HelloReply, error) {
	//TODO implement me
	panic("implement me")
}

func NewGreetServer(addr string) *rpc.RpcServer {
	return rpc.NewRpcServer(addr, func(server *grpc.Server) {
		srv := GreeterService{}
		demo.RegisterGreeterServer(server, srv)
	})
}
