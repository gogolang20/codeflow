package main

import (
	"fmt"
)

func main() {
	// 二维数组
	// 4是行 6是列
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	// 遍历数组
	for i := 0; i < 4; i++ {
		// fmt.Println(arr[i])
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println()

	// fmt.Println("arr", arr)//[[0 0 0 0 0 0] [0 0 1 0 0 0] [0 2 0 1 0 0] [0 0 0 0 0 0]]

	var arr2 [2][3]int // 分析 内存布局
	arr2[1][1] = 10
	fmt.Println(arr2)

	// 一个int 8个字节 (16进制)
	fmt.Printf("arr2[0]的地址%p\n", &arr2[0]) // 0xc042072030
	fmt.Printf("arr2[1]的地址%p\n", &arr2[1]) // 0xc042072048

	fmt.Printf("arr2[0][0]的地址%p\n", &arr2[0][0]) // 0xc042072030
	fmt.Printf("arr2[1][0]的地址%p\n", &arr2[1][0]) // 0xc042072048

	fmt.Println()
	var arr3 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr3=", arr3)

}
