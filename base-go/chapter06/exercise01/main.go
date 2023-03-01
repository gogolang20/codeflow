package main

import (
	"fmt"
)

// 斐波那契数
func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

// 求函数式 表达式
func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}

// 猴子吃桃问题
func eat(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入的天数不对")
		return 0 // 表示没有得到正确数量
	}
	if n == 10 {
		return 1
	} else {
		return 2 * (eat(n+1) + 1)
	}

}

func main() {
	// res := fbn(3)
	// //测试
	// fmt.Println("res=", res)
	// fmt.Println("res=", fbn(4))
	// fmt.Println("res=", fbn(5))
	// fmt.Println("res=", fbn(6))
	// fmt.Println("res=", fbn(7))

	// 测试f
	// fmt.Println("f(1)=", f(1))
	// fmt.Println("f(5)=", f(5))

	// 测试 猴子吃桃子
	n := eat(12)
	fmt.Println("n=", n)
}
