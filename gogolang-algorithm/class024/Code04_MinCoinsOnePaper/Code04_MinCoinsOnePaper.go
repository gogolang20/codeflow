package main

import (
	"container/list"
	"math"
)

/*
arr是货币数组，其中的值都是正数。再给定一个正数aim。
每个值都认为是一张货币，
返回组成aim的最少货币数
注意：
因为是求最少货币数，所以每一张货币认为是相同或者不同就不重要了
*/

func minCoins(arr []int, aim int) int {
	return process(arr, 0, aim)
}

func process(arr []int, index, rest int) int {
	if rest < 0 {
		return math.MaxInt
	}
	if index == len(arr) {
		if rest == 0 {
			return 0
		} else {
			return math.MaxInt
		}
	} else {
		p1 := process(arr, index+1, rest)
		p2 := process(arr, index+1, rest-arr[index])
		if p2 != math.MaxInt {
			p2++
		}
		return min(p1, p2)
	}
}

// dp1时间复杂度为：O(arr长度 * aim)
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
			p1 := dp[index+1][rest]
			p2 := math.MaxInt
			if rest-arr[index] >= 0 {
				p2 = dp[index+1][rest-arr[index]]
			}
			if p2 != math.MaxInt {
				p2++
			}
			dp[index][rest] = min(p1, p2)
		}
	}
	return dp[0][aim]
}

type Info struct {
	coins  []int
	zhangs []int
}

func NewInfo(c, z []int) *Info {
	return &Info{
		coins:  c,
		zhangs: z,
	}
}

func getInfo(arr []int) *Info {
	counts := make(map[int]int, 0)
	for _, value := range arr {
		if _, ok := counts[value]; !ok {
			counts[value] = 1
		} else {
			counts[value] += 1
		}
	}
	coins := make([]int, 0)
	zhangs := make([]int, 0)
	for key, value := range counts {
		coins = append(coins, key)
		zhangs = append(zhangs, value)
	}
	return NewInfo(coins, zhangs)
}

// dp3时间复杂度为：O(arr长度) + O(货币种数 * aim)
// 优化需要用到窗口内最小值的更新结构
func dp3(arr []int, aim int) int {
	if aim == 0 {
		return 0
	}
	// 得到info时间复杂度O(arr长度)
	info := getInfo(arr)
	c := info.coins
	z := info.zhangs
	N := len(c)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	dp[N][0] = 0
	for j := 1; j <= aim; j++ {
		dp[N][j] = math.MaxInt
	}
	//虽然是嵌套了很多循环，但是时间复杂度为O(货币种数 * aim)
	//因为用了窗口内最小值的更新结构
	for i := N - 1; i >= 0; i-- {
		for mod := 0; mod < min(aim+1, c[i]); mod++ {
			//当前面值 X
			//mod  mod + x   mod + 2*x   mod + 3 * x
			w := list.New()
			w.PushBack(mod)
			dp[i][mod] = dp[i+1][mod]
			for r := mod + c[i]; r <= aim; r += c[i] {
				for w.Len() > 0 && (dp[i+1][w.Back().Value.(int)] == math.MaxInt ||
					dp[i+1][w.Back().Value.(int)]+compensate(w.Back().Value.(int), r, c[i]) >= dp[i+1][r]) {
					w.Remove(w.Back())
				}
				w.PushBack(r)
				overdue := r - c[i]*(z[i]+1)
				if w.Front().Value.(int) == overdue {
					w.Remove(w.Front())
				}
				dp[i][r] = dp[i+1][w.Front().Value.(int)] + compensate(w.Front().Value.(int), r, c[i])
			}
		}
	}
	return dp[0][aim]
}

func compensate(pre, cur, coin int) int {
	return (cur - pre) / coin
}

func main() {



}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
