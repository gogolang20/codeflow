package service

import (
	"gobase/chapter13/customerManage/model"
)

// 包括增删改查
type CustomerService struct {
	// 一个 model 下 Customer 结构体的切片
	customers []model.Customer
	// 声明一个字段 表示当前切片含有多少个客户
	customerNum int
}

func NewCustomerService() *CustomerService {
	// 初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张思", "男", 10, "165", "sx@163.com")
	customerService.customers = append(customerService.customers, *customer)
	return customerService
}

func (this *CustomerService) List() []model.Customer {
	return this.customers
}

func (this *CustomerService) Add(customer model.Customer) bool {

	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

// 根据id删除客户
func (this *CustomerService) Delete(id int) bool {

	index := this.FindById(id)

	if index == -1 { // 说明没有这个客户
		return false
	}

	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true

}

// 删除方法  先查找 再删除
func (this *CustomerService) FindById(id int) int {

	// 遍历
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}

	return index
}

// 修改功能
// func (this *CustomersService) Modify() {
// 	//
// }
