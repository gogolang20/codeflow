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
实现二叉树的按层遍历
1）其实就是宽度优先遍历，用队列
2）可以通过设置flag变量的方式，来发现某一层的结束
*/

/*
两步实现
1 队列出一个 cur，打印
2 cur 有左入左  有右入右
*/
func level(head *Node) {
	if head == nil {
		return
	}
	queue := list.New()
	queue.PushBack(head)
	for queue.Len() > 0 {
		cur := queue.Front().Value.(*Node) // 队列中第一个元素
		queue.Remove(queue.Front())        // 移除队列
		fmt.Printf("%v ==> ", cur.value)   // 打印
		if cur.left != nil {
			queue.PushBack(cur.left)
		}
		if cur.right != nil {
			queue.PushBack(cur.right)
		}
	}
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

	level(node01)
	fmt.Println()
	pre(node01)
}

func pre(head *Node) {
	if head == nil {
		return
	}
	fmt.Printf("%v ==> ", head.value)
	pre(head.left)
	pre(head.right)
}
