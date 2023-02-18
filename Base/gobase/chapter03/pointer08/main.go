package main

import (
	"fmt" // 前面加_ 表示暂时不使用 也不会报错

	_ "strconv"
	_ "unsafe"
)

// golang 中的指针类型
func main() {

	var i int = 10
	// i 的地址是多少
	fmt.Println("i的地址是", &i)

	// var ptr *int = &i
	// ptr 是一个指针变量
	// ptr 的类型 *int
	// ptr 本身的值&i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr 的地址是 %v\n", &ptr)
	// 通过ptr 取出 i 的值
	fmt.Printf("ptr 指向的值=%v\n", *ptr)
	*ptr = 200
	fmt.Printf("i 的值=%v\n", i)
}
