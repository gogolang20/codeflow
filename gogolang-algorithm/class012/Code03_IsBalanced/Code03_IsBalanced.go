package main

import "math"

type Node struct {
	value int
	Left  *Node
	Right *Node
}

type Info struct {
	isBalanced bool
	height     int
}

func NewInfo(is bool, h int) *Info {
	return &Info{
		isBalanced: is,
		height:     h,
	}
}

// 递归
func isBalanced1(head *Node) bool {
	return process(head).isBalanced
}

func process(x *Node) *Info {
	if x == nil {
		return NewInfo(true, 0)
	}
	leftInfo := process(x.Left)
	rightInfo := process(x.Right)
	// 左右树的最大高度
	height := int(math.Max(float64(leftInfo.height), float64(rightInfo.height))) + 1
	isBalanced := true // 是否平衡
	if !leftInfo.isBalanced {
		isBalanced = false
	}
	if !rightInfo.isBalanced {
		isBalanced = false
	}
	if math.Abs(float64(leftInfo.height-rightInfo.height)) > 1 {
		isBalanced = false
	}
	return NewInfo(isBalanced, height)
}
