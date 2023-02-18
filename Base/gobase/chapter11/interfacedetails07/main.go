package main

import (
	"fmt"
)

type BInterface interface {
	test01()
}

type CInterface interface {
	test02()
}

type AInterface interface {
	BInterface
	CInterface
	test03()
}

// 如果实现AInterface 接口 需要将BInterface  CInterface 的方法搜都实现
type Stu struct {
}

func (stu Stu) test01() {

}

func (stu Stu) test02() {

}

func (stu Stu) test03() {

}

type T interface {
}

func main() {
	var stu Stu
	var a AInterface = stu
	a.test01()

	var t T = stu
	fmt.Println(t)
	// 声明一个空接口
	// 可以接收任何一个变量
	var t2 interface{} = stu

	var num1 float64 = 8.9
	t2 = num1
	t = num1
	fmt.Println(t2, t)
}
