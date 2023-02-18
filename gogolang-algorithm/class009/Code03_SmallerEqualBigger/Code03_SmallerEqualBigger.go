package main

import "fmt"

//将单向链表按某值划分成左边小、中间相等、右边大的形式
type Node struct {
	value int
	next  *Node
}

func listPartition2(head *Node, pivot int) *Node {
	var sH *Node = nil // small head
	var sT *Node = nil // small tail
	var eH *Node = nil // equal head
	var eT *Node = nil // equal tail
	var mH *Node = nil // big head
	var mT *Node = nil // big tail
	var Next *Node = nil
	for head != nil {
		Next = head.next
		head.next = nil
		if head.value < pivot {
			if sH == nil {
				sH = head
				sT = head
			} else {
				sT.next = head
				sT = head
			}
		} else if head.value == pivot {
			if eH == nil {
				eH = head
				eT = head
			} else {
				eT.next = head
				eT = head
			}
		} else {
			if mH == nil {
				mH = head
				mT = head
			} else {
				mT.next = head
				mT = head
			}
		}
		head = Next
	}
	if sT != nil {
		sT.next = eH
		if eT == nil {
			eT = sT
		}
	}
	if eT != nil {
		eT.next = mH
	}
	if sH != nil {
		return sH
	} else if eH != nil {
		return eH
	} else {
		return mH
	}
}

func main() {
	node := &Node{value: 6}
	node.next = &Node{value: 5}
	node.next.next = &Node{value: 4}
	node.next.next.next = &Node{value: 3}
	node.next.next.next.next = &Node{value: 2}
	node.next.next.next.next.next = &Node{value: 1}

	res := listPartition2(node, 4)
	printNode(res)
}

func printNode(head *Node) {
	for head != nil {
		fmt.Printf("%v ==> ", head.value)
		head = head.next
	}
	fmt.Println()
}
