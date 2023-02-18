package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 添加功能
// 1 能通过终端输入数据（输入一行 发送一行） 并发送给服务器端
// 2 在终端输入 exit 表示退出程序
func main() {
	// 先使用 net 包中的函数链接服务器
	conn, err := net.Dial("tcp", "192.168.0.101:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	// fmt.Println(conn)
	// 功能 客户端可以发送单行数据
	reader := bufio.NewReader(os.Stdin) // os.Stdin bn 表示从终端输入

	for {
		// 从终端读取一行用户输入 并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ReadString err=", err)
		}
		line = strings.Trim(line, " \n\r")
		if line == "exit" {
			fmt.Println("退出客户端")
			break
		}

		// 在将line 发送给服务器
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.write err=", err)
		}
		// fmt.Printf("客户端发送了 %d 字节的数据，并退出\n", n)
	}

}
