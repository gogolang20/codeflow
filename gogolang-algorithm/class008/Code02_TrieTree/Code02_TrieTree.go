package main

import "fmt"

// 该程序完全正确

type Node struct {
	pass int //节点通过次数
	end  int //字符串结尾的数量
	// 0    a
	// 1    b
	// 2    c
	// ..   ..
	// 25   z
	// nexts[i] == nil   i方向的路不存在
	// nexts[i] != nil   i方向的路存在
	nexts [26]*Node
}

//添加某个字符串，可以重复添加，每次算1个
func (node *Node) insert(str string) {
	if str == "" {
		return
	}
	arr := []byte(str)
	node.pass++
	for index := range arr {
		path := arr[index] - 'a'
		if node.nexts[path] == nil {
			node.nexts[path] = &Node{pass: 0, end: 0}
		}
		node = node.nexts[path]
		node.pass++
	}
	node.end++
}

//删掉某个字符串，可以重复删除，每次算1个
func (node *Node) delete(str string) {
	if node.search(str) != 0 {
		arr := []byte(str)
		node.pass--
		for i := 0; i < len(arr); i++ {
			path := arr[i] - 'a'
			node.nexts[path].pass--
			if node.nexts[path].pass == 0 {
				node.nexts[path] = nil
				return
			}
			node = node.nexts[path]
		}
		node.end--
	}
}

//查询某个字符串在结构中还有几个
func (node *Node) search(str string) int {
	if str == "" {
		return 0
	}
	arr := []byte(str)
	for i := 0; i < len(arr); i++ {
		index := arr[i] - 'a'
		if node.nexts[index] == nil {
			return 0
		}
		node = node.nexts[index]
	}
	return node.end
}

//查询有多少个字符串，是以str做前缀的
func (node *Node) prefixNumber(str string) int {
	if str == "" {
		return 0
	}
	arr := []byte(str)
	for i := 0; i < len(arr); i++ {
		index := arr[i] - 'a'
		if node.nexts[index] == nil {
			return 0
		}
		node = node.nexts[index]
	}
	return node.pass
}

func main() {
	node := &Node{pass: 0, end: 0}
	node.insert("string")
	node.insert("string")
	node.insert("string")
	node.insert("strings")

	fmt.Println(node.search("st"))
	fmt.Println(node.search("string"))
	fmt.Println(node.search("strings"))

	node.delete("string")

	fmt.Println(node.prefixNumber("st"))
	fmt.Println(node.prefixNumber("string"))
	fmt.Println(node.prefixNumber("strings"))
}
