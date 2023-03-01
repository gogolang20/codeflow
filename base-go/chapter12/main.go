package main

import (
	"fmt"
)

func main() {
	// 声明一个变量
	key := ""
	// 控制是否退出for循环
	loop := true
	balance := 10000.0
	money := 0.0
	note := ""
	// 定义一个变量 是否有收支行为
	flag := false
	// 表头
	details := "收支\t账户金额\t收支金额\t说明"

	// 显示主菜单
	for {
		fmt.Println("---------家庭收支记账软件---------")
		fmt.Println("         1:收支明细")
		fmt.Println("         2:登记收入")
		fmt.Println("         3:登记支出")
		fmt.Println("         4:退出软件")
		fmt.Print("请选择(1-4): ")

		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("\n---------当前收支明显记录---------")
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("当前没有收支 来一笔吧")
			}

		// 登记收入
		case "2":
			fmt.Println("本次收入金额:")
			fmt.Scanln(&money)
			balance += money // 修改账户余额
			fmt.Println("本次收入说明:")
			fmt.Scanln(&note)
			// 收入情况拼接到 details
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v",
				balance, money, note)
			flag = true
		// 登记支出
		case "3":
			fmt.Println("本次支出金额:")
			fmt.Scanln(&money)
			// 必要的判断
			if money > balance {
				fmt.Println("支出的余额不足")
				break
			}
			balance -= money
			fmt.Println("本次支出说明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v",
				balance, money, note)
			flag = true
		case "4":
			fmt.Println("你确定要退出吗？ y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("你的输入有误，请重新输入y/n")
			}
			if choice == "y" {
				loop = false
			}
		default:
			fmt.Println("输入有误")
			break
		}

		if !loop {
			break
		}
	}

	fmt.Println("你退出了家庭记账软件的使用")
}
