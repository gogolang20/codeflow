package utils

import (
	"fmt"
)

type FamilyAccount struct {
	key     string
	loop    bool
	balance float64
	money   float64
	note    string
	flag    bool
	details string
}

// 编写一个工厂模式的构造方法，返回一个*FamilyAccount 实例
func NewFamilyAccount() *FamilyAccount {

	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说明",
	}
}

// 将显示明细写成一个方法
func (this *FamilyAccount) showDetails() {
	fmt.Println("\n---------当前收支明显记录---------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支 来一笔吧")
	}
}

// 将登记收入写个一个方法，和 *FamilyAccount绑定
func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balance += this.money // 修改账户余额
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	// 收入情况拼接到 details
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v",
		this.balance, this.money, this.note)
	this.flag = true
}

// 将登记支出写个一个方法，和 *FamilyAccount绑定
func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	// 必要的判断
	if this.money > this.balance {
		fmt.Println("支出的余额不足")
		// break
	}
	this.balance -= this.money
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v",
		this.balance, this.money, this.note)
	this.flag = true
}

// 将退出写个一个方法，和 *FamilyAccount绑定
func (this *FamilyAccount) exit() {
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
		this.loop = false
	}
}

// 显示主菜单方法
func (this *FamilyAccount) MainMemu() {
	for {

		fmt.Println("---------家庭收支记账软件---------")
		fmt.Println("         1:收支明细")
		fmt.Println("         2:登记收入")
		fmt.Println("         3:登记支出")
		fmt.Println("         4:退出软件")
		fmt.Print("请选择(1-4): ")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()

		// 登记收入
		case "2":
			this.income()
		// 登记支出
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("输入有误")
			break
		}
		if !this.loop {
			break
		}

	}
}
