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

func NewNode(value int) *Node {
	return &Node{
		value: value,
	}
}

/*
实现二叉树的序列化和反序列化
1）先序方式序列化和反序列化
	中序方式无法序列化，会出现歧义
2）按层方式序列化和反序列化
*/

//先序方式序列化
func preSerial(head *Node) *list.List {
	ans := list.New()
	pres(head, ans)
	return ans
}

func pres(head *Node, ans *list.List) {
	if head == nil {
		ans.PushBack("#") //占位
	} else {
		ans.PushBack(head.value)
		pres(head.left, ans)
		pres(head.right, ans)
	}
}

//先序方式反序列化
func buildByPreQueue(prelist *list.List) *Node {
	if prelist.Len() == 0 {
		return nil
	}
	return preb(prelist)
}

func preb(prelist *list.List) *Node {
	cur := prelist.Front()
	prelist.Remove(cur)
	switch cur.Value.(type) {
	case string:
		return nil
	default:
	}
	head := NewNode(cur.Value.(int))
	head.left = preb(prelist)
	head.right = preb(prelist)
	return head
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
	pre(node01)
	fmt.Println()

	res := preSerial(node01)
	printQueue(res.Front(), res.Len())

	fmt.Println()
	newres := buildByPreQueue(res)
	pre(newres)
}

func printQueue(head *list.Element, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%v ==> ", head.Value)
		head = head.Next()
	}
}

func pre(head *Node) {
	if head == nil {
		return
	}
	fmt.Printf("%v ==> ", head.value)
	pre(head.left)
	pre(head.right)
}
