package main

import "sort"

/*
给定一个无序数组arr中，长度为N，给定一个正数k，返回top k个最大的数

不同时间复杂度三个方法：

1）O(N*logN)
2）O(N + K*logN)
3）O(n + k*logk)
*/

// 时间复杂度O(N*logN)
// 排序+收集
func maxTopK1(arr []int, k int) []int {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	N := len(arr)
	k = Min(N, k)
	sort.Ints(arr)
	ans := make([]int, k)
	for i, j := N-1, 0; j < k; i-- {
		ans[j] = arr[i]
		j++
	}
	return ans
}



func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
