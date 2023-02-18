package main

import (
	"fmt"
)

// 打印菱形
func main() {
	var level int = 8
	// 上半部分
	// 控制高度
	for i := 1; i <= level; i++ {
		// 控制前面的空格数
		for k := 1; k < level+1-i; k++ {
			fmt.Print(" ")
		}
		// 控制每行 * 的个数
		for j := 1; j <= 2*i-1; j++ {
			//
			if j == 1 || j == 2*i-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	// 下半部分
	// 控制高度
	for i := 1; i <= level; i++ {
		// 控制前面的空格数
		for k := 1; k < i; k++ {
			fmt.Print(" ")
		}
		// 控制每行 * 的个数
		for j := 1; j <= 2*(level-i)+1; j++ {
			if j == 1 || j == 2*(level-i)+1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
