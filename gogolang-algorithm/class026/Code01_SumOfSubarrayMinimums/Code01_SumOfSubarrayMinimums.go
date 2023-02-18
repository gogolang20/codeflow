package main

import "container/list"

/*
给定一个数组arr，
返回所有子数组最小值的累加和
*/

// 测试链接：https://leetcode.com/problems/sum-of-subarray-minimums/
// subArrayMinSum1是暴力解
// subArrayMinSum2是最优解的思路
// sumSubarrayMins是最优解思路下的单调栈优化
// Leetcode上只提交sumSubarrayMins方法，时间复杂度O(N)，可以直接通过

// https://leetcode.cn/problems/sum-of-subarray-minimums/submissions/
// 测试通过
func sumSubarrayMins(arr []int) int {
	left := nearLessEqualLeft(arr)
	right := nearLessRight(arr)
	ans := 0
	for i := 0; i < len(arr); i++ {
		start := i - left[i]
		end := right[i] - i
		ans += start * end * arr[i]
		ans %= 1000000007
	}
	return ans
}

func nearLessEqualLeft(arr []int) []int {
	N := len(arr)
	left := make([]int, N)

	stack := list.New()
	for i := N - 1; i >= 0; i-- {
		for stack.Len() > 0 && arr[i] <= arr[stack.Front().Value.(int)] {
			left[stack.Front().Value.(int)] = i
			stack.Remove(stack.Front())
		}
		stack.PushFront(i)
	}
	for stack.Len() > 0 {
		left[stack.Front().Value.(int)] = -1
		stack.Remove(stack.Front())
	}
	return left
}

func nearLessRight(arr []int) []int {
	N := len(arr)
	right := make([]int, N)
	stack := list.New()
	for i := 0; i < N; i++ {
		for stack.Len() > 0 && arr[stack.Front().Value.(int)] > arr[i] {
			right[stack.Front().Value.(int)] = i
			stack.Remove(stack.Front())
		}
		stack.PushFront(i)
	}
	for stack.Len() > 0 {
		right[stack.Front().Value.(int)] = N
		stack.Remove(stack.Front())
	}
	return right
}
