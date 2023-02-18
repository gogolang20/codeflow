package main

type Node struct {
	value int
	left  *Node
	right *Node
}

/*
给定一棵二叉树的头节点head，和另外两个节点a和b。
返回a和b的最低公共祖先
*/

type Info struct {
	findA bool
	findB bool
	ans   *Node
}

func NewInfo(fa, fb bool, ans *Node) *Info {
	return &Info{
		findA: fa,
		findB: fb,
		ans:   ans,
	}
}
func lowestAncestor2(head, a, b *Node) *Node {
	return process(head, a, b).ans
}
func process(x, a, b *Node) *Info {
	if x == nil {
		return NewInfo(false, false, nil)
	}
	leftInfo := process(x.left, a, b)
	rightInfo := process(x.right, a, b)
	findA := (x == a) || leftInfo.findA || rightInfo.findA
	findB := (x == b) || leftInfo.findB || rightInfo.findB
	var ans *Node = nil
	if leftInfo.ans != nil {
		ans = leftInfo.ans
	} else if rightInfo.ans != nil {
		ans = rightInfo.ans
	} else {
		if findA && findB {
			ans = x
		}
	}
	return NewInfo(findA, findB, ans)
}

func main() {

}
