package main

import (
	"fmt"
)

func main() {
	fmt.Println(10 / 4)
	var n1 float32 = 10 / 4
	fmt.Println(n1) // n1 = 2

	var n2 float32 = 10 / 4.0
	fmt.Println(n2) // n2 = 2.5

	// %的使用
	// a % b = a - a / b * b
	fmt.Println(10 % 3)   // 1
	fmt.Println(-10 % 3)  // -1
	fmt.Println(10 % -3)  // 1
	fmt.Println(-10 % -3) // -1

	// 使用++ --
	var i = 10
	i++ // i = i + 1
	fmt.Println("i = ", i)
	i-- // i = i - 1
	fmt.Println("i = ", i)
	// i++和i-- 只能独立使用
	// 没有 --i ++i 的使用方法

}
