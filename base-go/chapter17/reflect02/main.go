package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	// 看看 rVal 的Kind 是
	fmt.Printf("rVal Kind= %v\n", rVal.Kind())
	// rVal
	// Elem() 非常重要
	rVal.Elem().SetInt(20)
}

func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println("num=", num)

	// Elem() 类似
	// num2 := 9
	// ptr *int = &num2
	// num3 := *ptr
}
