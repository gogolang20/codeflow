package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 归并排序
// 归并排序基于分治思想
// 时间复杂度 O(N * logN)
func mergeSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	process(arr, 0, len(arr)-1)
}

func process(arr []int, L int, R int) {
	if L == R {
		return
	}
	mid := L + (R-L)>>1
	process(arr, L, mid)
	process(arr, mid+1, R)
	merge(arr, L, mid, R)
}

func merge(arr []int, L, mid, R int) {
	temp := make([]int, R-L+1)
	i := L                   // i 数组左下标
	j := mid + 1             // j 数组右下标
	index := 0               // temp 数组下标
	for i <= mid && j <= R { // 左右下标都不越界
		if arr[i] < arr[j] {
			temp[index] = arr[i]
			i++
		} else { // 相等先拷贝右边
			temp[index] = arr[j]
			j++
		}
		index++
	}
	copy(temp[index:], arr[i:mid+1])
	copy(temp[index:], arr[j:R+1])
	copy(arr[L:R+1], temp[:])
}

// 非递归方法实现
func mergeSort2(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	N := len(arr)
	// 步长
	mergeSize := 1
	for mergeSize < N { // log N
		// 当前左组的，第一个位置
		L := 0
		for L < N {
			if mergeSize >= N-L {
				break
			}
			M := L + mergeSize - 1
			R := M + Min(mergeSize, N-M-1)
			merge(arr, L, M, R)
			L = R + 1
		}
		// 防止溢出
		if mergeSize > N/2 {
			break
		}
		mergeSize <<= 1
	}
}

func main() {
	testTime := 100000 // 测试次数
	maxSize := 1000    // 数组长度
	maxValue := 1000   // 数组值的范围
	succeed := true
	for i := 0; i < testTime; i++ {
		arr := generateRandomArray(maxSize, maxValue)
		arr1 := make([]int, maxSize)
		copy(arr1, arr)

		mergeSort(arr1)
		mergeSort2(arr)
		// Insertion(arr)

		for i := 0; i < maxSize; i++ {
			if arr1[i] != arr[i] {
				succeed = false
				// fmt.Println(arr)
				break
			}
		}
	}
	if succeed == false {
		fmt.Println("Oops")
	} else {
		fmt.Println("Success")
	}
}

func generateRandomArray(maxSize, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, maxSize)
	for index := range arr {
		arr[index] = rand.Intn(maxValue)
	}
	return arr
}

func Insertion(arr []int) {
	if len(arr) < 2 || arr == nil {
		return
	}
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
