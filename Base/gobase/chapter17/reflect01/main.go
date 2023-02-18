package main

import (
	"fmt"
	"reflect"
)

// 编写一个函数 专门演示反射
func reflectTest01(b interface{}) {
	// 通过反射获取到传入变量的 type kind 值
	// 先获取到 reflect.Type
	// rType 是 reflect.Type 反射的 类型
	rType := reflect.TypeOf(b)
	fmt.Println("rType", rType)

	// 获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	// 如何取出真正的值
	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)

	fmt.Println("rVal", rVal)
	fmt.Printf("rVal type=%T \n", rVal)

	// 将rVal 转成 interface{}
	iV := rVal.Interface()
	// 将 interface{} 通过类型断言转换成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)

}

// 演示结构体反射
func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType", rType)

	// 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal", rVal)

	// 获取变量对应的 kind
	kind1 := rVal.Kind()
	rType.Kind()
	fmt.Printf("kind =%v kind = %v\n", kind1, rType.Kind())

	iV := rVal.Interface()
	fmt.Printf("iV=%v iv=%T\n", iV, iV)
	// 将 interface{} 通过类型断言转换成需要的类型
	// 方式一  stu, ok := iV.(Student)
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}

	// 方式二 关联第十一章的类型断言 switch 判断
}

type Student struct {
	Name string
	Age  int
}

func main() {
	// 基本数据类型的反射
	var num int = 100
	reflectTest01(num)

	fmt.Println()
	// 定义一个 Student
	stu := Student{
		Name: "tom",
		Age:  29,
	}

	reflectTest02(stu)
}
