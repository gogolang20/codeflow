package main

import (
	"fmt"
	"math"
)

/*
Morris遍历

一种遍历二叉树的方式，并且时间复杂度O(N)，额外空间复杂度O(1)

通过利用原树中大量空闲指针的方式，达到节省空间的目的
*/

/*
Morris遍历细节

假设来到当前节点cur，开始时cur来到头节点位置
1）如果cur没有左孩子，cur向右移动(cur = cur.right)
2）如果cur有左孩子，找到左子树上最右的节点mostRight：
	a.如果mostRight的右指针指向空，让其指向cur，
	然后cur向左移动(cur = cur.left)
	b.如果mostRight的右指针指向cur，让其指向null，
	然后cur向右移动(cur = cur.right)
3）cur为空时遍历停止
*/

/*
Morris遍历实质

建立一种机制：

对于没有左子树的节点只到达一次，

对于有左子树的节点会到达两次

morris遍历时间复杂度依然是O(N)
*/

type Node struct {
	value int
	left  *Node
	right *Node
}

func morris(head *Node) {
	if head == nil {
		return
	}
	cur := head
	var mostRight *Node = nil
	for cur != nil {
		mostRight = cur.left
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil { // 第一次遍历到
				mostRight.right = cur
				cur = cur.left
				continue
			} else { // 第二次遍历到
				mostRight.right = nil
			}
		}
		cur = cur.right
	}
}

func morrisPre(head *Node) {
	if head == nil {
		return
	}
	cur := head
	var mostRight *Node = nil
	for cur != nil {
		mostRight = cur.left
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil { // 第一次遍历到
				fmt.Printf("%v ==> ", cur.value)
				mostRight.right = cur
				cur = cur.left
				continue
			} else { // 第二次遍历到
				mostRight.right = nil
			}
		} else {
			fmt.Printf("%v ==> ", cur.value)
		}
		cur = cur.right
	}
	fmt.Println()
}

func morrisIn(head *Node) {
	if head == nil {
		return
	}
	cur := head
	var mostRight *Node = nil
	for cur != nil {
		mostRight = cur.left
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil { // 第一次遍历到
				mostRight.right = cur
				cur = cur.left
				continue
			} else { // 第二次遍历到
				mostRight.right = nil
			}
		}
		fmt.Printf("%v ==> ", cur.value)
		cur = cur.right
	}
	fmt.Println()
}

func morrisPos(head *Node) {
	if head == nil {
		return
	}
	cur := head
	var mostRight *Node = nil
	for cur != nil {
		mostRight = cur.left
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil { // 第一次遍历到
				mostRight.right = cur
				cur = cur.left
				continue
			} else { // 第二次遍历到
				mostRight.right = nil
				printEdge(cur.left)
			}
		}
		cur = cur.right
	}
	printEdge(head)
	fmt.Println()
}

func printEdge(head *Node) {
	tail := reverseEdge(head)
	cur := tail
	for cur != nil {
		fmt.Printf("%v ==> ", cur.value)
		cur = cur.right
	}
	reverseEdge(tail)
}

func reverseEdge(from *Node) *Node {
	var pre *Node = nil
	var next *Node = nil
	for from != nil {
		next = from.right
		from.right = pre
		pre = from
		from = next
	}
	return pre
}

func main() {
	head := &Node{value: 4}
	head.left = &Node{value: 2}
	head.right = &Node{value: 6}
	head.left.left = &Node{value: 1}
	head.left.right = &Node{value: 3}
	head.right.left = &Node{value: 5}
	head.right.right = &Node{value: 7}

	morrisIn(head)
	morrisPre(head)
	morrisPos(head)
}

func isBST(head *Node) bool {
	if head == nil {
		return true
	}
	cur := head
	var mostRight *Node = nil

	pre := math.MinInt // 上一个node的值
	ans := true
	for cur != nil {
		mostRight = cur.left
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil { // 第一次遍历到
				mostRight.right = cur
				cur = cur.left
				continue
			} else { // 第二次遍历到
				mostRight.right = nil
			}
		}
		if pre >= cur.value {
			ans = false
		}
		pre = cur.value
		cur = cur.right
	}
	return ans
}
