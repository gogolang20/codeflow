package main

import (
	"fmt"
	"reflect"
)

// 定义了一个 Monster 结构体
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

// 方法 显示s 的值
func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

// 方法 返回俩个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

// 方法 接收4个值 给s 赋值
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

// 反射函数
func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	// 获取到 a 对应的类别
	kd := val.Kind()
	// 判断传入的不是 struct 结构体 就退出
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	// 获取到该结构体的字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)

	// 变量结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("field %d: 值为=%v\n", i, val.Field(i))
		// 获取到 struct 标签，注意需要通过 reflect.Type 来获取 tag 标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: tag=%v\n", i, tagVal)
		}
	}

	// 获取到该结构体有多个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	// var params []reflect.Value
	// 方法的排序默认是按照函数名排序  ASCII码
	val.Method(1).Call(nil) // 获取到第二个方法

	// 调用结构体的第1个方法 Method(0)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) // 传入的参数是 []reflect.Value
	fmt.Println("res", res[0].Int())  // 返回结果 返回的结果是 []reflect.Value
}

/*
练习
func TestStruct(a  interface{}) {

}
*/
func main() {
	// 创建了一个 Monster 实例
	var a Monster = Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.2,
	}
	// 将 Monster
	TestStruct(a)
}
