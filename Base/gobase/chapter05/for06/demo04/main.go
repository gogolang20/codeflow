package main

import (
	"fmt"
	_ "math"
)

func main() {
	// 先打印矩形
	// 打印半个金字塔
	// 打印金字塔
	// 打印空心金字塔

	// 层数用变量表示
	var totalLevel int = 20

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

	// 打印9*9乘法表
	// 添加一个层数的变量
	var num int = 10
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v  ", j, i, (i * j))
		}
		fmt.Println()
	}

}
