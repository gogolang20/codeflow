package main

import (
	"fmt"
)

func main() {
	// 类型断言
	var a interface{}
	var b float32 = 1.25
	a = b

	// y := a.(float32)
	// 断言处添加判断
	if y, ok := a.(float32); ok {
		fmt.Println("convert success")
		fmt.Printf("y的类型 %T y的值是 %v", y, y)
	} else {
		fmt.Println("convert fail")
	}

	fmt.Println("继续执行...")
}
