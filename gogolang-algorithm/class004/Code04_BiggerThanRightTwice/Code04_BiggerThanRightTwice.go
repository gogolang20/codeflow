package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
面试问题三
一个数组某个数num的右边数乘以2依旧小于num，求个数
*/
func biggerTwice(arr []int) int {
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
	//左数比右数大两倍
	ans := 0
	windowR := mid + 1 // 目前囊括进来的数 [mid + 1, windowR)
	for i := L; i <= mid; i++ {
		for windowR <= R && (arr[i] > (arr[windowR] << 1)) {
			windowR++
		}
		ans += windowR - mid - 1
	}
	temp := make([]int, R-L+1)
	i, j := L, mid+1
	index := 0
	for i <= mid && j <= R {
		if arr[i] < arr[j] {
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
	return ans // 返回结果
}

// for test
func comparator(arr []int) int {
	ans := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > (arr[j] << 1) {
				ans++
			}
		}
	}
	return ans
}

func main() {
	arrLen := 100
	arr := make([]int, arrLen)
	arr1 := make([]int, arrLen)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(1000)
	}
	copy(arr1, arr)

	fmt.Println(biggerTwice(arr))
	fmt.Println(comparator(arr1))
}
