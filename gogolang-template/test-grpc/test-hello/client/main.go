package main

import (
	"context"
	"fmt"

	hello_grpc "test-grpc/test-hello/pb" // proto

	"google.golang.org/grpc"
)

func main() {
	grpcConn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		fmt.Println("conn err:", err)
	}
	defer grpcConn.Close()

	client := hello_grpc.NewHelloGRPCClient(grpcConn)
	req, err := client.SayHello(context.Background(), &hello_grpc.Req{Message: "I am come from client"})
	if err != nil {
		fmt.Println("client err:", err)
	}
	fmt.Println(req.GetMessage())
}
