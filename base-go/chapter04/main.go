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
}
