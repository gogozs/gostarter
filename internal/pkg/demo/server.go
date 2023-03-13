package demo

import (
	"context"

	"github.com/gogozs/zlib/xrpc/server/serverinterceptors"

	rpc "github.com/gogozs/zlib/xrpc/server"

	"google.golang.org/grpc"

	demo "github.com/gogozs/gostarter/internal/api"
)

type GreeterService struct {
	demo.UnimplementedGreeterServer
}

var _ demo.GreeterServer = (*GreeterService)(nil)

func (s GreeterService) SayHello(ctx context.Context, request *demo.HelloRequest) (*demo.HelloReply, error) {
	return &demo.HelloReply{Message: request.Name}, nil
}

func (s GreeterService) SayHelloAgain(ctx context.Context, request *demo.HelloRequest) (*demo.HelloReply, error) {
	panic("implement me")
}

func NewGreetServer(addr string) *rpc.RpcServer {
	registerFn := func(server *grpc.Server) {
		demo.RegisterGreeterServer(server, GreeterService{})
	}
	return rpc.NewRpcServer(
		addr,
		registerFn,
		rpc.WithUnaryServerOption(serverinterceptors.LogInterceptor),
	)
}
