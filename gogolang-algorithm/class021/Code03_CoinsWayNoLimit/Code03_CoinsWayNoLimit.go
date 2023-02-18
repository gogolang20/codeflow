package main

/*
arr是面值数组，其中的值都是正数且没有重复。再给定一个正数aim。
每个值都认为是一种面值，且认为张数是无限的。
返回组成aim的方法数
例如：arr = {1,2}，aim = 4
方法如下：1+1+1+1、1+1+2、2+2
一共就3种方法，所以返回3
*/

func coinsWay(arr []int, aim int) int {

	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}
	return process(arr, 0, aim)
}

// arr[index....] 所有的面值，每一个面值都可以任意选择张数，组成正好rest这么多钱，方法数多少？
func process(arr []int, index, rest int) int {
	if index == len(arr) { // 没钱了
		if rest == 0 {
			return 1
		} else {
			return 0
		}
	}
	ways := 0
	for zhang := 0; zhang*arr[index] <= rest; zhang++ {
		ways += process(arr, index+1, rest-(zhang*arr[index]))
	}
	return ways
}

func dp1(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}
	N := len(arr)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	dp[N][0] = 1
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			ways := 0
			for zhang := 0; zhang*arr[index] <= rest; zhang++ {
				ways += dp[index+1][rest-(zhang*arr[index])]
			}
			dp[index][rest] = ways
		}
	}
	return dp[0][aim]
}

func dp2(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}
	N := len(arr)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	dp[N][0] = 1
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			dp[index][rest] = dp[index+1][rest]
			if rest-arr[index] >= 0 {
				dp[index][rest] += dp[index][rest-arr[index]]
			}
		}
	}
	return dp[0][aim]
}
