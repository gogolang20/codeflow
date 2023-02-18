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

//求二叉树最宽的层有多少个节点
func maxWidthNoMap(head *Node) int {
	if head == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(head)
	curEnd := head          // 当前层，最右节点是谁
	var nextEnd *Node = nil // 下一层，最右节点是谁
	max := 0                //记录最大宽度
	curLevelNodes := 0
	for queue.Len() > 0 {
		cur := queue.Front() //要弹出的节点
		queue.Remove(cur)
		if cur.Value.(*Node).left != nil {
			queue.PushBack(cur.Value.(*Node).left)
			nextEnd = cur.Value.(*Node).left
		}
		if cur.Value.(*Node).right != nil {
			queue.PushBack(cur.Value.(*Node).right)
			nextEnd = cur.Value.(*Node).right
		}
		curLevelNodes++
		if cur.Value.(*Node) == curEnd {
			if max < curLevelNodes {
				max = curLevelNodes
			}
			curLevelNodes = 0
			curEnd = nextEnd
		}
	}
	return max
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

	fmt.Println(maxWidthNoMap(node01))
}
