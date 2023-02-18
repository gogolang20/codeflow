package main

import "container/list"

type Node struct {
	value int
	Left  *Node
	Right *Node
}

// 判断二叉树是否是完全二叉树

// 方法一
func isCBT1(head *Node) bool {
	if head == nil {
		return true
	}
	queue := list.New() // 二叉树按层遍历
	leaf := false
	var l *Node = nil
	var r *Node = nil
	queue.PushBack(head)
	for queue.Len() > 0 {
		temp := queue.Front()
		queue.Remove(temp)
		l = temp.Value.(*Node).Left
		r = temp.Value.(*Node).Right
		if (leaf && (l != nil || r != nil)) || (l == nil && r != nil) { // 如果遇到了不双全的节点之后，又发现当前节点不是叶节点
			return false
		}
		if l != nil {
			queue.PushBack(l)
		}
		if r != nil {
			queue.PushBack(r)
		}
		if l == nil || r == nil {
			leaf = true
		}
	}
	return true
}
