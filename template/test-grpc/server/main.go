package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"test-grpc/pb/person"

	"google.golang.org/grpc"
)

type personServer struct {
	person.UnimplementedSearchServiceServer
}

func (*personServer) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.GetName()
	res := &person.PersonRes{Name: `revive ` + name + ` message.`}
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
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("[server][main] listen err:", err)
	}
	defer lis.Close()

	// cred, _ := credentials.NewServerTLSFromFile("", "")
	// TLS: grpc.Creds(cred)
	// 拦截器: grpc.UnknownServiceHandler()
	s := grpc.NewServer()
	person.RegisterSearchServiceServer(s, &personServer{})

	fmt.Println("start")
	s.Serve(lis)
}
