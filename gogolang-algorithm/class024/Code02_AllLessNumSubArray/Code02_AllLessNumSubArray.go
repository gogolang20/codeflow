package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

/*
给定一个整型数组arr，和一个整数num
某个arr中的子数组sub，如果想达标，必须满足：
sub中最大值 – sub中最小值 <= num，
返回arr中达标子数组的数量
*/

// 暴力的对数器方法
func right(arr []int, sum int) int {
	if arr == nil || len(arr) == 0 || sum < 0 {
		return 0
	}
	N := len(arr)
	count := 0
	for L := 0; L < N; L++ {
		for R := L; R < N; R++ {
			max := arr[L]
			min := arr[L]
			for i := L + 1; i <= R; i++ {
				max = Max(max, arr[i])
				min = Min(min, arr[i])
			}
			if max-min <= sum {
				count++
			}
		}
	}
	return count
}

func num(arr []int, sum int) int {
	if arr == nil || len(arr) == 0 || sum < 0 {
		return 0
	}
	N := len(arr)
	count := 0
	maxWindow := list.New()
	minWindow := list.New()
	for L, R := 0, 0; L < N; L++ {
		for R < N {
			for maxWindow.Len() > 0 && arr[maxWindow.Back().Value.(int)] <= arr[R] {
				maxWindow.Remove(maxWindow.Back())
			}
			maxWindow.PushBack(R)
			for minWindow.Len() > 0 && arr[minWindow.Back().Value.(int)] >= arr[R] {
				minWindow.Remove(minWindow.Back())
			}
			minWindow.PushBack(R)
			if arr[maxWindow.Front().Value.(int)]-arr[minWindow.Front().Value.(int)] > sum { //初次不达标了
				break
			} else {
				R++
			}
		}
		count += R - L
		if maxWindow.Front().Value.(int) == L {
			maxWindow.Remove(maxWindow.Front())
		}
		if minWindow.Front().Value.(int) == L {
			minWindow.Remove(minWindow.Front())
		}
	}
	return count
}

func main() {
	rand.Seed(time.Now().UnixNano())
	testTime := 10000
	maxLen := 100
	maxValue := 200
	arr1 := generateRandomArray(maxLen, maxValue)
	arr2 := make([]int, maxLen)
	copy(arr2, arr1)
	sum := rand.Intn(maxLen)

	for i := 0; i < testTime; i++ {
		res1 := right(arr2, sum)
		res2 := num(arr1, sum)
		if res1 != res2 {
			fmt.Println("Oops!!!")
			fmt.Println(res1)
			fmt.Println(res2)
			break
		}
	}
	fmt.Println("Finish")
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func generateRandomArray(maxLen, maxValue int) []int {
	res := make([]int, 0)
	for i := 0; i < maxLen; i++ {
		res = append(res, rand.Intn(maxValue))
	}
	return res
}
