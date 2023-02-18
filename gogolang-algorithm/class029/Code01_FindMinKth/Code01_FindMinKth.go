package main

import (
	"fmt"
	"math/rand"
	"time"
)

//import (
//)

/*
在无序数组中求第K小的数

1）改写快排的方法

2）bfprt算法
*/

// 改写快排，时间复杂度O(N)
// k >= 1
func minKth2(array []int, k int) int {
	arr := make([]int, len(array))
	copy(arr, array)
	return process2(arr, 0, len(arr)-1, k-1)
}

// arr 第k小的数
// process2(arr, 0, N-1, k-1)
// arr[L..R]  范围上，如果排序的话(不是真的去排序)，找位于index的数
// index [L..R]
func process2(arr []int, L, R, index int) int {
	if L == R { // L = =R ==INDEX
		return arr[L]
	}
	// 不止一个数  L +  [0, R -L]
	rand.Seed(time.Now().UnixNano())
	pivot := arr[L+rand.Intn(R-L)]
	ran := partition(arr, L, R, pivot)
	if index >= ran[0] && index <= ran[1] {
		return arr[index]
	} else if index < ran[0] {
		return process2(arr, L, ran[0]-1, index)
	} else {
		return process2(arr, ran[1]+1, R, index)
	}
}

// 利用bfprt算法，时间复杂度O(N)
func minKth3(array []int, k int) int {
	arr := make([]int, len(array))
	copy(arr, array)
	return bfprt(arr, 0, len(arr)-1, k-1)
}

// arr[L..R]  如果排序的话，位于index位置的数，是什么，返回
func bfprt(arr []int, L, R, index int) int {
	if L == R {
		return arr[L]
	}
	// L...R  每五个数一组
	// 每一个小组内部排好序
	// 小组的中位数组成新数组
	// 这个新数组的中位数返回
	pivot := medianOfMedians(arr, L, R)
	ran := partition(arr, L, R, pivot)
	if index >= ran[0] && index <= ran[1] {
		return arr[index]
	} else if index < ran[0] {
		return bfprt(arr, L, ran[0]-1, index)
	} else {
		return bfprt(arr, ran[1]+1, R, index)
	}
}

// arr[L...R]  五个数一组
// 每个小组内部排序
// 每个小组中位数领出来，组成marr
// marr中的中位数，返回
func medianOfMedians(arr []int, L, R int) int {
	size := R - L + 1
	offset := 1
	if size%5 == 0 {
		offset = 0
	}
	mArr := make([]int, size/5+offset)
	for team := 0; team < len(mArr); team++ {
		teamFirst := L + team*5
		// L ... L + 4
		// L +5 ... L +9
		// L +10....L+14
		mArr[team] = getMedian(arr, teamFirst, Min(R, teamFirst+4))
	}
	// marr中，找到中位数
	// marr(0, marr.len - 1,  mArr.length / 2 )
	return bfprt(mArr, 0, len(mArr)-1, len(mArr)/2)
}

func getMedian(arr []int, L, R int) int {
	insertionSort(arr, L, R)
	return arr[(L+R)/2]
}

func insertionSort(arr []int, L, R int) {
	for i := L + 1; i <= R; i++ {
		for j := i - 1; j >= L && arr[j] > arr[j+1]; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}

func partition(arr []int, L, R, pivot int) []int {
	less := L - 1
	more := R + 1
	cur := L
	for cur < more {
		if arr[cur] < pivot {
			arr[cur], arr[less+1] = arr[less+1], arr[cur]
			less++
			cur++
		} else if arr[cur] > pivot {
			arr[cur], arr[more-1] = arr[more-1], arr[cur]
			more--
		} else {
			cur++
		}
	}
	return []int{less + 1, more - 1}
}

func main() {
	arr := []int{4, 2, 6, 8, 2, 7, 3, 9, 6, 3, 1, 2}
	k := 4
	fmt.Println(minKth2(arr, k))
	fmt.Println(minKth3(arr, k))
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
