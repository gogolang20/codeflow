package main

// https://leetcode.cn/problems/remove-boxes/

// 测试通过
func removeBoxes2(boxes []int) int {
	N := len(boxes)
	dp := make([][][]int, N)
	for i := range dp {
		dp[i] = make([][]int, N)
		for j := range dp[i] {
			dp[i][j] = make([]int, N)
		}
	}
	ans := process2(boxes, 0, N-1, 0, dp)
	return ans
}

func process2(boxes []int, L, R, K int, dp [][][]int) int {
	if L > R {
		return 0
	}
	if dp[L][R][K] > 0 {
		return dp[L][R][K]
	}
	// 找到开头，
	// 1,1,1,1,1,5
	// 3 4 5 6 7 8
	//         !
	last := L
	for last+1 <= R && boxes[last+1] == boxes[L] {
		last++
	}
	// K个1     (K + last - L) last
	pre := K + last - L
	ans := (pre+1)*(pre+1) + process2(boxes, last+1, R, 0, dp)
	for i := last + 2; i <= R; i++ {
		if boxes[i] == boxes[L] && boxes[i-1] != boxes[L] {
			ans = Max(ans, process2(boxes, last+1, i-1, 0, dp)+process2(boxes, i, R, pre+1, dp))
		}
	}
	dp[L][R][K] = ans
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
