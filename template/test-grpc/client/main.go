package main

import (
	"context"
	"fmt"
	"time"

	"codeflow/template/test-grpc/pb/person"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func MyUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	pairs := metadata.Pairs("token", "token999")
	ctx = metadata.NewOutgoingContext(context.Background(), pairs)

	now := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	sub := time.Now().Sub(now)
	fmt.Println("Client 执行时间: ", sub.Milliseconds())
	return err
}

func main() {
	/*
		normal
	*/
	// 拦截器
	opt := grpc.WithUnaryInterceptor(MyUnaryClientInterceptor)
	grpcConn, err := grpc.Dial("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()), opt)
	if err != nil {
		fmt.Println("conn err: ", err)
	}
	defer grpcConn.Close()

	// metadata
	pairs := metadata.Pairs("token1", "token999")
	ctx := metadata.NewOutgoingContext(context.Background(), pairs)

	client := person.NewSearchServiceClient(grpcConn)
	// res, err := client.Search(context.Background(), &person.PersonReq{Name: "i am scs"})
	res, err := client.Search(ctx, &person.PersonReq{Name: "I am scs"})
	if err != nil {
		fmt.Println("call err: ", err)
	}
	fmt.Println("res: ", res)

	/*
		In
	*/
	// grpcConn, err := grpc.Dial("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	fmt.Println("[client][main] dial err:", err)
	// }
	// defer grpcConn.Close()
	//
	// client := person.NewSearchServiceClient(grpcConn)
	// c, err := client.SearchIn(context.Background())
	// if err != nil {
	// 	fmt.Println("[client][main] call err:", err)
	// }
	// for i := 0; ; i++ {
	// 	time.Sleep(time.Second)
	// 	c.Send(&person.PersonReq{Name: "come in message"})
	// 	if i > 10 {
	// 		res, err := c.CloseAndRecv()
	// 		if err != nil {
	// 			fmt.Println("[][] err: ", err)
	// 		}
	// 		fmt.Println(res)
	// 		break
	// 	}
	// }

	/*
		Out
	*/
	// grpcConn, err := grpc.Dial("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	fmt.Println("[client][main] dial err:", err)
	// }
	// defer grpcConn.Close()
	//
	// client := person.NewSearchServiceClient(grpcConn)
	// c, err := client.SearchOut(context.Background(), &person.PersonReq{Name: "scs"})
	// if err != nil {
	// 	fmt.Println("call err:", err)
	// }
	// for {
	// 	req, err := c.Recv()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}
	// 	fmt.Println(req)
	// }

	/*
		IO
	*/
	// grpcConn, err := grpc.Dial("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	fmt.Println("conn err:", err)
	// }
	// defer grpcConn.Close()
	//
	// client := person.NewSearchServiceClient(grpcConn)
	// c, err := client.SearchIO(context.Background())
	// if err != nil {
	// 	fmt.Println("call err:", err)
	// }
	// var wg sync.WaitGroup
	//
	// wg.Add(2)
	// go func() {
	// 	for {
	// 		time.Sleep(time.Second)
	// 		err := c.Send(&person.PersonReq{Name: "scs"})
	// 		if err != nil {
	// 			wg.Done()
	// 			break
	// 		}
	// 	}
	// }()
	// go func() {
	// 	for {
	// 		req, err := c.Recv()
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			wg.Done()
	// 			break
	// 		}
	// 		fmt.Println(req)
	// 	}
	// }()
	// wg.Wait()
}
