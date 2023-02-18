package main

import "math"

/*
一座大楼有 0~N 层，地面算作第 0 层，最高的一层为第 N 层。已知棋子从第 0 层掉落肯定 不会摔碎，从第 i 层掉落可能会摔碎，
	也可能不会摔碎(1≤i≤N)。给定整数 N 作为楼层数， 再给定整数 K 作为棋子数，返 回如果想找到棋子不会摔碎的最高层数，即使在最差的情况下扔 的最少次数。一次只能扔一个棋子。
【举例】
N=10，K=1。
返回 10。因为只有 1 棵棋子，所以不得不从第 1 层开始一直试到第 10 层，在最差的情况 下，即第 10 层 是不会摔坏的最高层，最少也要扔 10 次。
N=3，K=2。
返回 2。先在 2 层扔 1 棵棋子，如果碎了，试第 1 层，如果没碎，试第 3 层。 N=105，K=2 返回 14。
第一个棋子先在 14 层扔，碎了则用仅存的一个棋子试 1~13。 若没碎，第一个棋子继续在 27 层扔，碎了则 用仅存的一个棋子试 15~26。
	若没碎，第一个棋子继续在 39 层扔，碎了则用仅存的一个棋子试 28~38。
	若 没碎，第一个棋子继续在 50 层扔，碎了则用仅存的一个棋子试 40~49。
	若没碎，第一个棋子继续在 60 层扔， 碎了则用仅存的一个棋子试 51~59。
	若没碎，第一个棋子继续在 69 层扔，碎了则用仅存的一个棋子试 61~68。
	若没碎，第一个棋子继续在 77 层扔，碎了则用仅存的一个棋子试 70~76。
	若没碎，第一个棋子继续在 84 层 扔，碎了则用仅存的一个棋子试 78~83。
	若没碎，第一个棋子继续在 90 层扔，碎了则用仅存的一个棋子试 85~89。
	若没碎，第一个棋子继续在 95 层扔，碎了则用仅存的一个棋子试 91~94。
	若没碎，第一个棋子继续 在 99 层扔，碎了则用仅存的一个棋子试 96~98。
	若没碎，第一个棋子继续在 102 层扔，碎了则用仅存的一 个棋子试 100、101。
	若没碎，第一个棋子继续在 104 层扔，碎了则用仅存的一个棋子试 103。
	若没碎，第 一个棋子继续在 105 层扔，若到这一步还没碎，那么 105 便是结果。
*/

// leetcode测试链接：https://leetcode.com/problems/super-egg-drop
// https://leetcode.cn/problems/super-egg-drop/submissions/

// 四边形不等式：超出时间限制
func superEggDrop3(kChess, nLevel int) int {
	if nLevel < 1 || kChess < 1 {
		return 0
	}
	if kChess == 1 {
		return nLevel
	}
	dp := make([][]int, nLevel+1)
	for i := range dp {
		dp[i] = make([]int, kChess+1)
	}
	for i := 1; i != len(dp); i++ {
		dp[i][1] = i
	}
	best := make([][]int, nLevel+1)
	for i := range dp {
		best[i] = make([]int, kChess+1)
	}
	for i := 1; i != len(dp[0]); i++ {
		dp[1][i] = 1
		best[1][i] = 1
	}
	for i := 2; i < nLevel+1; i++ {
		for j := kChess; j > 1; j-- {
			ans := math.MaxInt
			bestChoose := -1
			down := best[i-1][j]
			up := i
			if j != kChess {
				up = best[i][j+1]
			}
			for first := down; first <= up; first++ {
				cur := Max(dp[first-1][j-1], dp[i-first][j])
				if cur <= ans {
					ans = cur
					bestChoose = first
				}
			}
			dp[i][j] = ans + 1
			best[i][j] = bestChoose
		}
	}
	return dp[nLevel][kChess]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
最优解思路：
	i 个棋子，扔 j 次，能解决几层楼 ？
*/

// 测试通过
func superEggDrop4(kChess, nLevel int) int {
	if nLevel < 1 || kChess < 1 {
		return 0
	}
	dp := make([]int, kChess)
	res := 0
	for {
		res++
		previous := 0
		for i := 0; i < len(dp); i++ {
			tmp := dp[i]
			dp[i] = dp[i] + previous + 1
			previous = tmp
			if dp[i] >= nLevel {
				return res
			}
		}
	}
}

// 测试通过
func superEggDrop5(kChess, nLevel int) int {
	if nLevel < 1 || kChess < 1 {
		return 0
	}
	bsTimes := log2N(nLevel) + 1
	if kChess >= bsTimes {
		return bsTimes
	}
	dp := make([]int, kChess)
	res := 0
	for {
		res++
		previous := 0
		for i := 0; i < len(dp); i++ {
			tmp := dp[i]
			dp[i] = dp[i] + previous + 1
			previous = tmp
			if dp[i] >= nLevel {
				return res
			}
		}
	}
}

func log2N(n int) int {
	res := -1
	for n != 0 {
		res++
		n >>= 1
	}
	return res
}
