package main

import (
	"context"
	"fmt"
	"net"
	hello_grpc "test-grpc/test-hello/pb" // proto

	"google.golang.org/grpc"
)

// get server
type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

// guazai method
func (s *server) SayHello(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "I am the content return from service of grpc!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen err:", err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &server{})
	s.Serve(lis)
}
