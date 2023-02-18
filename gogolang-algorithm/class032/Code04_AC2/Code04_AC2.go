package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	end    string
	endUse bool
	fail   *Node
	nexts  []*Node
}

func NewNode() *Node {
	return &Node{
		endUse: false,
		end:    "",
		fail:   nil,
		nexts:  make([]*Node, 26),
	}
}

func (node *Node) insert(s string) {
	str := []byte(s)
	for i := 0; i < len(str); i++ {
		index := str[i] - 'a'
		if node.nexts[index] == nil {
			node.nexts[index] = NewNode()
		}
		node = node.nexts[index]
	}
	node.end = s
}

// 添加完成所有字符串后，链接 fail 指针
func (node *Node) build() {
	queue := list.New()
	queue.PushBack(node)
	for queue.Len() > 0 {
		// 某个父亲，cur
		cur := queue.Front().Value.(*Node)
		queue.Remove(queue.Front())
		for i := 0; i < 26; i++ { // 所有的路
			// cur -> 父亲  i号儿子，必须把i号儿子的fail指针设置好！
			if cur.nexts[i] != nil { // 如果真的有i号儿子
				cur.nexts[i].fail = node
				cfail := cur.fail  // 开始链寻找fail指针
				for cfail != nil { // 寻找我的子该链谁
					if cfail.nexts[i] != nil {
						cur.nexts[i].fail = cfail.nexts[i] // 找到了，重新设置指针
						break
					}
					cfail = cfail.fail
				}
				queue.PushBack(cur.nexts[i])
			}
		}
	}
}

// 大文章：content
func (node *Node) containWords(content string) []string {
	str := []byte(content)
	cur := node
	ans := make([]string, 0)
	for i := 0; i < len(str); i++ {
		index := str[i] - 'a' // 路
		// 如果当前字符在这条路上没配出来，就随着fail方向走向下条路径
		for cur.nexts[index] == nil && cur != node {
			cur = cur.fail
		}
		// 1) 现在来到的路径，是可以继续匹配的
		// 2) 现在来到的节点，就是前缀树的根节点
		if cur.nexts[index] != nil {
			cur = cur.nexts[index]
		} else {
			cur = node
		}
		follow := cur
		for follow != node {
			if follow.endUse {
				break
			}
			// 不同的需求，在这一段之间修改
			if follow.end != "" {
				ans = append(ans, follow.end)
				follow.endUse = true
			}
			// 不同的需求，在这一段之间修改
			follow = follow.fail
		}
	}
	return ans
}

func main() {
	node := NewNode()
	file1 := "abcde"
	node.insert(file1)
	file2 := "abcd"
	node.insert(file2)
	file3 := "abc"
	node.insert(file3)
	file4 := "ab"
	node.insert(file4)
	file5 := "a"
	node.insert(file5)

	node.build()

	testFile := "a"
	fmt.Println(node.containWords(testFile))
}