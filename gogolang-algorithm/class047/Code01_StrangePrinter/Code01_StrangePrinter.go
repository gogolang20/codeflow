package main

import "gogolang-algorithm/utils"

// https://leetcode.cn/problems/strange-printer/

func strangePrinter1(s string) int {
	if s == "" {
		return 0
	}
	str := []byte(s)
	return process1(str, 0, len(str)-1)
}

// 要想刷出str[L...R]的样子！
// 返回最少的转数
func process1(str []byte, L, R int) int {
	if L == R {
		return 1
	}
	// L...R
	ans := R - L + 1
	for k := L + 1; k <= R; k++ {
		// L...k-1 k....R
		temp := 0
		if str[L] == str[k] {
			temp = 1
		}
		ans = utils.Min(ans, process1(str, L, k-1)+process1(str, k, R)-temp)
	}
	return ans
}

// 测试通过
// 时间复杂度 O(N^3) 最优解，没有四边形不等式技巧
func strangePrinter3(s string) int {
	if s == "" {
		return 0
	}
	str := []byte(s)
	N := len(str)
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, N)
	}
	dp[N-1][N-1] = 1
	for i := 0; i < N-1; i++ {
		dp[i][i] = 1
		if str[i] == str[i+1] {
			dp[i][i+1] = 1
		} else {
			dp[i][i+1] = 2
		}
	}
	for L := N - 3; L >= 0; L-- {
		for R := L + 2; R < N; R++ {
			dp[L][R] = R - L + 1
			for k := L + 1; k <= R; k++ {
				temp := 0
				if str[L] == str[k] {
					temp = 1
				}
				dp[L][R] = utils.Min(dp[L][R], dp[L][k-1]+dp[k][R]-temp)
			}
		}
	}
	return dp[0][N-1]
}
