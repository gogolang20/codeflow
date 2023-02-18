package main

import (
	"fmt"
	"gobase/chapter11/encapexe03/model"
)

func main() {
	a := model.NewAccount("gs111111", "666666", 72.0)
	if a != nil {
		fmt.Println("创建成功", a)
	} else {
		fmt.Println("创建失败")
	}

	a.SetAccountNo("111111")
	a.SetBalance()
	a.SetPwd()
}
