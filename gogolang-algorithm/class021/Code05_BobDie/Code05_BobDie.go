package main

import "math"

/*
给定5个参数，N，M，row，col，k
表示在N*M的区域上，醉汉Bob初始在(row,col)位置
Bob一共要迈出k步，且每步都会等概率向上下左右四个方向走一个单位
任何时候Bob只要离开N*M的区域，就直接死亡
返回k步之后，Bob还在N*M的区域的概率
*/
func livePosibility1(row, col, k, N, M int) float64 {
	return float64(process(row, col, k, N, M)) / math.Pow(float64(4), float64(k))
}

// 目前在row，col位置，还有rest步要走，走完了如果还在棋盘中就获得1个生存点，返回总的生存点数
func process(row, col, rest, N, M int) int {
	if row < 0 || row == N || col < 0 || col == M {
		return 0
	}
	// 还在棋盘中！
	if rest == 0 {
		return 1
	}
	// 还在棋盘中！还有步数要走
	up := process(row-1, col, rest-1, N, M)
	down := process(row+1, col, rest-1, N, M)
	left := process(row, col-1, rest-1, N, M)
	right := process(row, col+1, rest-1, N, M)
	return up + down + left + right
}

func livePosibility2(row, col, k, N, M int) float64 {
	dp := make([][][]int, N)
	for i := range dp {
		dp[i] = make([][]int, M)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			dp[i][j][0] = 1
		}
	}
	for rest := 1; rest <= k; rest++ {
		for r := 0; r < N; r++ {
			for c := 0; c < M; c++ {
				dp[r][c][rest] = pick(dp, N, M, r-1, c, rest-1)
				dp[r][c][rest] += pick(dp, N, M, r+1, c, rest-1)
				dp[r][c][rest] += pick(dp, N, M, r, c-1, rest-1)
				dp[r][c][rest] += pick(dp, N, M, r, c+1, rest-1)
			}
		}
	}
	return float64(dp[row][col][k]) / math.Pow(float64(4), float64(k))
}

func pick(dp [][][]int, N, M, r, c, rest int) int {
	if r < 0 || r == N || c < 0 || c == M {
		return 0
	}
	return dp[r][c][rest]
}
