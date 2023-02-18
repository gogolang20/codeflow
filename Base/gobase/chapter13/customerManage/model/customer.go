package model

import (
	"fmt"
)

// 声明一个customer 结构体
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) *Customer {
	return &Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// 第二种不带Id的方法
func NewCustomer2(name string, gender string,
	age int, phone string, email string) *Customer {
	return &Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",
		this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return info
}
