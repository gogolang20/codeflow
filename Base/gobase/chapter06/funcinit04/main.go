package main

import (
	"fmt"
)

// 若引用的包有 init 函数  会先调用

// 变量
var age = test()

// 为了看到全局变量是先被初始化的 我们先写函数
func test() int {
	fmt.Println("test...")
	return 90
}

// init函数 初始化 在main之前执行
func init() {
	fmt.Println("init...")
}

func main() {
	fmt.Println("main...age=", age)
}
