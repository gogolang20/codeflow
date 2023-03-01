package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type MyInterface interface {
	SayHello(string, *string) error
}

func myRegister(i MyInterface) error {
	return rpc.RegisterName("hello", i)
}

type World struct{}

func (w *World) SayHello(name string, resp *string) error {
	*resp = name + "你好！"
	return nil
}

func main() {
	// err := rpc.RegisterName("hello", new(World))
	// err := rpc.Register(new(World))
	err := myRegister(new(World))
	if err != nil {
		fmt.Println("Register error:")
		return
	}
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Listen error:")
		return
	}
	defer listener.Close()

	println("Start Listen...")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept error:")
		return
	}
	defer conn.Close()

	println("Start Accept...")

	// rpc.ServeConn(conn)
	jsonrpc.ServeConn(conn)

}
