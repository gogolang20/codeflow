package main

import "math"

/*
如果一个字符相邻的位置没有相同字符，那么这个位置的字符出现不能被消掉。比如:"ab"，其中a和b都不能被消掉
如果一个字符相邻的位置有相同字符，就可以一起消掉。比如:“abbbc”，中间一串的b是可以被消掉的，
消除之后剩下“ac”。某些字符如果消掉了，剩下的字符认为重新靠在一起
给定一个字符串，你可以决定每一步消除的顺序，目标是请尽可能多的消掉字符，返回最少的剩余字符数量
比如："aacca", 如果先消掉最左侧的"aa"，那么将剩下"cca"，然后把"cc"消掉，剩下的"a"将无法再消除，返回1
但是如果先消掉中间的"cc"，那么将剩下"aaa"，最后都消掉就一个字符也不剩了，返回0，这才是最优解。
再比如："baaccabb"，如果先消除最左侧的两个a，剩下"bccabb"，如果再消除最左侧的两个c，剩下"babb"，
最后消除最右侧的两个b，剩下"ba"无法再消除，返回2
而最优策略是：先消除中间的两个c，剩下"baaabb"，再消除中间的三个a，剩下"bbb"，最后消除三个b，
不留下任何字符，返回0，这才是最优解
*/

// 优良尝试的动态规划版本
func restMin3(s string) int {
	if s == "" {
		return 0
	}
	if len(s) < 2 {
		return len(s)
	}
	str := []byte(s)
	N := len(str)
	dp := make([][][]int, N)
	for i := range dp {
		dp[i] = make([][]int, N)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < 2; k++ {
				dp[i][j][k] = -1
			}
		}
	}
	return dpProcess(str, 0, N-1, false, dp)
}

func dpProcess(str []byte, L, R int, has bool, dp [][][]int) int {
	if L > R {
		return 0
	}
	K := 0
	if has {
		K = 1
	}
	if dp[L][R][K] != -1 {
		return dp[L][R][K]
	}
	ans := 0
	if L == R {
		if K == 0 {
			ans = 1
		}
	} else {
		index := L
		all := K
		for index <= R && str[index] == str[L] {
			all++
			index++
		}
		way1 := 0
		if all > 1 {
			way1 = dpProcess(str, index, R, false, dp)
		} else {
			way1 = 1 + dpProcess(str, index, R, false, dp)
		}
		way2 := math.MaxInt
		for split := index; split <= R; split++ {
			if str[split] == str[L] && str[split] != str[split-1] {
				if dpProcess(str, index, split-1, false, dp) == 0 {
					way2 = Min(way2, dpProcess(str, split, R, all > 0, dp))
				}
			}
		}
		ans = Min(way1, way2)
	}
	dp[L][R][K] = ans
	return ans
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
