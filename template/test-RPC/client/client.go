package main

import (
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type MyClient struct {
	c *rpc.Client
}

func InitClient(addr string) MyClient {
	conn, _ := jsonrpc.Dial("tcp", addr)
	return MyClient{c: conn}
}

func (m *MyClient) SayHello(a string, b *string) error {
	return m.c.Call("hello.SayHello", a, b)
}

func main() {
	// conn, err := rpc.Dial("tcp", "127.0.0.1:8800")
	// conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8800")
	// if err != nil {
	// 	fmt.Println("Listen error:")
	// 	return
	// }

	var reply string

	myClient := InitClient("127.0.0.1:8800")
	err := myClient.SayHello("Mr 张", &reply)
	if err != nil {
		fmt.Println("Listen error:")
		return
	}
	fmt.Println(reply)
}
