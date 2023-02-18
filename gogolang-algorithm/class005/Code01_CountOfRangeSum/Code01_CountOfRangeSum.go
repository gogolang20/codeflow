package main

/*
LeetCode 327题

https://leetcode.cn/problems/count-of-range-sum/

时间复杂度 O(N*logN)
*/

/*
1 求数组前缀和
2 原 [lower,upper] --> 前缀和求 [x-upper, x-lower]
3 前缀和数组 中之前有多少数落在 [x-upper, x-lower] 上
4 归并排序 中求
5 merge 左右不回退
*/

// 测试通过
func countRangeSum(arr []int, lower, upper int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	presumArr := make([]int, len(arr)) // 前缀和数组
	presumArr[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		presumArr[i] = presumArr[i-1] + arr[i]
	}
	return process(presumArr, 0, len(arr)-1, lower, upper)
}

func process(presumArr []int, L, R, lower, upper int) int {
	if L == R { // 代表原始数组 0...L 的前缀和
		if presumArr[L] >= lower && presumArr[L] <= upper {
			return 1
		} else {
			return 0
		}
	}
	mid := L + ((R - L) >> 1)
	return process(presumArr, L, mid, lower, upper) +
		process(presumArr, mid+1, R, lower, upper) +
		merge(presumArr, L, mid, R, lower, upper)
}

func merge(presumArr []int, L, mid, R, lower, upper int) int {
	ans := 0
	windowL := L // 滑动不回退
	windowR := L // 滑动不回退
	for i := mid + 1; i <= R; i++ {
		min := presumArr[i] - upper
		max := presumArr[i] - lower
		for windowR <= mid && presumArr[windowR] <= max {
			windowR++
		}
		for windowL <= mid && presumArr[windowL] < min {
			windowL++
		}
		ans += windowR - windowL
	}
	// 经典 merge 过程
	temp := make([]int, R-L+1)
	i, j := L, mid+1
	index := 0
	for i <= mid && j <= R {
		if presumArr[i] <= presumArr[j] {
			temp[index] = presumArr[i]
			i++
		} else {
			temp[index] = presumArr[j]
			j++
		}
		index++
	}
	copy(temp[index:], presumArr[i:mid+1])
	copy(temp[index:], presumArr[j:R+1])
	copy(presumArr[L:R+1], temp[:])
	return ans // 返回结果
}
