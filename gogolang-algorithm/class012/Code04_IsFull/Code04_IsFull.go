package main

import (
	"fmt"
	"math"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type Info struct {
	nodes  int
	height int
}

func NewInfo(n, h int) *Info {
	return &Info{
		nodes:  n,
		height: h,
	}
}

func isFull2(head *Node) bool {
	all := process(head)
	return (1<<all.height)-1 == all.nodes
}

func process(x *Node) *Info {
	if x == nil {
		return NewInfo(0, 0)
	}
	leftInfo := process(x.left)
	rightInfo := process(x.right)
	height := int(math.Max(float64(leftInfo.height), float64(rightInfo.height))) + 1
	nodes := leftInfo.nodes + rightInfo.nodes + 1
	return NewInfo(nodes, height)
}

func main() {
	var node01 = &Node{value: 1}
	var node02 = &Node{value: 2}
	var node03 = &Node{value: 3}
	var node04 = &Node{value: 4}
	var node05 = &Node{value: 5}
	var node06 = &Node{value: 6}
	var node07 = &Node{value: 7}

	node01.left = node02
	node01.right = node03
	node02.left = node04
	node02.right = node05
	node03.left = node06
	node03.right = node07

	fmt.Println(isFull2(node01))
}
