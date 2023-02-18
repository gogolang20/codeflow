package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

//快慢指针

// 1）输入链表头节点，奇数长度返回中点，偶数长度返回上中点
func midOrUpMidNode(head *Node) *Node {
	if head == nil || head.next == nil || head.next.next == nil {
		return head
	}
	slow := head.next
	fast := head.next.next
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// 2）输入链表头节点，奇数长度返回中点，偶数长度返回下中点
func midOrDownMidNode(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	slow := head.next
	fast := head.next
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// 3）输入链表头节点，奇数长度返回中点前一个，偶数长度返回上中点前一个
func midOrUpMidPreNode(head *Node) *Node {
	if head == nil || head.next == nil || head.next.next == nil {
		return nil
	}
	slow := head
	fast := head.next.next
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// 4）输入链表头节点，奇数长度返回中点前一个，偶数长度返回下中点前一个
func midOrDownMidPreNode(head *Node) *Node {
	if head == nil || head.next == nil {
		return nil
	}
	if head.next.next == nil {
		return head
	}
	slow := head
	fast := head.next
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func main() {
	node := &Node{value: 1}
	node.next = &Node{value: 2}
	node.next.next = &Node{value: 3}
	node.next.next.next = &Node{value: 4}
	node.next.next.next.next = &Node{value: 5}
	node.next.next.next.next.next = &Node{value: 6}
	node.next.next.next.next.next.next = &Node{value: 7}
	node.next.next.next.next.next.next.next = &Node{value: 8}
	node.next.next.next.next.next.next.next.next = &Node{value: 9}

	fmt.Println(midOrUpMidNode(node).value)
	fmt.Println(midOrDownMidNode(node).value)
	fmt.Println(midOrUpMidPreNode(node).value)
	fmt.Println(midOrDownMidPreNode(node).value)

}

func printNode(head *Node) {
	for head != nil {
		fmt.Printf("%v ==> ", head.value)
		head = head.next
	}
	fmt.Println()
}
