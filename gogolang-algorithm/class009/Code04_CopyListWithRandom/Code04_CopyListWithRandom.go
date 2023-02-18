package main

/*
一种特殊的单链表节点
rand指针是单链表节点结构中新增的指针，rand可能指向链表中的任意一个节点，也可能指向null。
给定一个由Node节点类型组成的无环单链表的头节点 head，请实现一个函数完成这个链表的复制，并返回复制的新链表的头节点。

[要求] 时间复杂度O(N)，额外空间复杂度O(1)
*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node //随机指针，可能指向任意 Node
}

// https://leetcode.cn/problems/copy-list-with-random-pointer/
// 测试通过
func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	//1 将clone节点插入到 老节点的next
	cur := head
	var next *Node = nil
	for cur != nil {
		next = cur.Next
		cur.Next = &Node{Val: cur.Val}
		cur.Next.Next = next
		cur = next
	}
	//2 根据老链表先将clone节点的 rand 设置好
	cur = head
	var deepcopy *Node = nil
	for cur != nil {
		next = cur.Next.Next
		deepcopy = cur.Next
		if cur.Random != nil {
			deepcopy.Random = cur.Random.Next // 新链表的 rand 指针
		} else {
			deepcopy.Random = nil
		}
		cur = next
	}
	//3 将clone链表的 next 与老链表断开
	cur = head
	res := head.Next //新链表的头
	for cur != nil {
		next = cur.Next.Next
		deepcopy = cur.Next
		cur.Next = next
		if next != nil {
			deepcopy.Next = next.Next
		} else {
			deepcopy.Next = nil
		}
		cur = next
	}
	return res
}
