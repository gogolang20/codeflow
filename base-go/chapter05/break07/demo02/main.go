package main

import (
	"fmt"
	_ "math/rand"
	_ "time"
)

func main() {

	// 100以内的数求和 当和第一次大于 20的当前数
	// var sum int = 0
	// var max int = 100
	// for i := 1; i <= max; i++ {
	// 	sum += i
	// 	if sum > 20 {
	// 		fmt.Println("当前数=", i)
	// 		break
	// 	}
	// }
	// fmt.Println("当前数和=", sum)

	// 实现登录验证，有三次机会用户名“张无忌” 密码“888”登录成功 否则提示还有几次机会
	var name string
	var code string
	var loginChance = 3
	for i := 1; i <= 3; i++ {
		fmt.Println("请输入用户名：")
		fmt.Scanln(&name)
		fmt.Println("请输入密码：")
		fmt.Scanln(&code)

		if name == "张无忌" && code == "888" {
			fmt.Println("登录成功")
			break
		} else {
			loginChance--
			fmt.Printf("当前还有%v次机会\n", loginChance)
		}
	}

	if loginChance == 0 {
		fmt.Printf("没有登录成功")
	}

}
