package main

import "fmt"

/*
给定一个正数数组arr，
请把arr中所有的数分成两个集合，尽量让两个集合的累加和接近
返回：
最接近的情况下，较小集合的累加和
*/

func right(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	sum := 0
	for i := range arr {
		sum += arr[i] //数组累加和
	}
	return process(arr, 0, sum/2)
}

// arr[i...]可以自由选择，请返回累加和尽量接近rest，但不能超过rest的情况下，最接近的累加和是多少？
func process(arr []int, i, rest int) int {
	if i == len(arr) {
		return 0
	} else { // 还有数，arr[i]这个数
		p1 := process(arr, i+1, rest) // 可能性1，不使用arr[i]
		p2 := 0                       // 可能性2，要使用arr[i]
		if arr[i] <= rest {           // 加上 arr[i] 位置的数不会越界
			p2 = arr[i] + process(arr, i+1, rest-arr[i])
		}
		return max(p1, p2)
	}
}

func dps(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	sum /= 2
	N := len(arr)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, sum+1)
	}
	for i := N - 1; i >= 0; i-- {
		for rest := 0; rest <= sum; rest++ {
			p1 := dp[i+1][rest] // 可能性1，不使用arr[i]
			p2 := 0             // 可能性2，要使用arr[i]
			if arr[i] <= rest {
				p2 = arr[i] + dp[i+1][rest-arr[i]]
			}
			dp[i][rest] = max(p1, p2)
		}
	}
	return dp[0][sum]
}

func main() {
	arr := []int{5, 8, 3, 9, 2, 4}
	res1 := right(arr)
	res2 := dps(arr)
	if res1 == res2 {
		fmt.Println("测试成功")
	} else {
		fmt.Println("测试失败")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
