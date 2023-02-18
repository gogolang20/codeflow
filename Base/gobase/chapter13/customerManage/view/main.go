package main

import (
	"fmt"

	"gobase/chapter13/customerManage/model"
	"gobase/chapter13/customerManage/service"
)

type customerView struct {
	// 定义必要的字段
	key             string // 接收客户输入。。。
	loop            bool   // 表示是否循环的显示菜单
	customerService *service.CustomerService
}

// 显示所有的客户
func (this *customerView) List() {
	customers := this.customerService.List()
	// 显示
	fmt.Println("---------客户列表---------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n---------客户列表完成---------\n\n")

}

// 得到用户的输入 信息构建新的客户
func (this *customerView) add() {
	fmt.Println("---------添加客户---------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scan(&name)

	fmt.Println("性别：")
	gender := ""
	fmt.Scan(&gender)

	fmt.Println("年龄：")
	age := 0
	fmt.Scan(&age)

	fmt.Println("电话：")
	phone := ""
	fmt.Scan(&phone)

	fmt.Println("邮箱：")
	email := ""
	fmt.Scan(&email)

	// Id号没有让用户输入 系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	// 调用
	if this.customerService.Add(*customer) {
		fmt.Println("---------添加完成---------")
	} else {
		fmt.Println("---------添加失败---------")
	}

}

func (this *customerView) delete() {
	fmt.Println("---------删除客户---------")
	fmt.Println("请选择待删除客户编号(-1退出):")
	id := -1
	fmt.Scan(&id)
	if id == -1 {
		return
	}

	fmt.Println("确认是否删除(Y/N):")
	choice := ""
	fmt.Scan(&choice)

	if choice == "y" || choice == "Y" {
		// 调用customerService 的Delete方法
		if this.customerService.Delete(id) {
			fmt.Println("---------删除成功---------")
		} else {
			fmt.Println("---------删除失败 输入的id不存在---------")
		}
	}

}

func (this *customerView) exit() {
	fmt.Println("确认是否退出(Y/N)")
	for {
		fmt.Scan(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}

		fmt.Println("你的输入有误，确认是否退出(Y/N) ")
	}

	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}
}

// 与结构体绑定的方法
func (this *customerView) mainMenu() {

	for {
		fmt.Println("---------客户信息管理软件---------")
		fmt.Println("         1 添加客户")
		fmt.Println("         2 修改客户")
		fmt.Println("         3 删除客户")
		fmt.Println("         4 客户列表")
		fmt.Println("         5 退    出")
		fmt.Print("请选择(1-5):")

		fmt.Scan(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Print("1") // this.Mdify()
		case "3":
			this.delete()
		case "4":
			this.List()
		case "5":
			this.exit()
		default:
			fmt.Println("输入错误 请重新输入...")
		}

		if !this.loop {
			break
		}
	}

	fmt.Println("你退出了客户关系管理系统")

}

func main() {
	// 在主函数中创建一个customerView 并运行
	customerView := customerView{
		key:  "",
		loop: true,
	}
	//
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}
