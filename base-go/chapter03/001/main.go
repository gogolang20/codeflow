package main

import (
	"fmt"
)

// 演示golang中字符类型使用
func main() {
	var c1 byte = 'a'
	var c2 byte = '0'

	// 当我们直接输出byte值，就是输出了对应的字符的码值 ASCII码表
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	// 如果希望输出对应的字符需要使用格式化的输出
	fmt.Printf("c1=%c c2=%c\n", c1, c2)

	// var c3 byte ='北'//overflows byte
	var c4 int = '北'
	fmt.Printf("c4=%c c4的码值是 %d\n", c4, c4)

	var c5 int = 22269 // 22269 对应”国“
	fmt.Printf("c5=%c\n", c5)

	// 与字符相加
	var n1 = 10 + 'a'       // 10 + 97 =107
	fmt.Printf("n1=%T", n1) // n1 类型是int32
}
