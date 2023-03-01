package main

import (
	"fmt"
)

var (
	// Fun1 就是全局匿名函数
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	fmt.Println("Fun1=", Fun1(10, 29)) // 290
	// 定义匿名函数时直接调用 这种方式匿名函数只能调用一次
	// 案列演示 使用匿名函数的方式完成
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1=", res1)

	// 将匿名函数func (n1 int, n2 int) int 赋给a 变量
	// 则 a 的数据类型就是函数类型
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}

	res2 := a(10, 30)
	fmt.Println("res2=", res2)
	res3 := a(60, 30)
	fmt.Println("res3=", res3)

	// 全局匿名函数的使用
	res4 := Fun1(4, 9)
	fmt.Println("res4=", res4)
}
