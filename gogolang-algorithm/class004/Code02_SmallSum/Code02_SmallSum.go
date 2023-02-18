package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
面试问题一  小和问题
数组每个左边位置比自己小的数都累加
最后整体再累加起来，时间复杂度O(N * logN)
每个merge左组比右组数小的时候产生 小和，
左右相等的时候先拷贝右组，不产生 小和
*/
func smallSum(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return process(arr, 0, len(arr)-1)
}

func process(arr []int, L int, R int) int {
	if L == R {
		return 0
	}
	mid := L + (R-L)>>1
	return process(arr, L, mid) + process(arr, mid+1, R) + merge(arr, L, mid, R)
}

func merge(arr []int, L, mid, R int) int {
	temp := make([]int, R-L+1)
	i, j := L, mid+1
	index := 0
	res := 0 // 记录最小和结果
	for i <= mid && j <= R {
		if arr[i] < arr[j] { // 相等的时候，先拷贝右边!!!
			res += (R - j + 1) * arr[i]
			temp[index] = arr[i]
			i++
		} else {
			temp[index] = arr[j]
			j++
		}
		index++
	}
	copy(temp[index:], arr[i:mid+1])
	copy(temp[index:], arr[j:R+1])
	copy(arr[L:R+1], temp[:])
	return res // 返回最小和结果
}

// for test
func comparator(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	res := 0
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				res += arr[j]
			}
		}
	}
	return res
}

func main() {
	arr := make([]int, 50)
	arr1 := make([]int, 50)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	copy(arr1, arr)

	fmt.Println(smallSum(arr))
	fmt.Println(comparator(arr1))
}
