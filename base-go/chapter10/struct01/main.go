package main

import (
	"fmt"
)

// 定义一个结构体
type Cat struct {
	Name   string // 大写 代表公开的
	Age    int
	Color  string
	Hobby  string
	Scores [3]int
}

func main() {
	// 养猫问题

	// //使用变量
	// var cat1Name string = "小白"
	// var cat1Age int = 3
	// var cat1Color string = "白色"

	// var cat2Name string = "小花"
	// var cat2Age int = 100
	// var cat2Color string = "花色"

	// //使用数组
	// var catName [2]string = [...]string{"小白", "小花"}
	// var catAge [2]int = [...]int{3, 100}
	// var catColor [2]string = [...]string{"白色", "花色"}

	// 结构体 struct 完成

	// 创建一个Cat的变量
	var cat1 Cat

	fmt.Printf("cat1的地址=%p\n", &cat1)
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "吃鱼"
	fmt.Println("cat1", cat1)

	fmt.Println("猫的信息如下：")
	fmt.Println("name=", cat1.Name)
	fmt.Println("age=", cat1.Age)
	fmt.Println("color=", cat1.Color)
	fmt.Println("Hobby=", cat1.Hobby)
}
