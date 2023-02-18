package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

// need n extra space
func isPalindrome1(head *Node) bool {
	stack := list.New()
	cur := head
	for cur != nil {
		stack.PushFront(cur)
		cur = cur.next
	}
	for head != nil {
		if head.value != stack.Front().Value.(*Node).value {
			stack.Remove(stack.Front())
			return false
		}
		stack.Remove(stack.Front())
		head = head.next
	}
	return true
}

//判断链表是否回文
func isPalindrome3(head *Node) bool {
	if head == nil || head.next == nil {
		return true
	}
	slow := head                                    //slow
	fast := head                                    //fast
	for fast.next != nil && fast.next.next != nil { //找到链表的中点，奇书返回中点，偶数返回上中点
		slow = slow.next      // slow -> mid
		fast = fast.next.next // fast -> end
	}
	fast = slow.next // 后半段链表的 head'
	slow.next = nil  // 前半段链表指向 nil
	//准备反转后半段链表
	var temp *Node = nil
	for fast != nil {
		temp = fast.next
		fast.next = slow //修改指针方向
		slow = fast
		fast = temp
	}
	temp = slow //后半段链表的头
	fast = head //前半段链表的头
	res := true //要返回的节点
	//比对节点是否相等
	for slow != nil && fast != nil {
		if slow.value != fast.value {
			res = false
			break
		}
		slow = slow.next
		fast = fast.next
	}
	//返回结果之前还原链表
	slow = temp.next // n3是后半段的头节点，反转后成为最后一个节点
	temp.next = nil  // 将 n3 变量指针指向nil
	for slow != nil {
		fast = slow.next // fast 也可以作为临时变量
		slow.next = temp
		temp = slow // n3 最终到这个位置相当与后半段的头，且不会指向 nil
		slow = fast // slow 和 fast 最终等于 nil
	}
	return res
}

func main() {
	node := &Node{value: 1}
	node.next = &Node{value: 2}
	node.next.next = &Node{value: 3}
	node.next.next.next = &Node{value: 3}
	node.next.next.next.next = &Node{value: 2}
	node.next.next.next.next.next = &Node{value: 0}

	fmt.Println(isPalindrome1(node))
	fmt.Println(isPalindrome3(node))
	printNode(node)
}

func printNode(head *Node) {
	for head != nil {
		fmt.Printf("%v ==> ", head.value)
		head = head.next
	}
	fmt.Println()
}
