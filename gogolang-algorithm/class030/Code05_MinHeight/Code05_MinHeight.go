package main

import "math"

type Node struct {
	value int
	left  *Node
	right *Node
}

/*
给定一棵二叉树的头节点head

求以head为头的树中，最小深度是多少？
*/

func minHeight1(head *Node) int {
	if head == nil {
		return 0
	}
	return p(head)
}

// 返回x为头的树，最小深度是多少
func p(x *Node) int {
	if x.left == nil && x.right == nil {
		return 1
	}
	// 左右子树起码有一个不为空
	leftH := math.MaxInt
	if x.left != nil {
		leftH = p(x.left)
	}
	rightH := math.MaxInt
	if x.right != nil {
		rightH = p(x.right)
	}
	return 1 + Min(leftH, rightH)
}

// 根据morris遍历改写
func minHeight2(head *Node) int {
	if head == nil {
		return 0
	}
	cur := head
	var mostRight *Node = nil
	curLevel := 0
	minHeight := math.MaxInt
	for cur != nil {
		mostRight = cur.left
		if mostRight != nil {
			rightBoardSize := 1
			for mostRight.right != nil && mostRight.right != cur {
				rightBoardSize++
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				curLevel++
				mostRight.right = cur
				cur = cur.left
				continue
			} else {
				if mostRight.left == nil {
					minHeight = Min(minHeight, curLevel)
				}
				curLevel -= rightBoardSize
				mostRight.right = nil
			}
		} else { // 只有一次到达
			curLevel++
		}
		cur = cur.right
	}
	finalRight := 1
	cur = head
	for cur.right != nil {
		finalRight++
		cur = cur.right
	}
	if cur.left == nil && cur.right == nil {
		minHeight = Min(minHeight, finalRight)
	}
	return minHeight
}

func main() {

}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
