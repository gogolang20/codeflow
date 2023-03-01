package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		// 等待客户端通过conn 发送信息
		// 如果客户端没有write[发送] 那协程阻塞在这里
		// fmt.Printf("服务器在等待客户端 %s 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) // 从conn读取
		if err != nil {
			fmt.Println("[Server] Read err: ", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("[Server] Start listen...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("[Server] listen err: ", err)
		return
	}
	defer listen.Close()

	// 循环等待客户端来链接我
	for {
		// 等待客户端连接
		fmt.Println("等特客户端来链接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("[Server] Accept error: ", err)
		} else {
			fmt.Printf("[Server] Accept conn success: [%v], Client IP: [%v]\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
}

// go build -o server server.go
// go build -o client client.go
