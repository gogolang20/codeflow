package main

import "fmt"

//leetcode 547题地址
//https://leetcode-cn.com/problems/number-of-provinces/submissions/

//数组实现并查集
type UnionFind struct {
	parent []int // parent[i] = k ： i的父亲是k
	// i所在的集合大小是多少
	size []int // size[i] = k ： 如果i是代表节点，size[i]才有意义，否则无意义
	help []int // 辅助结构：压缩路径，做栈使用
	sets int   // 一共有多少个集合
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	help := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i //所有节点的父节点是自己
		size[i] = 1   //所有集合初始大小为 1
	}
	return &UnionFind{
		parent: parent,
		size:   size,
		help:   help,
		sets:   n,
	}
}
func (un *UnionFind) union(i, j int) {
	f1 := un.find(i)
	f2 := un.find(j)
	if f1 != f2 {
		if un.size[f1] >= un.size[f2] {
			un.size[f1] += un.size[f2]
			un.parent[f2] = f1
		} else {
			un.size[f2] += un.size[f1]
			un.parent[f1] = f2
		}
		un.sets--
	}
}

// 从i开始一直往上，往上到不能再往上，代表节点，返回
// 这个过程要做路径压缩
func (un *UnionFind) find(i int) int {
	hi := 0
	for i != un.parent[i] {
		un.help[hi] = i
		hi++
		i = un.parent[i]
	}
	for hi--; hi >= 0; hi-- { //路径压缩
		un.parent[un.help[hi]] = i
	}
	return i
}

func (un *UnionFind) setss() int {
	return un.sets
}

func findCircleNum(arr [][]int) int {
	unionFind := NewUnionFind(len(arr))
	//正方形二维数组，遍历上半部分
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i][j] == 1 {
				unionFind.union(i, j)
			}
		}
	}
	return unionFind.setss()
}

func main() {
	arr := NewArr(5, 5)
	fmt.Println(findCircleNum(arr))
}

//生成一个初始化完成的二维数组，不用提交
func NewArr(row, col int) [][]int {
	cir := make([][]int, row)
	for index := range cir {
		cir[index] = make([]int, col)
	}
	return cir
}
