package main

import (
	"fmt"
	_ "math"
)

func main() {
	// 1~100间9的倍数 个数及总和
	// 记录个数 总和
	var max int64 = 100
	var count int64 = 0
	var sum int64 = 0
	var i int64 = 1
	for ; i <= max; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}

	fmt.Printf("count=%v sum=%v \n", count, sum)

	// 练习2
	var n1 int = 6
	for i := 0; i <= n1; i++ {
		fmt.Printf("%v + %v = %v \n", i, (n1 - i), n1)
	}
}
