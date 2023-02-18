package main

/*
给定一个非负数组arr，和一个正数m。
返回arr的所有子序列中累加和%m之后的最大值。
*/

func max1(arr []int, m int) int {
	set := make(map[int]struct{})
	process(arr, 0, 0, set)
	max := 0
	for sum := range set {
		max = Max(max, sum%m)
	}
	return max
}

func process(arr []int, index, sum int, set map[int]struct{}) {
	if index == len(arr) {
		set[sum] = struct{}{}
	} else {
		process(arr, index+1, sum, set)
		process(arr, index+1, sum+arr[index], set)
	}
}

func max2(arr []int, m int) int {
	sum := 0
	N := len(arr)
	for i := 0; i < N; i++ {
		sum += arr[i]
	}
	dp := make([][]bool, N)
	//for i := range dp {
	//	dp[i] = make([]bool, sum+1)
	//}
	for i := 0; i < N; i++ {
		res := make([]bool,sum+1)
		dp[i] = res
	}
	for i := 0; i < N; i++ {
		dp[i][0] = true
	}
	dp[0][arr[0]] = true
	for i := 1; i < N; i++ {
		for j := 1; j <= sum; j++ {
			dp[i][j] = dp[i-1][j]
			if j-arr[i] >= 0 {
				dp[i][j] |= dp[i-1][j-arr[i]]
			}
		}
	}
	ans := 0
	for j := 0; j <= sum; j++ {
		if dp[N-1][j] {
			ans = Max(ans, j%m)
		}
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
