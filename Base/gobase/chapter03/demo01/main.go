package main

import "fmt"

func main() {
	// 定义变量
	var i int
	// 赋值
	i = 10
	// 使用变量
	fmt.Println("i=", i)
	// 第一种：声明后变量不赋值，使用默认值，
	// int的默认值是0
	// 第二种：根据自行判定变量类型
	var num = 10.11
	fmt.Println("num=", num)
	// 第三中 省略 var,:= 左侧的变量不应该是已经声明过的，否则会编译错误
	// 以下 var name string  name = "tom"
	name := "tom"
	fmt.Println("name=", name)
}
