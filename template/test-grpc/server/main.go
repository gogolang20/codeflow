package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"test-grpc/pb/person"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type personServer struct {
	person.UnimplementedSearchServiceServer
}

func (*personServer) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.GetName()
	res := &person.PersonRes{Name: `revive message: ` + name + `.`}

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

func (*personServer) SearchIn(server person.SearchService_SearchInServer) error {
	for {
		req, err := server.Recv()
		fmt.Println(req)
		if err != nil {
			server.SendAndClose(&person.PersonRes{Name: "finish!!!"})
			break
		}
	}
	return nil
}

func (*personServer) SearchOut(req *person.PersonReq, server person.SearchService_SearchOutServer) error {
	name := req.Name
	for i := 0; ; i++ {
		if i > 10 {
			break
		}
		time.Sleep(time.Second)
		server.Send(&person.PersonRes{Name: "what i get " + name})
	}
	return nil
}

func (*personServer) SearchIO(server person.SearchService_SearchIOServer) error {
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
			server.Send(&person.PersonRes{Name: s})
			break
		}
		server.Send(&person.PersonRes{Name: s})
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("[server][main] listen err: ", err)
	}
	defer listen.Close()

	// cred, _ := credentials.NewServerTLSFromFile("", "")
	// TLS: grpc.Creds(cred)

	// 拦截器
	interceptor := grpc.UnaryInterceptor(MyUnaryServerInterceptor)
	s := grpc.NewServer(interceptor)
	person.RegisterSearchServiceServer(s, &personServer{})

	fmt.Println("Start server...")
	s.Serve(listen)
}

// 拦截器
func MyUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	now := time.Now()
	resp, err = handler(ctx, req)
	sub := time.Now().Sub(now)
	fmt.Println("Server 执行时间: ", sub.Milliseconds())
	return
}
