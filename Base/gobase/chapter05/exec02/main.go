package main

import (
	"fmt"
)

func main() {
	// 练习5
	var n1 int32 = 23
	var n2 int32 = 59
	if n1+n2 >= 50 {
		fmt.Println("hello world")
	}

	// 练习6
	var n3 float32 = 23.0
	var n4 float32 = 83.0
	if n3 > 10.0 && n4 > 12.0 {
		fmt.Println("和=", (n3 + n4))
	}

	// 判断闰年
	var year int = 2400
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("今年是闰年 ")
	}
}
