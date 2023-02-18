package main

import (
	"fmt"
)

type student struct {
}

// 编写一个函数 可以判断输入的类型
func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数的是 bool 类型，值是%v\n", index, x)
		case float32:
			fmt.Printf("第%v个参数的是 float32 类型，值是%v\n", index, x)
		case float64:
			fmt.Printf("第%v个参数的是 float64 类型，值是%v\n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数的是 整数 类型，值是%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数的是 string 类型，值是%v\n", index, x)
		case student:
			fmt.Printf("第%v个参数的是 student 类型，值是%v\n", index, x)
		case *student:
			fmt.Printf("第%v个参数的是 *student 类型，值是%v\n", index, x)
		default:
			fmt.Printf("第%v个参数的是 类型不确定，值是%v\n", index, x)
		}
	}

}

func main() {
	var n1 float32 = 1.2
	var n2 float64 = 1.8
	var n3 int = 100
	n4 := 300
	name := "jack"
	address := "北京"

	stu := student{}
	stu2 := &student{}

	TypeJudge(n1, n2, n3, n4, name, address, stu, stu2)
}
