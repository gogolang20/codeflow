package main

import (
	"context"
	"fmt"

	"test-grpc/test-server/pb/person"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	/*
		normal
	*/
	grpcConn, err := grpc.Dial("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("conn err:", err)
	}
	defer grpcConn.Close()

	client := person.NewSearchServiceClient(grpcConn)
	res, err := client.Search(context.Background(), &person.PersonReq{Name: "i am scs"})
	if err != nil {
		fmt.Println("call err:", err)
	}
	fmt.Println("res:", res)

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
