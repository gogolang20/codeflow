package main

import (
	"container/list"
	"fmt"
)

/*
并查集
1 有若干个样本a、b、c、d…类型假设是V
2 在并查集中一开始认为每个样本都在单独的集合里
3 用户可以在任何时候调用如下两个方法：
       boolean isSameSet(V x, V y) : 查询样本x和样本y是否属于一个集合
       void union(V x, V y) : 把x和y各自所在集合的所有样本合并成一个集合
4 isSameSet和union方法的代价越低越好

1）每个节点都有一条往上指的指针
2）节点a往上找到的头节点，叫做a所在集合的代表节点
3）查询x和y是否属于同一个集合，就是看看找到的代表节点是不是一个
4）把x和y各自所在集合的所有点合并成一个集合，只需要小集合的代表点挂在大集合的代表点的下方即可


并查集的优化
1）节点往上找代表点的过程，压缩路径
2）小集合挂在大集合的下面
3）如果方法调用很频繁，那么单次调用的代价为O(1)，两个方法都如此
*/

type Node struct { //给值包一层
	value string
}

type UnionFind struct {
	nodes   map[string]Node // key是样本，value是对应包的一层圈
	parents map[Node]Node   // parent表记录：key 代表子节点，value 代表父节点
	sizeMap map[Node]int    // 只记录作为代表的节点
}

//传入一个样本集合，初始化完成后返回
func NewUnionFind(arr []string) *UnionFind {
	nodes := make(map[string]Node, len(arr))
	parents := make(map[Node]Node, len(arr))
	sizeMap := make(map[Node]int, len(arr))
	for _, v := range arr {
		node := Node{value: v} //每个样本生成对应的代表节点
		nodes[v] = node
		parents[node] = node
		sizeMap[node] = 1
	}
	return &UnionFind{
		nodes:   nodes,
		parents: parents,
		sizeMap: sizeMap,
	}
}

// 给你一个节点，请你往上到不能再往上，把代表返回
func (uf *UnionFind) findFather(cur Node) Node {
	stack := list.New()
	for cur != uf.parents[cur] {
		stack.PushFront(cur)
		cur = uf.parents[cur]
	}
	for stack.Len() > 0 { //优化：路径压缩
		res := stack.Front()
		stack.Remove(res)
		uf.parents[res.Value.(Node)] = cur
	}
	return cur
}
func (uf *UnionFind) isSameSet(a, b string) bool {
	//a,b 的代表节点是否是一个
	return uf.findFather(uf.nodes[a]) == uf.findFather(uf.nodes[b])
}

func (uf *UnionFind) union(a, b string) {
	aHead := uf.findFather(uf.nodes[a])
	bHead := uf.findFather(uf.nodes[b])
	if aHead != bHead {
		aSetSize := uf.sizeMap[aHead]
		bSetSize := uf.sizeMap[bHead]
		var big, small Node //优化：小的挂大的
		if aSetSize > bSetSize {
			big = aHead
			small = bHead
		} else {
			big = bHead
			small = aHead
		}
		uf.parents[small] = big
		uf.sizeMap[big] = aSetSize + bSetSize
		delete(uf.sizeMap, small)
	}
}

//没啥作用
func (uf *UnionFind) setss() int {
	return len(uf.sizeMap)
}

func main() {
	//传入样本的类型string，Node是给 样本包的一层
	uf := NewUnionFind([]string{"asd", "qwe", "123"})
	fmt.Println(uf.setss())
	uf.union("asd", "qwe")
	fmt.Println(uf.setss())
	fmt.Println(uf.findFather(Node{value: "qwe"}))
	fmt.Println(uf.findFather(Node{value: "asd"}))
	fmt.Println(uf.findFather(Node{value: "123"}))
}
