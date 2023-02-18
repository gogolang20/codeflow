package main

import (
	"fmt"
)

var age int = 50          // 整个包可以使用
var Name string = "jack~" // 其他包也可以使用 首字母大写 作用域在整个程序有效

func test() {
	// age Name 的作用域就只在test函数内部
	age := 10
	Name := "tom~"
	fmt.Println("age=", age)   // 10
	fmt.Println("Name=", Name) // tom
}

func main() {

	fmt.Println("age=", age)   // 50
	fmt.Println("Name=", Name) // jack
	test()

	// 如果变量在一个代码块中 比如 for if作用域就是 该代码块
	for i := 0; i <= 10; i++ {
		fmt.Println("i=", i)
	}
}
