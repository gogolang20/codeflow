package main

/*
arr是货币数组，其中的值都是正数。再给定一个正数aim。
每个值都认为是一张货币，
认为值相同的货币没有任何不同，
返回组成aim的方法数
例如：arr = {1,2,1,1,2,1,2}，aim = 4
方法：1+1+1+1、1+1+2、2+2
一共就3种方法，所以返回3
*/

type Info struct {
	coins  []int
	zhangs []int
}

func getInfo(arr []int) *Info {
	counts := make(map[int]int)
	for _, value := range arr {
		if _, ok := counts[value]; ok {
			counts[value] = 1
		} else {
			counts[value] ++
		}
	}
	N := len(counts)
	coins := make([]int, N)
	zhangs := make([]int, N)
	index := 0
	for key, valye := range counts {
		coins[index] = key
		zhangs[index] = valye
		index++
	}
	return &Info{coins: coins, zhangs: zhangs}
}
func coinsWay(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}
	info := getInfo(arr)
	return process(info.coins, info.zhangs, 0, aim)
}

// coins 面值数组，正数且去重
// zhangs 每种面值对应的张数
func process(coins, zhangs []int, index, rest int) int {
	if index == len(coins) {
		if rest == 0 {
			return 1
		} else {
			return 0
		}
	}
	ways := 0
	for zhang := 0; zhang*coins[index] <= rest && zhang <= zhangs[index]; zhang++ {
		ways += process(coins, zhangs, index+1, rest-(zhang*coins[index]))
	}
	return ways
}

func dp1(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}
	info := getInfo(arr)
	coins := info.coins
	zhangs := info.zhangs
	N := len(coins)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	dp[N][0] = 1
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			ways := 0
			for zhang := 0; zhang*coins[index] <= rest && zhang <= zhangs[index]; zhang++ {
				ways += dp[index+1][rest-(zhang*coins[index])]
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
	info := getInfo(arr)
	coins := info.coins
	zhangs := info.zhangs
	N := len(coins)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	dp[N][0] = 1
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			dp[index][rest] = dp[index+1][rest]
			if rest-coins[index] >= 0 {
				dp[index][rest] += dp[index][rest-coins[index]]
			}
			if rest-coins[index]*(zhangs[index]+1) >= 0 {
				dp[index][rest] -= dp[index+1][rest-coins[index]*(zhangs[index]+1)]
			}
		}
	}
	return dp[0][aim]
}
