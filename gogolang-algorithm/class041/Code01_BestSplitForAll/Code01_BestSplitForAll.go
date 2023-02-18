package main

import "fmt"

/*
给定一个非负数组arr，长度为N，
那么有N-1种方案可以把arr切成左右两部分
每一种方案都有，min{左部分累加和，右部分累加和}
求这么多方案中，min{左部分累加和，右部分累加和}的最大值是多少？
整个过程要求时间复杂度O(N)
*/

func bestSplit2(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	sumAll := 0
	for index := range arr {
		sumAll += arr[index]
	}
	ans := 0
	sumL := 0
	for s := 0; s < len(arr)-1; s++ {
		sumL += arr[s]
		sumR := sumAll - sumL
		ans = max(ans, min(sumL, sumR))
	}
	return ans
}

/*
模型存在不回退
ans = 最优 {最差 {左(某指标)  右(某指标)}}
ans = 最差 {最优 {左(某指标)  右(某指标)}}

拓展情况：
指标和区间 存在单调性
eg：数组区间增大，前缀和变大
*/


func bestSplit3(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return nil
	}
	ans := make([]int, len(arr))
	ans[0] = 0
	// arr =   {5, 3, 1, 3}
	//          0  1  2  3
	// sum ={0, 5, 8, 9, 12}
	//       0  1  2  3   4
	// 0~2 ->  sum[3] - sum[0]
	// 1~3 ->  sum[4] - sum[1]
	sum := make([]int, len(arr)+1) //前缀和数组，长度要比arr多一个
	for i := 0; i < len(arr); i++ {
		sum[i+1] = sum[i] + arr[i]
	}
	best := 0 // 最优划分下标  0~range-1上，最优划分是左部分[0~best]  右部分[best+1~range-1]
	for rank := 0; rank < len(arr); rank++ {
		for best+1 < rank {
			before := min(sumF(sum, 0, best), sumF(sum, best+1, rank))
			after := min(sumF(sum, 0, best+1), sumF(sum, best+2, rank))
			if after >= before { // 注意，一定要是>=，只是>会出错
				best++
			} else {
				break
			}
		}
		ans[rank] = min(sumF(sum, 0, best), sumF(sum, best+1, rank))
	}
	return ans
}

// 求原来的数组arr中，arr[L...R]的累加和
func sumF(sum []int, L, R int) int {
	return sum[R+1] - sum[L]
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(bestSplit2(arr))

	res := bestSplit3(arr)
	fmt.Println(res)
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
