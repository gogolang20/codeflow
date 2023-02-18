package main

import "fmt"

/*
单向链表反转
*/
type node struct {
	value int
	next  *node
}

// head
//
//	a    ->   b    ->  c  ->  nil
//	c    ->   b    ->  a  ->  nil
func reverseLinkedList(head *node) *node {
	var pre *node = nil
	var next *node = nil
	for head != nil {
		next = head.next
		head.next = pre
		pre = head
		head = next
	}
	return pre
}

/*
双向链表反转
*/
type Node struct {
	value int
	prev  *Node // 节点的前一个
	next  *Node // 节点的后一个
}

func reverseDoubleList(head *Node) *Node {
	var pre *Node = nil
	var next *Node = nil
	for head != nil {
		next = head.next
		head.next = pre  // 修改节点的指向
		head.prev = next // 修改节点的指向
		pre = head
		head = next
	}
	return pre
}

// 打印链表
func printNode(head *Node) {
	for head != nil {
		fmt.Printf("%v ==> ", head.value)
		head = head.next
	}
	fmt.Println()
}

func main() {
	node1 := &Node{value: 0}
	node2 := &Node{value: 1}
	node3 := &Node{value: 2}
	node4 := &Node{value: 3}

	node1.prev = nil
	node1.next = node2

	node2.prev = node1
	node2.next = node3

	node3.prev = node2
	node3.next = node4

	node4.prev = node3
	node4.next = nil

	// 反转前打印
	printNode(node1)
	// 反转后打印
	head := reverseDoubleList(node1)
	printNode(head)
}
