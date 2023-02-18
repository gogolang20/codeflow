package main

import "fmt"

/*
把给定值都删除
    找到第一个不需要删除的节点，做为新的节点返回
    记录第一个节点，在不断追加后续的节点
*/

type Node struct {
	value int
	next  *Node
}

func removeValue(head *Node, num int) *Node {
	// 找到第一个不需要删除的 Node
	for head != nil {
		if head.value != num {
			break
		}
		head = head.next
	}
	// 开始删除
	pre := head
	cur := head
	for cur != nil {
		if cur.value == num {
			pre.next = cur.next
		} else {
			pre = cur
		}
		cur = cur.next
	}
	return head
}
func printNode(head *Node) {
	for head != nil {
		fmt.Printf("%v ===> ", head.value)
		head = head.next
	}
	fmt.Println()
}
func main() {
	node1 := &Node{value: 1}
	node2 := &Node{value: 1}
	node3 := &Node{value: 2}
	node4 := &Node{value: 1}
	node5 := &Node{value: 4}
	node6 := &Node{value: 1}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6

	printNode(node1)
	res := removeValue(node1, 1)
	printNode(res)
}
