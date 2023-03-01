package model

import (
	"fmt"
)

type account struct {
	accountNo string
	pwd       string
	balance   float64
}

func NewAccount(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("输入账户错误")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("输入密码错误")
		return nil
	}
	if balance < 20 {
		fmt.Println("输入数目错误")
		return nil
	}

	return &account{
		accountNo: accountNo,
		pwd:       pwd,
		balance:   balance,
	}
}

func (a *account) SetAccountNo(str string) {

}

func (a *account) SetBalance() {

}

func (a *account) SetPwd() {

}
