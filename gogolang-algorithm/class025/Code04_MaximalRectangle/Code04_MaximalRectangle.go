package main

import "container/list"

/*
给定一个二维数组matrix，其中的值不是0就是1，
返回全部由1组成的最大子矩形，内部有多少个1
*/

// 测试链接：https://leetcode.com/problems/maximal-rectangle/
// https://leetcode.cn/problems/maximal-rectangle/

// 通过测试  时间复杂度 O(n^2)
func maximalRectangle(maps [][]byte) int {
	if maps == nil || len(maps) == 0 || len(maps[0]) == 0 {
		return 0
	}
	maxArea := 0
	height := make([]int, len(maps[0]))
	for i := 0; i < len(maps); i++ {
		for j := 0; j < len(maps[0]); j++ {
			if maps[i][j] == '0' {
				height[j] = 0
			} else {
				height[j] = height[j] + 1
			}
		}
		maxArea = Max(maxRecFromBottom(height), maxArea)
	}
	return maxArea
}

// height是正方图数组
func maxRecFromBottom(height []int) int {
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

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
