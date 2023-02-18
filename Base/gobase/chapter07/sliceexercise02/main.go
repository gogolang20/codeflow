package main

import (
	"fmt"
)

func fbn(n int) []uint64 {
	// 声明一个切片大小
	var fbnSlice []uint64 = make([]uint64, n)
	// 第一个 和第二个数的斐波那契 为1
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {

		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func main() {
	// 思路
	// 声明一个函数 fbn(n uint64) ([]uint64)
	// 使用for循环存放斐波那契数列

	// 测试
	fbnSlice := fbn(10)
	fmt.Println("fbnSlice=", fbnSlice)

}
