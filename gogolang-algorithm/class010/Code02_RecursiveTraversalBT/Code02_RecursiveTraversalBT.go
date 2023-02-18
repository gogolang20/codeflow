package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

/*
递归序
a b d d d b e e e b a c f f f c g g g c a
*/
// 先序遍历
func pre(head *Node) {
	if head == nil {
		return
	}
	fmt.Printf("%v ==> ", head.value)
	pre(head.left)
	pre(head.right)
}

// 中序遍历
func in(head *Node) {
	if head == nil {
		return
	}
	in(head.left)
	fmt.Printf("%v ==> ", head.value)
	in(head.right)
}

// 后序遍历
func pos(head *Node) {
	if head == nil {
		return
	}
	pos(head.left)
	pos(head.right)
	fmt.Printf("%v ==> ", head.value)
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

	//pre(node01)
	//fmt.Println()
	in(node01)
	fmt.Println()
	pos(node01)
}
