package main

import (
	"fmt"
)

// 管理文件 func06
// AddUpper 函数返回值是匿名函数 func(int) int
func AddUpper() func(int) int {
	var n int = 20
	return func(x int) int {
		n += x
		return n
	}

}

func main() {
	f := AddUpper()
	fmt.Println(f(1)) // 21
	// fmt.Println(f(2))
	fmt.Println(f(3)) // 如果注销掉上一行 会在第一行的基础上加3
}
