package main

import "fmt"

/*
IndexTree

特点：
1）支持区间查询
2）没有线段树那么强，但是非常容易改成一维、二维、三维的结构
3）只支持单点更新
*/

type IndexTree struct {
	tree []int // 下标从1开始！！！
	N    int
}

func NewIndexTree(size int) *IndexTree {
	return &IndexTree{
		tree: make([]int, size+1), // 0位置弃而不用！！！
		N:    size,
	}
}

// 1~index 累加和是多少？
func (it *IndexTree) sum(index int) int {
	ret := 0
	for index > 0 {
		ret += it.tree[index]
		index -= index & -index // index 减去二进制最右侧的 1
	}
	return ret
}

// index & -index : 提取出index最右侧的1出来
// index :           0011001000
// index & -index :  0000001000
func (it *IndexTree) add(index, d int) {
	for index <= it.N {
		it.tree[index] += d
		index += index & -index
	}
}

func main() {
	size := 8
	tree := NewIndexTree(size)
	tree.add(1,5)
	tree.add(2,6)
	fmt.Println(tree.sum(1))
}
