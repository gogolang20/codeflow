package main

import (
	"fmt"
)

/*
	递归的退出条件：base case
	任何递归都可以改成非递归
*/

// 求arr中的最大值
func getMax(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	return process(arr, 0, len(arr)-1)
}

// arr[L..R]范围上求最大值  L ... R   N
func process(arr []int, L, R int) int {
	// arr[L..R]范围上只有一个数，直接返回，base case
	if L == R {
		return arr[L]
	}
	// L...R 不只一个数
	// mid = (L + R) / 2
	mid := L + ((R - L) >> 1)
	leftMax := process(arr, L, mid)
	rightMax := process(arr, mid+1, R)
	return Max(leftMax, rightMax)
}

func main() {
	arr := []int{1, 2, 5, 7, 8, 9, 0}
	fmt.Println(getMax(arr))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
