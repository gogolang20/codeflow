package main

import (
	"container/list"
	"math"
)

/*
给定一个只包含正数的数组arr，arr中任何一个子数组sub，
一定都可以算出(sub累加和 )* (sub中的最小值)是什么，
那么所有子数组中，这个值最大是多少？
*/

/*、
思路：
以自己做最小值 来选择子数组
找到左边最近一个比自己小，和右边最近一个比自己小。两个下标
*/
func max1(arr []int) int {
	max := math.MinInt
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			minNum := math.MaxInt
			sum := 0
			for k := i; k <= j; k++ {
				sum += arr[k]
				minNum = Min(minNum, arr[k])
			}
			max = Max(max, minNum*sum)
		}
	}
	return max
}

// 单调栈方法
func max2(arr []int) int {
	size := len(arr)
	sums := make([]int, size)
	sums[0] = arr[0]
	for i := 1; i < size; i++ {
		sums[i] = sums[i-1] + arr[i]
	}
	// 使用没有重复值的单调栈
	max := math.MinInt
	stack := list.New()
	for i := 0; i < size; i++ {
		for stack.Len() > 0 && arr[stack.Front().Value.(int)] >= arr[i] {
			j := stack.Front().Value.(int)
			stack.Remove(stack.Front())
			if stack.Len() == 0 {
				max = Max(max, sums[i-1])
			} else {
				max = Max(max, (sums[i-1]-sums[stack.Front().Value.(int)])*arr[j])
			}
		}
		stack.PushFront(i)
	}
	for stack.Len() > 0 {
		j := stack.Front().Value.(int)
		stack.Remove(stack.Front())
		if stack.Len() == 0 {
			max = Max(max, sums[size-1])
		} else {
			max = Max(max, (sums[size-1]-sums[stack.Front().Value.(int)])*arr[j])
		}
	}
	return max
}

func main() {

}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
