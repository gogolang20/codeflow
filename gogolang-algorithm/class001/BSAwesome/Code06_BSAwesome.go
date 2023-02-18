package main

import "fmt"

// 局部最小值问题
// 无序的数组，任意两个相邻的数不相等
// 只需要返回一个下标即可

func getLessIndex(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1 // no exist
	}
	if len(arr) == 1 || arr[0] < arr[1] {
		return 0
	}
	if arr[len(arr)-1] < arr[len(arr)-2] {
		return len(arr) - 1
	}
	left := 1             // 左侧开始下标
	right := len(arr) - 2 // 右侧开始下标
	mid := 0
	for left < right { // 范围内至少有两个数
		mid = left + (right-left)/2
		if arr[mid] > arr[mid-1] {
			right = mid - 1
		} else if arr[mid] > arr[mid+1] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left // 范围内至少有两个数，所以要加
}

func main() {
	var arr2 = []int{6, 5, 1, 2, 3, 2, 3, 4, 3, 4, 6, 4, 5, 9, 6, 2, 5, 7}
	fmt.Println(len(arr2))
	num := getLessIndex(arr2)
	fmt.Printf("num=%d \n", num)
}
