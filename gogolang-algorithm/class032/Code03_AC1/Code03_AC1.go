package main

/*
AC自动机

解决在一个大字符串中，找到多个候选字符串的问题
*/

/*
AC自动机算法核心

1）把所有匹配串生成一棵前缀树

2）前缀树节点增加fail指针

3）fail指针的含义：如果必须以当前字符结尾，当前形成的路径是str，剩下哪一个字符串的前缀和str的后缀，拥有最大的匹配长度。fail指针就指向那个字符串的最后一个字符所对应的节点。（迷不迷？听讲述！）

*/


//type Node struct {
//	end   int // 有多少个字符串以该节点结尾
//	fail  *Node
//	nexts []*Node
//}
//
//func NewNode() *Node {
//	return &Node{
//		end:   0,
//		fail:  nil,
//		nexts: make([]*Node, 26),
//	}
//}
//
//// 你有多少个匹配串，就调用多少次insert
//func (node *Node) insert(s string) {
//	str := []byte(s)
//	for i := 0; i < len(str); i++ {
//		index := str[i] - 'a'
//		if node.nexts[index] == nil {
//			next := NewNode()
//			node.nexts[index] = next
//		}
//		node = node.nexts[index]
//	}
//	node.end++
//}
