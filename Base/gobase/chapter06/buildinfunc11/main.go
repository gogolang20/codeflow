package main

import (
	"fmt"
	_ "strconv"
	_ "time"
)

// 内置函数 builtin 中的 new
func main() {
	// 内置函数buildin
	num1 := 100
	// num1的类型是int num1的值是100 num1的地址是0xc042056058
	fmt.Printf("num1的类型是%T num1的值是%v num1的地址是%v\n", num1, num1, &num1)

	num2 := new(int) // *int
	// num2的类型是 *int
	// num2的值是 地址 是由系统分配的
	// num2的地址是 地址 是由系统分配的
	*num2 = 10
	fmt.Printf("num2的类型是%T num2的值是%v num2的地址是%v num2指向的值是%v\n",
		num2, num2, &num2, *num2)
	// num2的类型是*int num2的值是0xc0420560a0 num2的地址是0xc042074020 num2指向的值是10
}
