package main

import (
	"fmt"

	"gobase/chapter10/factory09/model"
)

func main() {
	// var stu = model.Student{
	// 	Name : "tom",
	// 	Score : 78.9,
	// }

	// 通过工厂模拟调用
	var stu = model.NewStudent("jack", 99)

	fmt.Println(*stu)
	fmt.Println("name=", stu.Name, "score=", stu.GetScore())
}
