package main

import (
	"fmt"
)

func main() {
	// continue 案例
	// here:
	// for i := 0; i <= 4; i++ {
	// 	//lable1: //设置了一个标签
	// 	for j := 0; j<= 10; j++ {
	// 		if j == 2 {
	// 			continue
	// 		}
	// 		fmt.Println("j=", j)
	// 	}
	// }

	// continue 实现 打印1-100之内的奇数 使用for循环 continue
	// for i := 1; i <= 100; i++ {
	// 	if i % 2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("i=",i)
	// }

	// 从键盘读入个数不确定的整数，并判断读入的正数和负数的个数，输入为0时结束程序
	// 使用for break continue 实现 positive negative

	// 定义一个个数
	var count int
	var countPositive int = 0
	var countNegative int = 0
	for {
		fmt.Println("请输入一个整数")
		fmt.Scanln(&count)
		if count > 0 {
			countPositive++
			continue
		} else if count < 0 {
			countNegative++
			continue
		} else if count == 0 {
			fmt.Printf("countPositive=%d countNegative=%d", countPositive, countNegative)
			break
		}
	}
}
