package main

import (
	"fmt"
)

func main() {
	// 二维数组遍历
	var arr [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	// 传统方式for
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v\t", arr[i][j])
		}
		fmt.Println()
	}

	// for i, v := range 方式
	// v 代表底下的一维数组 i是一位数组的个数
	for i, v := range arr {
		// v2 是一维数组中的值 j是一维数组的下标
		for j, v2 := range v {
			fmt.Printf("arr[%v][%v]=%v\t", i, j, v2)
		}
		fmt.Println()
	}

}
