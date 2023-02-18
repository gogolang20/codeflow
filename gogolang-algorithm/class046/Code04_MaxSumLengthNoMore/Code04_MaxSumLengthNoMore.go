package main

import (
	"container/list"
	"math"
)

/*
给定一个数组arr，和一个正数M
返回在子数组长度不大于M的情况下，最大的子数组累加和
*/

// O(N^2)的解法，暴力解，用作对数器
func test(arr []int, M int) int {
	if arr == nil || len(arr) == 0 || M < 1 {
		return 0
	}
	N := len(arr)
	max := math.MaxInt
	for L := 0; L < N; L++ {
		sum := 0
		for R := L; R < N; R++ {
			if R-L+1 > M {
				break
			}
			sum += arr[R]
			max = Max(max, sum)
		}
	}
	return max
}

// O(N)的解法，最优解
func maxSum(arr []int, M int) int {
	if arr == nil || len(arr) == 0 || M < 1 {
		return 0
	}
	N := len(arr)
	sum := make([]int, N)
	sum[0] = arr[0]
	for i := 1; i < N; i++ {
		sum[i] = sum[i-1] + arr[i]
	}
	qmax := list.New()
	i := 0
	end := Min(N, M)
	for ; i < end; i++ {
		for qmax.Len() > 0 && sum[qmax.Back().Value.(int)] <= sum[i] {
			qmax.Remove(qmax.Back())
		}
		qmax.PushBack(i)
	}
	max := sum[qmax.Front().Value.(int)]
	L := 0
	for ; i < N; L++ {
		if qmax.Front().Value.(int) == L {
			qmax.Remove(qmax.Front())
		}
		for qmax.Len() > 0 && sum[qmax.Back().Value.(int)] <= sum[i] {
			qmax.Remove(qmax.Back())
		}
		qmax.PushBack(i)
		max = Max(max, sum[qmax.Front().Value.(int)]-sum[L])
		i++
	}
	for ; L < N-1; L++ {
		if qmax.Front().Value.(int) == L {
			qmax.Remove(qmax.Front())
		}
		max = Max(max, sum[qmax.Front().Value.(int)]-sum[L])
	}
	return max
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
