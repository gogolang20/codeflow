package main

import "fmt"

// 整型数组arr长度为n(3 <= n <= 10^4)，最初每个数字是<=200的正数且满足如下条件：
// 1. 0位置的要求：arr[0]<=arr[1]
// 2. n-1位置的要求：arr[n-1]<=arr[n-2]
// 3. 中间i位置的要求：arr[i]<=max(arr[i-1],arr[i+1])
// 但是在arr有些数字丢失了，比如k位置的数字之前是正数，丢失之后k位置的数字为0
// 请你根据上述条件，计算可能有多少种不同的arr可以满足以上条件
// 比如 [6,0,9] 只有还原成 [6,9,9]满足全部三个条件，所以返回1种，即[6,9,9]达标
func ways0(arr []int) int {
	return process0(arr, 0)
}
func process0(arr []int, index int) int {
	if index == len(arr) {
		if isValid(arr) {
			return 1
		} else {
			return 0
		}
	} else {
		if arr[index] != 0 {
			return process0(arr, index+1)
		} else {
			ways := 0
			for v := 1; v < 201; v++ {
				arr[index] = v
				ways += process0(arr, index+1)
			}
			arr[index] = 0
			return ways
		}
	}
}

func isValid(arr []int) bool {
	if arr[0] > arr[1] {
		return false
	}
	if arr[len(arr)-1] > arr[len(arr)-2] {
		return false
	}
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > Max(arr[i-1], arr[i+1]) {
			return false
		}
	}
	return true
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ways3(arr []int) int {
	N := len(arr)
	dp := make([][][]int, N)
	for i := range dp {
		dp[i] = make([][]int, 201)
		for j := range dp[i] {
			dp[i][j] = make([]int, 3)
		}
	}
	if arr[0] != 0 {
		dp[0][arr[0]][0] = 1
		dp[0][arr[0]][1] = 1
	} else {
		for v := 1; v < 201; v++ {
			dp[0][v][0] = 1
			dp[0][v][1] = 1
		}
	}
	presum := make([][]int, 201)
	for i := range presum {
		presum[i] = make([]int, 3)
	}

	for v := 1; v < 201; v++ {
		for s := 0; s < 3; s++ {
			presum[v][s] = presum[v-1][s] + dp[0][v][s]
		}
	}
	for i := 1; i < N; i++ {
		for v := 1; v < 201; v++ {
			for s := 0; s < 3; s++ {
				if arr[i] == 0 || v == arr[i] {
					if s == 0 || s == 1 {
						dp[i][v][s] += sum(1, v-1, 0, presum)
					}
					dp[i][v][s] += dp[i-1][v][1]
					dp[i][v][s] += sum(v+1, 200, 2, presum)
				}
			}
		}
		for v := 1; v < 201; v++ {
			for s := 0; s < 3; s++ {
				presum[v][s] = presum[v-1][s] + dp[i][v][s]
			}
		}
	}
	if arr[N-1] != 0 {
		return dp[N-1][arr[N-1]][2]
	} else {
		return sum(1, 200, 2, presum)
	}
}

func sum(begin, end, relation int, presum [][]int) int {
	return presum[end][relation] - presum[begin-1][relation]
}

func main() {
	arr := []int{1, 3, 5, 6, 8, 0}
	fmt.Println(ways0(arr))
	fmt.Println(ways3(arr))
}
