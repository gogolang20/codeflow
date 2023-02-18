package main

import "math"

type Node struct {
	value int
	Left  *Node
	Right *Node
}

/*
给定一棵二叉树的头节点head，任何两个节点之间都存在距离，
返回整棵二叉树的最大距离
*/

// 1 与 x 节点有关：经过 x
// 2 与 x 节点无关

type Info struct {
	maxDistance int
	height      int
}

func NewInfo(m, h int) *Info {
	return &Info{
		maxDistance: m,
		height:      h,
	}
}

func maxDistance2(head *Node) int {
	if head == nil {
		return 0
	}
	return process(head).maxDistance
}
func process(x *Node) *Info {
	if x == nil {
		return NewInfo(0, 0)
	}
	leftInfo := process(x.Left)
	rightInfo := process(x.Right)
	height := int(math.Max(float64(leftInfo.height), float64(rightInfo.height))) + 1
	p1 := leftInfo.maxDistance
	p2 := rightInfo.maxDistance
	p3 := leftInfo.height + rightInfo.height + 1
	maxDistance := int(math.Max(math.Max(float64(p1), float64(p2)), float64(p3)))
	return NewInfo(maxDistance, height)
}
