package main

import (
	"fmt"
)

func main() {
	// 循环打印输入的月份的天数 countine 实现
	// 要有判断输入的月份是否错误的语句

	// 输入天数
	// 判断是否错误
	// 继续输入
	var month int
	var year int

	for {
		fmt.Println("输入年：")
		fmt.Scanln(&year)
		fmt.Println("输入月份：")
		fmt.Scanln(&month)

		if month == 1 || month == 3 || month == 5 || month == 7 ||
			month == 8 || month == 10 || month == 12 {
			fmt.Printf("%v年 %v月有31天\n", year, month)
		} else if month == 4 || month == 6 || month == 9 ||
			month == 11 {
			fmt.Printf("%v年 %v月有30天\n", year, month)
		} else if month == 2 {
			if year%400 == 0 || (year%100 != 0 && year%4 == 0) {
				fmt.Printf("%v年 %v月有29天\n", year, month)
			} else {
				fmt.Printf("%v年 %v月有28天\n", year, month)
			}
		} else {
			fmt.Println("输入错误")
		}
		continue
	}

}
