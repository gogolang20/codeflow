package main

import (
	"fmt"
)

// 演示defer
func sum(n1 int, n2 int) int {
	// 当函数执行完毕后 再从defer栈 按照先入后出的方式出栈 执行
	// defer 会将n1 n2的值也压入栈中
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)
	// 增加一句话
	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)
}
