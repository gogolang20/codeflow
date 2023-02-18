package main

import (
	"fmt"
)

// 一个账户字段的结构体
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// 结构体绑定的方法
// 存款 取款 查询
func (account *Account) Deposite(Pwd string, money float64) {
	if account.Pwd != Pwd {
		fmt.Println("输入密码错误")
		return
	}
	if money <= 0 {
		fmt.Println("输入金额错误")
		return
	}
	account.Balance += money
	fmt.Println("存款成功")
}

func (account *Account) WithDraw(Pwd string, money float64) {
	if account.Pwd != Pwd {
		fmt.Println("输入密码错误")
		return
	}
	if money <= 0 || money > account.Balance {
		fmt.Println("输入金额错误")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功")
}

// 查询
func (account *Account) Query(Pwd string) {
	if account.Pwd != Pwd {
		fmt.Println("输入密码错误")
		return
	}
	fmt.Printf("你的账户为：%v 余额：%v", account.AccountNo, account.Balance)
}

func main() {
	var money = Account{
		AccountNo: "gs111111",
		Pwd:       "666666",
		Balance:   100.0,
	}

	money.Query("666666")

	fmt.Println()
	money.Deposite("666666", 200)
	money.Query("666666")

	fmt.Println()
	money.WithDraw("666666", 150)
	money.Query("666666")
}
