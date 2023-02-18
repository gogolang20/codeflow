package main

import "math"

type Node struct {
	value int
	left  *Node
	right *Node
}

type Info struct {
	isBST bool
	max   int
	min   int
}

func NewInfo(is bool, max, min int) *Info {
	return &Info{
		isBST: is,
		max:   max,
		min:   min,
	}
}

func isBST1(head *Node) bool {
	if head == nil {
		return true
	}
	return process(head).isBST
}

func process(x *Node) *Info {
	if x == nil {
		return nil
	}
	leftInfo := process(x.left)
	rightInfo := process(x.right)
	max, min := x.value, x.value
	if leftInfo != nil {
		max = int(math.Max(float64(leftInfo.max), float64(max)))
		min = int(math.Min(float64(leftInfo.min), float64(min)))
	}
	if rightInfo != nil {
		max = int(math.Max(float64(rightInfo.max), float64(max)))
		min = int(math.Min(float64(rightInfo.min), float64(min)))
	}
	isBST := true
	if leftInfo != nil && !leftInfo.isBST {
		isBST = false
	}
	if rightInfo != nil && !rightInfo.isBST {
		isBST = false
	}
	if leftInfo != nil && leftInfo.max >= x.value {
		isBST = false
	}
	if rightInfo != nil && rightInfo.min <= x.value {
		isBST = false
	}
	return NewInfo(isBST, max, min)
}
