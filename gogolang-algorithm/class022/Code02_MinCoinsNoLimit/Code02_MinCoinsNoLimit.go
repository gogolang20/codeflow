package main

import (
	"fmt"
	"math"
)

/*
arr是面值数组，其中的值都是正数且没有重复。再给定一个正数aim。
每个值都认为是一种面值，且认为张数是无限的。
返回组成aim的最少货币数
*/

func minCoins(arr []int, aim int) int {
	return process(arr, 0, aim)
}

// arr[index...]面值，每种面值张数自由选择，
// 搞出rest正好这么多钱，返回最小张数
// 拿Integer.MAX_VALUE标记怎么都搞定不了
func process(arr []int, index, rest int) int {
	if index == len(arr) {
		if rest == 0 {
			return 0
		} else {
			return math.MaxInt
		}
	} else {
		ans := math.MaxInt
		for zhang := 0; zhang*arr[index] <= rest; zhang++ {
			next := process(arr, index+1, rest-zhang*arr[index])
			if next != math.MaxInt {
				ans = min(ans, zhang+next)
			}
		}
		return ans
	}
}

func dp1(arr []int, aim int) int {
	if aim == 0 {
		return 0
	}
	N := len(arr)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	dp[N][0] = 0
	for j := 1; j <= aim; j++ {
		dp[N][j] = math.MaxInt
	}
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			ans := math.MaxInt
			for zhang := 0; zhang*arr[index] <= rest; zhang++ {
				next := dp[index+1][rest-zhang*arr[index]]
				if next != math.MaxInt {
					ans = min(ans, zhang+next)
				}
			}
			dp[index][rest] = ans
		}
	}
	return dp[0][aim]
}

func dp2(arr []int, aim int) int {
	if aim == 0 {
		return 0
	}
	N := len(arr)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	dp[N][0] = 0
	for j := 1; j <= aim; j++ {
		dp[N][j] = math.MaxInt
	}
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			dp[index][rest] = dp[index+1][rest]                                    //提前填上上面的值
			if rest-arr[index] >= 0 && dp[index][rest-arr[index]] != math.MaxInt { //斜率优化
				dp[index][rest] = min(dp[index][rest], dp[index][rest-arr[index]]+1)
			}
		}
	}
	return dp[0][aim]
}
func main() {
	arr := []int{1, 5, 10, 20}
	fmt.Println(minCoins(arr, 474))
	fmt.Println(dp1(arr, 474))
	fmt.Println(dp2(arr, 474))
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
