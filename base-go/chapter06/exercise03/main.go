package main

import (
	"fmt"
)

// 将打印金字塔的代码封装到函数中
func tower(totalLevel int) {
	// i表示层数
	for i := 1; i <= totalLevel; i++ {
		// 打印*之前 先打印空格" "
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}
		// j表示*个数
		for j := 1; j <= 2*i-1; j++ {
			// 判断以第一个和最后一个输出*
			if j == 1 || j == 2*i-1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	// 如何打印空心菱形？？？
}

func math99(num int) {
	// 打印9*9乘法表
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v  ", j, i, (i * j))
		}
		fmt.Println()
	}
}

func main() {

	// var n int
	// fmt.Println("请输入打印金字塔的层数")
	// fmt.Scanln(&n)
	// //调用打印金字塔的函数
	// tower(n)

	// 调用乘法口诀表
	var num int
	fmt.Println("请输入打印乘法口诀表的层数")
	fmt.Scanln(&num)
	math99(num)
}
