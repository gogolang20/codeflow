package main

import (
	"fmt"
	"math"
)

func main() {
	// 多分支应用 if-else
	var score int
	fmt.Println("请输入成绩：")
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Println("奖励一辆BMW")
	} else if score <= 99 && score > 88 {
		fmt.Println("奖励一个Iphone")
	} else if score <= 80 && score >= 60 {
		fmt.Println("奖励一个Ipad")
	} else {
		fmt.Println("什么都不奖励")
	}

	// 求方程的根
	// 分析：
	// a b c 三个float64
	// 引入一个math 包 math.Sqrt()
	// 使用多分支

	var a float64 = 2.0
	var b float64 = 4.0
	var c float64 = 2.0

	m := b*b - 4*a*c
	if m > 0 {
		x1 := (-b + math.Sqrt(m)/2*a)
		x2 := (-b - math.Sqrt(m)/2*a)
		fmt.Printf("x1=%v x2=%v \n", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)/2*a)
		fmt.Printf("x1=%v \n", x1)
	} else {
		fmt.Printf("无解 \n")
	}

}
