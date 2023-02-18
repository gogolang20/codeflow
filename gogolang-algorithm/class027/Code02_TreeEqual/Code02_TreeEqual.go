package main

import (
	"container/list"
	"strconv"
)

/*
给定两棵二叉树的头节点head1和head2

想知道head1中是否有某个子树的结构和head2完全一样
*/

type Node struct {
	value int
	left  *Node
	right *Node
}

func containsTree1(big, small *Node) bool {
	if small == nil {
		return true
	}
	if big == nil {
		return false
	}
	if isSameValueStructure(big, small) {
		return true
	}
	return containsTree1(big.left, small) || containsTree1(big.right, small)
}

func isSameValueStructure(head1, head2 *Node) bool {
	if head1 == nil && head2 != nil {
		return false
	}
	if head1 != nil && head2 == nil {
		return false
	}
	if head1 == nil && head2 == nil {
		return true
	}
	if head1.value != head2.value {
		return false
	}
	return isSameValueStructure(head1.left, head2.left) && isSameValueStructure(head1.right, head2.right)
}

// 代码有待验证 ！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！
// 先序序列化后判断
func containsTree2(big, small *Node) bool {
	if small == nil {
		return true
	}
	if big == nil {
		return false
	}
	b := preSerial(big)
	s := preSerial(small)
	str := make([]string, b.Len())
	for i := 0; i < len(str); i++ {
		str[i] = b.Front().Value.(string)
		b.Remove(s.Front())
	}

	match := make([]string, s.Len())
	for i := 0; i < len(match); i++ {
		match[i] = s.Front().Value.(string)
		s.Remove(s.Front())
	}
	return getIndexOf(str, match) != -1
}

func preSerial(head *Node) *list.List {
	ans := list.New()
	pres(head, ans)
	return ans
}

func pres(head *Node, ans *list.List) {
	if head == nil {
		ans.PushBack(nil)
	} else {
		ans.PushBack(strconv.Itoa(head.value))
		pres(head.left, ans)
		pres(head.right, ans)
	}
}

func getIndexOf(str1, str2 []string) int {
	if str1 == nil || str2 == nil || len(str1) < 1 || len(str1) < len(str2) {
		return -1
	}
	x := 0
	y := 0
	next := getNextArray(str2)
	for x < len(str1) && y < len(str2) {
		if isEqual(str1[x], str2[y]) {
			x++
			y++
		} else if next[y] == -1 {
			x++
		} else {
			y = next[y]
		}
	}
	if y == len(str2) {
		return x - y
	} else {
		return -1
	}
}

func getNextArray(ms []string) []int {
	if len(ms) == 1 {
		return []int{-1}
	}
	next := make([]int, len(ms))
	next[0] = -1
	next[1] = 0
	i := 2
	cn := 0
	for i < len(next) {
		if isEqual(ms[i-1], ms[cn]) {
			next[i] = cn + 1
			i++
			cn++
		} else if cn > 0 {
			cn = next[cn]
		} else {
			next[i] = 0
			i++
		}
	}
	return next
}

func isEqual(a, b string) bool {
	if a == "" && b == "" {
		return true
	} else {
		if a == "" || b == "" {
			return false
		} else {
			return a == b
		}
	}
}
