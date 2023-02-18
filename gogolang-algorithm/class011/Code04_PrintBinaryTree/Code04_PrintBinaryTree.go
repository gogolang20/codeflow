package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func printTree(head *Node) {
	fmt.Println("Binary Tree:")
	printInOrder(head, 0, "H", 17)
	fmt.Println()
}

func printInOrder(head *Node, height int, to string, lens int) {
	if head == nil {
		return
	}
	printInOrder(head.right, height+1, "v", lens)
	val := to + strconv.Itoa(head.value) + to
	lenM := len(val)
	lenL := (lens - lenM) / 2
	lenR := lens - lenM - lenL
	val = getSpace(lenL) + val + getSpace(lenR)
	fmt.Println(getSpace(height*lens) + val)
	printInOrder(head.left, height+1, "^", lens)
}

func getSpace(num int) string {
	space := " "
	buf := make([]byte, 0)
	for i := 0; i < num; i++ {
		buf = append(buf, []byte(space))
	}
	return string(buf)
}
