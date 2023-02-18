package model

import (
	"fmt"
)

type person struct {
	Name string
	age  int
	sal  float64
}

// 写一个工厂模式的函数 相当与构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// 为了访问age 和 sal 编写一对set 的方法和get 的方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	} else {
		fmt.Println("薪水范围不正确")
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}
