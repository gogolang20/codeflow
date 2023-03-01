package main

import (
	"fmt"
	"strings"
)

// 累加器
func AddUpper() func(int) int {
	// 闭包 n不需要再次初始化了
	var n int = 10
	var str string = "hello"
	return func(x int) int {
		n += x
		str += string(byte(36))
		fmt.Println("str=", str)
		return n
	}
}

// 编写一个闭包函数
// return返回的函数与外部提供的suffix 构成了闭包
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		// 判断传入的名字是否有后缀
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	// 调用
	f := AddUpper()
	fmt.Println(f(1)) // 11
	fmt.Println(f(2)) // 13
	fmt.Println(f(3)) // 16

	// 测试makeSuffix 的使用
	// 返回一个闭包
	f2 := makeSuffix(".jpg") // 如果使用闭包完成，好处是只需要传入一次后缀
	fmt.Println("文件名处理后=", f2("winter"))
	fmt.Println("文件名处理后=", f2("bird.jpg"))
}
