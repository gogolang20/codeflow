package main

import "math"

type Node struct {
	value int
	Left  *Node
	Right *Node
}

/*
给定一棵二叉树的头节点head，
返回这颗二叉树中最大的二叉搜索子树的大小
*/

type Info struct {
	maxBSTSubtreeSize int // 整个节点 最大子树node 数量
	allSize           int // 整个节点 node 数量
	max               int
	min               int
}

func NewInfo(m, a, max, min int) *Info {
	return &Info{
		maxBSTSubtreeSize: m,
		allSize:           a,
		max:               max,
		min:               min,
	}
}

func maxSubBSTSize2(head *Node) int {
	if head == nil {
		return 0
	}
	return process(head).maxBSTSubtreeSize
}
func process(x *Node) *Info {
	if x == nil {
		return nil // 无法确定 max min，所有返回 nil
	}
	leftInfo := process(x.Left)
	rightInfo := process(x.Right)
	max, min := x.value, x.value
	allSize := 1
	if leftInfo != nil {
		max = int(math.Max(float64(leftInfo.max), float64(max)))
		min = int(math.Min(float64(leftInfo.min), float64(min)))
		allSize += leftInfo.allSize
	}
	if rightInfo != nil {
		max = int(math.Max(float64(rightInfo.max), float64(max)))
		min = int(math.Min(float64(rightInfo.min), float64(min)))
		allSize += rightInfo.allSize
	}
	// 求 maxBSTSubtreeSize
	p1, p2, p3 := -1, -1, -1
	if leftInfo != nil {
		p1 = leftInfo.maxBSTSubtreeSize
	}
	if rightInfo != nil {
		p2 = rightInfo.maxBSTSubtreeSize
	}
	var leftBST, rightBST bool // 左右子树是否都为搜索二叉树
	if leftInfo == nil {
		leftBST = true
	} else {
		leftBST = leftInfo.maxBSTSubtreeSize == leftInfo.allSize
	}
	if rightInfo == nil {
		rightBST = true
	} else {
		rightBST = rightInfo.maxBSTSubtreeSize == rightInfo.allSize
	}
	if leftBST && rightBST { // 左右都为搜索二叉树，p3可以修改值
		var leftMaxLessX, rightMinMoreX bool // 判断左右子树有没有值
		if leftInfo == nil {
			leftMaxLessX = true
		} else {
			leftMaxLessX = leftInfo.max < x.value
		}
		if rightInfo == nil {
			rightMinMoreX = true
		} else {
			rightMinMoreX = x.value < rightInfo.min
		}
		if leftMaxLessX && rightMinMoreX {
			var leftSize, rightSize int // 左右子树的size
			if leftInfo == nil {
				leftSize = 0
			} else {
				leftSize = leftInfo.allSize
			}
			if rightInfo == nil {
				rightSize = 0
			} else {
				rightSize = rightInfo.allSize
			}
			p3 = leftSize + rightSize + 1
		}
	}
	maxBSTSubtreeSize := int(math.Max(math.Max(float64(p1), float64(p2)), float64(p3)))
	return NewInfo(maxBSTSubtreeSize, allSize, max, min)
}
