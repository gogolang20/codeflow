package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
面试问题二
左边数和右边数是降序
称为逆序对，求逆序对的个数？？？
从大到小开始merge，相等先拷贝右边
*/
func reverPairNumber(arr []int) int {
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
	p1 := mid
	p2 := R
	index := len(temp) - 1
	res := 0 // 逆序对的和
	for p1 >= L && p2 > mid {
		if arr[p1] > arr[p2] {
			res += p2 - mid // 累加结果
			temp[index] = arr[p1]
			p1--
		} else {
			temp[index] = arr[p2]
			p2--
		}
		index--
	}
	copy(temp[:index+1], arr[L:p1+1])
	copy(temp[:index+1], arr[mid+1:p2+1])
	copy(arr[L:R+1], temp[:])
	return res // 返回个数
}

// for test
func comparator(arr []int) int {
	ans := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				ans++
			}
		}
	}
	return ans
}

func main() {
	maxSize := 200
	arr := make([]int, maxSize)
	arr1 := make([]int, maxSize)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(1000)
	}
	copy(arr1, arr)

	fmt.Println(reverPairNumber(arr))
	fmt.Println(comparator(arr1))
}
