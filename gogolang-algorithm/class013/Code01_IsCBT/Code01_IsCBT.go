package main

import "math"

type Node struct {
	value int
	left  *Node
	right *Node
}

/*
二叉树dp套路解 判断是否为完全二叉树
1 左满二叉树    右满    左高 == 右高
2 左完全二叉树   右满   左高 == 右高 + 1
3 左满二叉树    右满    左高 == 右高 + 1
4 左满二叉树    右完全  左高 == 右高

*/

type Info struct {
	isFull bool //满二叉树
	isCBT  bool //完全二叉树
	height int
}

func NewInfo(full, cbt bool, h int) *Info {
	return &Info{
		isFull: full,
		isCBT:  cbt,
		height: h,
	}
}

//推荐使用 dp 套路
func isCBT2(head *Node) bool {
	return process(head).isCBT
}

func process(x *Node) *Info {
	if x == nil {
		return NewInfo(true, true, 0)
	}
	leftInfo := process(x.left)
	rightInfo := process(x.right)
	height := int(math.Max(float64(leftInfo.height), float64(rightInfo.height))) + 1
	isFull := leftInfo.isFull && rightInfo.isFull && (leftInfo.height == rightInfo.height)
	isCBT := false
	if leftInfo.isFull && rightInfo.isFull && (leftInfo.height == rightInfo.height) {
		isCBT = true // 可能性一
	} else if leftInfo.isCBT && rightInfo.isFull && (leftInfo.height == rightInfo.height+1) {
		isCBT = true // 可能性二
	} else if leftInfo.isFull && rightInfo.isFull && (leftInfo.height == rightInfo.height+1) {
		isCBT = true // 可能性三
	} else if leftInfo.isFull && rightInfo.isCBT && (leftInfo.height == rightInfo.height) {
		isCBT = true // 可能性四
	}
	return NewInfo(isFull, isCBT, height)
}

func main() {

}
