package main

import "container/list"

/*
给定一个非负数组arr，代表直方图
返回直方图的最大长方形面积
*/

// 测试链接：https://leetcode.com/problems/largest-rectangle-in-histogram
// https://leetcode.cn/problems/largest-rectangle-in-histogram/

// 测试通过
func largestRectangleArea1(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	maxArea := 0
	stack := list.New()
	for i := 0; i < len(height); i++ {
		for stack.Len() > 0 && height[i] <= height[stack.Front().Value.(int)] {
			j := stack.Front().Value.(int)
			stack.Remove(stack.Front())
			k := 0
			if stack.Len() > 0 {
				k = stack.Front().Value.(int)
			} else {
				k = -1
			}
			curArea := (i - k - 1) * height[j]
			maxArea = Max(maxArea, curArea)
		}
		stack.PushFront(i)
	}
	for stack.Len() > 0 {
		j := stack.Front().Value.(int)
		stack.Remove(stack.Front())
		k := 0
		if stack.Len() > 0 {
			k = stack.Front().Value.(int)
		} else {
			k = -1
		}
		curArea := (len(height) - k - 1) * height[j]
		maxArea = Max(maxArea, curArea)
	}
	return maxArea
}

// 测试通过
func largestRectangleArea2(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	N := len(height)
	stack := make([]int, N)
	si := -1
	maxArea := 0
	for i := 0; i < len(height); i++ {
		for si != -1 && height[i] <= height[stack[si]] {
			j := stack[si]
			si--
			k := 0
			if si == -1 {
				k = -1
			} else {
				k = stack[si]
			}
			curArea := (i - k - 1) * height[j]
			maxArea = Max(maxArea, curArea)
		}
		stack[si+1] = i
		si++
	}

	for si != -1 {
		j := stack[si]
		si--
		k := 0
		if si == -1 {
			k = -1
		} else {
			k = stack[si]
		}
		curArea := (len(height) - k - 1) * height[j]
		maxArea = Max(maxArea, curArea)
	}
	return maxArea
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
