package main

import "fmt"

// 二分法：可以不是有序数组，主要根据数据状况
// 1 在一个有序数组中，找某个数是否存在
// 2 在一个有序数组中，找>=某个数最左侧的位置
// 3 在一个有序数组中，找<=某个数最右侧的位置

// 在一个有序数组中，找>=某个数最左侧的位置
func BSExist(arr []int, num int) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}
	L := 0
	R := len(arr) - 1
	mid := 0
	for L < R { // 至少有两个数才二分
		mid = L + ((R - L) >> 1) // 中点位置
		if arr[mid] == num {
			return true
		} else if arr[mid] > num {
			R = mid - 1
		} else {
			L = mid + 1 // 砍去左边一半
		}
	}
	return arr[L] == num
}

// 在arr上，找满足>=value的最左位置
func BSNearLeft(arr []int, value int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	l, r, mid := 0, len(arr)-1, 0
	index := -1  // 记录最左的对号
	for l <= r { // 至少一个数的时候
		mid = l + ((r - l) >> 1)
		if arr[mid] >= value {
			index = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return index
}

// 在arr上，找满足<=value的最右位置
func BSNearRight(arr []int, value int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	l, r, mid := 0, len(arr)-1, 0
	index := -1
	for l <= r {
		mid = l + ((r - l) >> 1)
		if arr[mid] <= value {
			index = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return index
}

func main() {
	// 数组必须有序
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res := BSExist(arr, 5)
	fmt.Println(res)
}
