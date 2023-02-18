package main

import (
	"fmt"
	"math"
)

/*
给定两个可能有环也可能无环的单链表，头节点head1和head2。请实现一个函数，如果两个链表相交，请返回相交的 第一个节点。
如果不相交，返回 nil
要求：
如果两个链表长度之和为N，时间复杂度请达到O(N)，额外空间复杂度 请达到O(1)。
*/

type Node struct {
	value int
	next  *Node
}

func getIntersectNode(head1, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}
	loop1 := getLoopNode(head1)
	loop2 := getLoopNode(head2)
	if loop1 == nil && loop2 == nil {
		return noLoop(head1, head2)
	}
	// 有环的三种情况：1 各自成环 2 环内相交，两个loopnode 3 环外相交
	if loop1 != nil && loop2 != nil {
		return bothLoop(head1, head2, loop1, loop2)
	}
	//一个有环 一个无环
	return nil
}

// 找到链表第一个入环节点，如果无环，返回nil
func getLoopNode(head *Node) *Node {
	if head == nil || head.next == nil || head.next.next == nil {
		return nil
	}
	slow := head.next
	fast := head.next.next
	for slow != fast {
		if fast.next == nil || fast.next.next == nil {
			return nil
		}
		slow = slow.next
		fast = fast.next.next
	}
	fast = head
	for fast != slow {
		fast = fast.next
		slow = slow.next
	}
	return slow
}

//两个链表都无环
func noLoop(head1, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}
	cur1 := head1
	cur2 := head2
	n := 0
	for cur1.next != nil {
		n++
		cur1 = cur1.next
	}
	for cur2.next != nil {
		n--
		cur2 = cur2.next
	}
	if cur1 != cur2 {
		return nil
	}
	//长的链表给 cur1
	if n > 0 {
		cur1 = head1
		cur2 = head2
	} else {
		cur1 = head2
		cur2 = head1
	}
	n = int(math.Abs(float64(n)))
	for n != 0 {
		n--
		cur1 = cur1.next
	}
	for cur1 != cur2 {
		cur1 = cur1.next
		cur2 = cur2.next
	}
	return cur1
}

// 两个有环链表，返回第一个相交节点，如果不想交返回nil
func bothLoop(head1, head2, loop1, loop2 *Node) *Node {
	var cur1 *Node = nil
	var cur2 *Node = nil
	if loop1 == loop2 {
		cur1 = head1
		cur2 = head2
		n := 0
		for cur1 != loop1 {
			n++
			cur1 = cur1.next
		}
		for cur2 != loop2 {
			n--
			cur2 = cur2.next
		}
		if n > 0 {
			cur1 = head1
			cur2 = head2
		} else {
			cur1 = head2
			cur2 = head1
		}
		n = int(math.Abs(float64(n)))
		for n != 0 {
			n--
			cur1 = cur1.next
		}
		for cur1 != cur2 {
			cur1 = cur1.next
			cur2 = cur2.next
		}
		return cur1
	} else {
		cur1 = loop1.next
		for cur1 != loop1 {
			if cur1 == loop2 {
				return loop1
			}
			cur1 = cur1.next
		}
		return nil
	}
}

func main() {
	// 1->2->3->4->5->6->7->nil
	head1 := &Node{value: 1}
	head1.next = &Node{value: 2}
	head1.next.next = &Node{value: 3}
	head1.next.next.next = &Node{value: 4}
	head1.next.next.next.next = &Node{value: 5}
	head1.next.next.next.next.next = &Node{value: 6}
	head1.next.next.next.next.next.next = &Node{value: 7}

	// 0->9->8->6->7->nil
	head2 := &Node{value: 0}
	head2.next = &Node{value: 9}
	head2.next.next = &Node{value: 8}
	head2.next.next.next = head1.next.next.next.next.next // 8->6
	fmt.Println(getIntersectNode(head1, head2).value)

	//// 1->2->3->4->5->6->7->4...
	head1 = &Node{value: 1}
	head1.next = &Node{value: 2}
	head1.next.next = &Node{value: 3}
	head1.next.next.next = &Node{value: 4}
	head1.next.next.next.next = &Node{value: 5}
	head1.next.next.next.next.next = &Node{value: 6}
	head1.next.next.next.next.next.next = &Node{value: 7}
	head1.next.next.next.next.next.next = head1.next.next.next // 7->4

	// 0->9->8->2...
	head2 = &Node{value: 0}
	head2.next = &Node{value: 9}
	head2.next.next = &Node{value: 8}
	head2.next.next.next = head1.next // 8->2
	fmt.Println(getIntersectNode(head1, head2).value)

	// 0->9->8->6->4->5->6..
	head2 = &Node{value: 0}
	head2.next = &Node{value: 9}
	head2.next.next = &Node{value: 8}
	head2.next.next.next = head1.next.next.next.next.next // 8->6
	fmt.Println(getIntersectNode(head1, head2).value)
}
