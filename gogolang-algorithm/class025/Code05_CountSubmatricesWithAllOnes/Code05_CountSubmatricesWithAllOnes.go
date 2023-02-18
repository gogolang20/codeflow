package main

/*
给定一个二维数组matrix，其中的值不是0就是1，
返回全部由1组成的子矩形数量
*/

// 测试链接：https://leetcode.com/problems/count-submatrices-with-all-ones
// https://leetcode.cn/problems/count-submatrices-with-all-ones/

// 测试通过
func numSubmat(mat [][]int) int {
	if mat == nil || len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}
	nums := 0
	height := make([]int, len(mat[0]))
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				height[j] = 0
			} else {
				height[j] = height[j] + 1
			}
		}
		nums += countFromBottom(height)
	}
	return nums
}

func countFromBottom(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	nums := 0
	stack := make([]int, len(height)) // 作为栈
	si := -1
	for i := 0; i < len(height); i++ {
		for si != -1 && height[stack[si]] >= height[i] {
			cur := stack[si]
			si--
			if height[cur] > height[i] {
				left := 0
				if si == -1 {
					left = -1
				} else {
					left = stack[si]
				}
				n := i - left - 1
				down := 0
				if left == -1 {
					down = Max(0, height[i])
				} else {
					down = Max(height[left], height[i])
				}
				nums += (height[cur] - down) * num(n)
			}
		}
		stack[si+1] = i
		si++
	}
	for si != -1 {
		cur := stack[si]
		si--
		left := 0
		if si == -1 {
			left = -1
		} else {
			left = stack[si]
		}
		n := len(height) - left - 1
		down := 0
		if left != -1 {
			down = height[left]
		}
		nums += (height[cur] - down) * num(n)
	}
	return nums
}

func num(n int) int {
	return ((n * (1 + n)) >> 1)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
