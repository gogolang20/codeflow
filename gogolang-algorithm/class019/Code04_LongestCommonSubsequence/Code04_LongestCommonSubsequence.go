package main

/*
给定两个字符串str1和str2，
返回这两个字符串的最长公共子序列长度

比如 ： str1 = “a12b3c456d”,str2 = “1ef23ghi4j56k”
最长公共子序列是“123456”，所以返回长度6
*/

// 这个问题leetcode上可以直接测
// 链接：https://leetcode.com/problems/longest-common-subsequence/
func longestCommonSubsequence1(s1, s2 string) int {
	if s1 == "" || s2 == "" {
		return 0
	}
	str1 := []byte(s1)
	str2 := []byte(s2)
	return process1(str1, str2, len(str1)-1, len(str2)-1)
}

func process1(str1, str2 []byte, i, j int) int {
	if i == 0 && j == 0 {
		if str1[i] == str2[j] {
			return 1
		} else {
			return 0
		}
	} else if i == 0 {
		if str1[i] == str2[j] {
			return 1
		} else {
			return process1(str1, str2, i, j-1)
		}
	} else if j == 0 {
		if str1[i] == str2[j] {
			return 1
		} else {
			return process1(str1, str2, i-1, j)
		}
	} else { // i != 0 && j != 0
		p1 := process1(str1, str2, i-1, j)
		p2 := process1(str1, str2, i, j-1)
		p3 := 0
		if str1[i] == str2[j] {
			p3 = 1 + process1(str1, str2, i-1, j-1)
		}
		return max(p1, max(p2, p3))
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
方法二
通过测试！！！
*/
func longestCommonSubsequence2(s1, s2 string) int {
	if s1 == "" || s2 == "" {
		return 0
	}
	str1 := []byte(s1)
	str2 := []byte(s2)
	N, M := len(str1), len(str2)
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, M)
	}
	if str1[0] == str2[0] {
		dp[0][0] = 1
	} else {
		dp[0][0] = 0
	}
	for j := 1; j < M; j++ { //第0行
		if str1[0] == str2[j] {
			dp[0][j] = 1
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < N; i++ { //第0列
		if str1[i] == str2[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 1; i < N; i++ {
		for j := 1; j < M; j++ {
			p1 := dp[i-1][j]
			p2 := dp[i][j-1]
			p3 := 0
			if str1[i] == str2[j] {
				p3 = 1 + dp[i-1][j-1]
			}
			dp[i][j] = max(p1, max(p2, p3))
		}
	}
	return dp[N-1][M-1]
}
