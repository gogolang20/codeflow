package main

import (
	"fmt"
)

func modify(map1 map[int]int) {
	map1[10] = 900
}

// 定义一个学生结构体
type Stu struct {
	Name    string
	Age     int
	Address string
}

func main() {

	// map是引用类型
	map1 := make(map[int]int)
	map1[1] = 80
	map1[2] = 69
	map1[10] = 4
	map1[20] = 8
	modify(map1)
	fmt.Println(map1)
	// map的value 经常使用struct 类型 更适合管理复杂数据
	// map 的key 为学生的学号
	// map 的value 为struct

	// map
	students := make(map[string]Stu)
	// 创建学生
	stu1 := Stu{"tom", 19, "上海"}
	stu2 := Stu{"mary", 29, "深圳"}
	students["no1"] = stu1
	students["no2"] = stu2

	fmt.Println(students)
	// 遍历各个学生的信息
	for k, v := range students {
		fmt.Printf("学生的编号是%v\n", k)
		fmt.Printf("学生的名字是%v\n", v.Name)
		fmt.Printf("学生的年龄是%v\n", v.Age)
		fmt.Printf("学生的住址是%v\n", v.Address)
		fmt.Println()
	}
}
