package main

import (
	"fmt"
)

func main() {
	// 编写一个程序 可以输入人的年龄 如果该同志的年龄大于18 输出
	// "你的年龄大于18 要对自己的行为负责"

	// 分析：
	// 年龄
	// 接收一个输入
	// 输出
	var age int
	fmt.Println("请输入你的年龄:")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("你的年龄大于18 要对自己的行为负责")
	}
}
