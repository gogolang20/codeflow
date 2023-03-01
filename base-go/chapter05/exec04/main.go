package main

import (
	"fmt"
	_ "math"
)

func main() {
	// 输出成绩和性别
	var second float64

	fmt.Println("请输入秒数：")
	fmt.Scanln(&second)

	if second < 8 {
		// 进入决赛 男子组还是女子组
		var gender string
		fmt.Println("请输入性别：")
		fmt.Scanln(&gender)
		if gender == "男" {
			fmt.Println("进入男子决赛组")
		} else {
			fmt.Println("进入女子决赛组")
		}
	} else {
		fmt.Println("被淘汰！！！")
	}

}
