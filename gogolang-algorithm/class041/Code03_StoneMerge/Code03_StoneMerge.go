package main

import "math"

/*
摆放着n堆石子。现要将石子有次序地合并成一堆
规定每次只能选相邻的2堆石子合并成新的一堆，
并将新的一堆石子数记为该次合并的得分
求出将n堆石子合并成一堆的最小得分（或最大得分）合并方案
*/

func min2(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	N := len(arr)
	s := sum(arr)
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, N)
	}
	// dp[i][i] = 0
	for L := N - 2; L >= 0; L-- {
		for R := L + 1; R < N; R++ {
			next := math.MaxInt
			// dp(L..leftEnd)  + dp[leftEnd+1...R]  + 累加和[L...R]
			for leftEnd := L; leftEnd < R; leftEnd++ {
				next = Min(next, dp[L][leftEnd]+dp[leftEnd+1][R])
			}
			dp[L][R] = next + w(s, L, R)
		}
	}
	return dp[0][N-1]
}

func min3(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	N := len(arr)
	s := sum(arr)
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, N)
	}
	best := make([][]int, N)
	for i := range dp {
		best[i] = make([]int, N)
	}
	for i := 0; i < N-1; i++ {
		best[i][i+1] = i
		dp[i][i+1] = w(s, i, i+1)
	}
	for L := N - 3; L >= 0; L-- {
		for R := L + 2; R < N; R++ {
			next := math.MaxInt
			choose := -1
			for leftEnd := best[L][R-1]; leftEnd <= best[L+1][R]; leftEnd++ {
				cur := dp[L][leftEnd] + dp[leftEnd+1][R]
				if cur <= next {
					next = cur
					choose = leftEnd
				}
			}
			best[L][R] = choose
			dp[L][R] = next + w(s, L, R)
		}
	}
	return dp[0][N-1]
}

func w(s []int, l, r int) int {
	return s[r+1] - s[l]
}

func sum(arr []int) []int {
	N := len(arr)
	s := make([]int, N+1)
	s[0] = 0
	for i := 0; i < N; i++ {
		s[i+1] = s[i] + arr[i]
	}
	return s
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
