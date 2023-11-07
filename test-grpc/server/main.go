package main

import (
	person2 "codeflow/test-grpc/pb/person"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type personServer struct {
	person2.UnimplementedSearchServiceServer
}

func (*personServer) Search(ctx context.Context, req *person2.PersonReq) (*person2.PersonRes, error) {
	name := req.GetName()
	res := &person2.PersonRes{Name: `revive message: ` + name + `.`}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Fatalln("no metadata.")
	}

	for key, value := range md {
		fmt.Println("key: ", key)
		fmt.Println("value: ", value)
	}

	return res, nil
}

func (*personServer) SearchIn(server person2.SearchService_SearchInServer) error {
	for {
		req, err := server.Recv()
		fmt.Println(req)
		if err != nil {
			server.SendAndClose(&person2.PersonRes{Name: "finish!!!"})
			break
		}
	}
	return nil
}

func (*personServer) SearchOut(req *person2.PersonReq, server person2.SearchService_SearchOutServer) error {
	name := req.Name
	for i := 0; ; i++ {
		if i > 10 {
			break
		}
		time.Sleep(time.Second)
		server.Send(&person2.PersonRes{Name: "what i get " + name})
	}
	return nil
}

func (*personServer) SearchIO(server person2.SearchService_SearchIOServer) error {
	str := make(chan string)
	i := 0
	go func(i int) {
		for ; ; i++ {
			req, err := server.Recv()
			if err != nil {
				str <- "finish"
				break
			}
			if i > 10 {
				str <- "finish"
				break
			}
			str <- req.Name
		}
	}(i)

	for {
		s := <-str
		if s == "finish" {
			server.Send(&person2.PersonRes{Name: s})
			break
		}
		server.Send(&person2.PersonRes{Name: s})
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("[server][main] listen err: ", err)
	}
	defer listen.Close()

	// TODO TLS
	// cred, _ := credentials.NewServerTLSFromFile("", "")
	// creds := grpc.Creds(cred)

	// 拦截器
	interceptor := grpc.UnaryInterceptor(MyUnaryServerInterceptor)
	// s := grpc.NewServer(interceptor, creds)
	s := grpc.NewServer(interceptor)
	person2.RegisterSearchServiceServer(s, &personServer{})

	fmt.Println("Start server...")
	s.Serve(listen)
}

// 拦截器
func MyUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("[MyUnaryServerInterceptor] no metadata error.")
		return nil, status.Error(codes.Unauthenticated, "认证失败")
	}
	token, ok := md["token"]
	if !ok {
		fmt.Println("[MyUnaryServerInterceptor] no token error.")
		return nil, status.Error(codes.Unauthenticated, "没有 token, 认证失败")
	}
	fmt.Printf("[MyUnaryServerInterceptor] token: [%s]", token)

	now := time.Now()
	resp, err = handler(ctx, req)
	sub := time.Now().Sub(now)
	fmt.Println("Server 执行时间: ", sub.Milliseconds())
	return
}
