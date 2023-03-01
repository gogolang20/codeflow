package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world!")
	}
}

func test() {
	// 捕捉错误 defer + recover
	defer func() {
		// 捕获 test 报出的 panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang" // error 还没有make 初始化
}

func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("main() hello world!")
	}
}
