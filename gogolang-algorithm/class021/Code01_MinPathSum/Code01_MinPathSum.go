package main

import (
	"fmt"
	"gogolang-algorithm/utils"
)

/*
给定一个二维数组matrix，一个人必须从左上角出发，最后到达右下角
沿途只可以向下或者向右走，沿途的数字都累加就是距离累加和
返回最小距离累加和
*/

func minPathSum1(m [][]int) int {
	if m == nil || len(m) == 0 || m[0] == nil || len(m[0]) == 0 {
		return 0
	}
	row := len(m)
	col := len(m[0])
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}
	dp[0][0] = m[0][0]
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + m[i][0]
	}
	for j := 1; j < col; j++ {
		dp[0][j] = dp[0][j-1] + m[0][j]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = utils.Min(dp[i-1][j], dp[i][j-1]) + m[i][j]
		}
	}
	return dp[row-1][col-1]
}

func minPathSum2(m [][]int) int {
	if m == nil || len(m) == 0 || m[0] == nil || len(m[0]) == 0 {
		return 0
	}
	row := len(m)
	col := len(m[0])
	dp := make([]int, col)
	dp[0] = m[0][0]
	for j := 1; j < col; j++ { // 第 0 行的值
		dp[j] = dp[j-1] + m[0][j]
	}
	for i := 1; i < row; i++ { // 从第一行出发
		dp[0] += m[i][0]
		for j := 1; j < col; j++ {
			dp[j] = utils.Min(dp[j-1], dp[j]) + m[i][j]
		}
	}
	return dp[col-1]
}
func main() {
	m1 := [][]int{
		{1, 3, 5, 9},
		{8, 1, 3, 4},
		{5, 0, 6, 1},
		{8, 8, 4, 0},
	}

	fmt.Println(minPathSum1(m1))

	m2 := [][]int{{1, 3, 5, 9}, {8, 1, 3, 4}, {5, 0, 6, 1}, {8, 8, 4, 0}}
	fmt.Println(minPathSum2(m2))
}
