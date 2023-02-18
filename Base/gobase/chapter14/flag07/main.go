package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义几个变量  用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int
	// 非常重要！！！
	flag.StringVar(&user, "u", "", "user name empty")
	flag.StringVar(&pwd, "pwd", "", "password name empty")
	flag.StringVar(&host, "h", "localhost", "localhost name empty")
	flag.IntVar(&port, "port", 3306, "empty")
	flag.Parse()
	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)
}
