package main

import (
	"fmt"
)

// 定义一个结构体
// 如果结构体的字段类型是：指针 slice map 的零值都是 nil，即还没有分配空间
// 如果需要使用 需要先make 才能使用
type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              // 指针 new
	slice  []int             // 切片
	map1   map[string]string // map
}

type Monster struct {
	Name string
	Age  int
}

func main() {

	// 定义一个结构体变量
	var p1 Person
	fmt.Println(p1.ptr)
	fmt.Println(p1.slice)
	fmt.Println(p1.map1)

	if p1.ptr == nil {
		fmt.Println("ok1")
	}

	if p1.slice == nil {
		fmt.Println("ok2")
	}

	if p1.map1 == nil {
		fmt.Println("ok3")
	}

	// 使用slice
	p1.slice = make([]int, 10)
	p1.slice[0] = 100

	// 使用map
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom"

	fmt.Println("p1=", p1)

	// 创建一个 Monster 结构体变量
	var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.Age = 500

	monster2 := monster1
	monster2.Name = "狐狸精"

	monster3 := Monster{
		Name: "qqq",
		Age:  466,
	}

	fmt.Println("monster1=", monster1)
	fmt.Println("monster2=", monster2)
	fmt.Println("monster3=", monster3)

}
