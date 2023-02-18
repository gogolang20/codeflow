package main

import (
	"fmt"
)

/*
给定一个正数数组arr，请把arr中所有的数分成两个集合
如果arr长度为偶数，两个集合包含数的个数要一样多
如果arr长度为奇数，两个集合包含数的个数必须只差一个
请尽量让两个集合的累加和接近
返回：
最接近的情况下，较小集合的累加和
*/

func right(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	if (len(arr) & 1) == 0 {
		return process(arr, 0, len(arr)/2, sum/2)
	} else {
		return max(process(arr, 0, len(arr)/2, sum/2), process(arr, 0, len(arr)/2+1, sum/2))
	}
}

// arr[i....]自由选择，挑选的个数一定要是picks个，累加和<=rest, 离rest最近的返回
func process(arr []int, i, picks, rest int) int {
	if i == len(arr) {
		if picks == 0 {
			return 0
		} else {
			return -1
		}
	} else {
		p1 := process(arr, i+1, picks, rest)
		// 就是要使用arr[i]这个数
		p2 := -1
		next := -1
		if arr[i] <= rest {
			next = process(arr, i+1, picks-1, rest-arr[i])
		}
		if next != -1 {
			p2 = arr[i] + next
		}
		return max(p1, p2)
	}
}

//课上没有讲dp1方法
//func dp1(arr []int) int {
//	if arr == nil || len(arr) < 2 {
//		return 0
//	}
//	sum := 0
//	for i := range arr {
//		sum += arr[i]
//	}
//	sum >>= 1
//	N := len(arr)
//	M := (len(arr) + 1) >> 1 // 向上取整
//	dp := make([][][]int, N)
//	for i := range dp {
//		dp[i] = make([][]int, M+1)
//		for j := range dp[i] {
//			dp[i][j] = make([]int, sum+1)
//		}
//	}
//	for i := 0; i < N; i++ {
//		for j := 0; j <= M; j++ {
//			for k := 0; k <= sum; k++ {
//				dp[i][j][k] = math.MinInt
//			}
//		}
//	}
//	for i := 0; i < N; i++ {
//		for k := 0; k <= sum; k++ {
//			dp[i][0][k] = 0
//		}
//	}
//	for k := 0; k <= sum; k++ {
//		if arr[0] <= k {
//			dp[0][1][k] = arr[0]
//		} else {
//			dp[0][1][k] = math.MinInt
//		}
//	}
//	for i := 1; i < N; i++ {
//		for j := 1; j <= min(i+1, M); j++ {
//			for k := 0; k <= sum; k++ {
//				dp[i][j][k] = dp[i-1][j][k]
//				if k-arr[i] >= 0 {
//					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j-1][k-arr[i]]+arr[i])
//				}
//			}
//		}
//	}
//	return max(dp[N-1][M][sum], dp[N-1][N-M][sum])
//}

func dp2(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	sum >>= 1
	N := len(arr)
	M := (len(arr) + 1) >> 1
	dp := make([][][]int, N+1)
	for i := range dp {
		dp[i] = make([][]int, M+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, sum+1)
		}
	}
	for i := 0; i <= N; i++ { //默认所有值是无效的
		for j := 0; j <= M; j++ {
			for k := 0; k <= sum; k++ {
				dp[i][j][k] = -1
			}
		}
	}
	for rest := 0; rest <= sum; rest++ {
		dp[N][0][rest] = 0
	}
	for i := N - 1; i >= 0; i-- {
		for picks := 0; picks <= M; picks++ {
			for rest := 0; rest <= sum; rest++ {
				p1 := dp[i+1][picks][rest]
				// 就是要使用arr[i]这个数
				p2 := -1
				next := -1
				if picks-1 >= 0 && arr[i] <= rest {
					next = dp[i+1][picks-1][rest-arr[i]]
				}
				if next != -1 {
					p2 = arr[i] + next
				}
				dp[i][picks][rest] = max(p1, p2)
			}
		}
	}
	if (len(arr) & 1) == 0 {
		return dp[0][len(arr)/2][sum]
	} else {
		return max(dp[0][len(arr)/2][sum], dp[0][(len(arr)/2)+1][sum])
	}
}

func main() {
	arr := []int{6, 3, 7, 9, 2, 6}
	res1 := right(arr)
	//res2 := dp1(arr)
	res3 := dp2(arr)
	if res1 == res3 {
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
