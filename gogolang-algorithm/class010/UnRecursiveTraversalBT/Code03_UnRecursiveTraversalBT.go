package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

/*
X 祖先节点 有且只有
先序遍历的左边 && 后序遍历的右边
两个集合的交集
*/

func pre(head *Node) {
	fmt.Println("pre-order: ")
	if head != nil {
		stack := list.New()
		stack.PushFront(head)
		for stack.Len() > 0 {
			head = stack.Front().Value.(*Node)
			stack.Remove(stack.Front())
			fmt.Printf("%v ==> ", head.value)
			if head.right != nil {
				stack.PushFront(head.right)
			}
			if head.left != nil {
				stack.PushFront(head.left)
			}
		}
	}
	fmt.Println()
}

func in(cur *Node) {
	fmt.Println("in-order: ")
	if cur != nil {
		stack := list.New()
		for stack.Len() > 0 || cur != nil {
			if cur != nil {
				stack.PushFront(cur)
				cur = cur.left
			} else {
				cur = stack.Front().Value.(*Node)
				stack.Remove(stack.Front())
				fmt.Printf("%v ==> ", cur.value)
				cur = cur.right
			}
		}
	}
	fmt.Println()
}

func pos1(head *Node) {
	fmt.Println("pos-order: ")
	if head != nil {
		s1 := list.New()
		s2 := list.New()
		s1.PushFront(head)
		for s1.Len() > 0 {
			head = s1.Front().Value.(*Node) // 头 右 左
			s1.Remove(s1.Front())
			s2.PushFront(head)
			if head.left != nil {
				s1.PushFront(head.left)
			}
			if head.right != nil {
				s1.PushFront(head.right)
			}
		}
		// 左 右 头
		for s2.Len() > 0 {
			fmt.Printf("%v ==> ", s2.Front().Value.(*Node).value)
			s2.Remove(s2.Front())
		}
	}
	fmt.Println()
}

func pos2(h *Node) {
	fmt.Println("pos-order: ")
	if h != nil {
		stack := list.New()
		stack.PushFront(h)
		var c *Node = nil
		for stack.Len() > 0 {
			c = stack.Front().Value.(*Node)
			if c.left != nil && h != c.left && h != c.right {
				stack.PushFront(c.left)
			} else if c.right != nil && h != c.right {
				stack.PushFront(c.right)
			} else {
				fmt.Printf("%v ==> ", stack.Front().Value.(*Node).value)
				stack.Remove(stack.Front())
				h = c
			}
		}
	}
	fmt.Println()
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

	pre(node01)
	fmt.Println()

	in(node01)
	// 4 ==> 2 ==> 5 ==> 1 ==> 6 ==> 3 ==> 7 ==>
	// 4 ==> 2 ==> 5 ==> 1 ==> 6 ==> 3 ==> 7 ==>
	fmt.Println()

	pos1(node01)
	pos2(node01)
	// 4 ==> 5 ==> 2 ==> 6 ==> 7 ==> 3 ==> 1 ==>
	// 4 ==> 5 ==> 2 ==> 6 ==> 7 ==> 3 ==> 1 ==>
}
