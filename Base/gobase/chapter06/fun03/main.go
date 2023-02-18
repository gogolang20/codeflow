package main

import (
	"fmt"
)

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("test n=", n)
}

func test2(n int) {
	if n > 2 {
		n-- // 递归必须向推出递归条件逼近 否则就是无线循环调用
		test2(n)
	} else {
		fmt.Println("test2 n=", n)
	}

}

func main() {

	// 递归函数
	// test(4)
	test2(4)
}
