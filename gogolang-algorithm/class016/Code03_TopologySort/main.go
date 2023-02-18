package main

import (
	"container/list"
)

type Node struct { //图中的点
	value int
	in    int     //入度，有多个边指向自己
	out   int     //出度，有多个边直接指向其他
	nexts []*Node //从自己出发可以找到的直接邻居
	edges []*Edge //从自己出发能找到的边（没有就不用）
}

type Edge struct { //图中的边
	weight int //边的权重
	from   *Node
	to     *Node
}

type Graph struct { //图
	nodes map[int]*Node //编号 --> Node
	edges map[*Edge]struct{}
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
		in:    0, //没有指向自己的Node，值为0
		out:   0,
		nexts: make([]*Node, 0),
		edges: make([]*Edge, 0),
	}
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int]*Node),
		edges: make(map[*Edge]struct{}),
	}
}

/*
图的拓扑排序算法
1）在图中找到所有入度为0的点输出
2）把所有入度为0的点在图中删掉，继续找入度为0的点输出，周而复始
3）图的所有点都被删除后，依次输出的顺序就是拓扑排序

要求：有向图且其中没有环
应用：事件安排、编译顺序
*/
func sortedTopology(graph Graph) []*Node {
	// key 某个节点   value 剩余的入度
	inMap := make(map[*Node]int)
	// 只有剩余入度为0的点，才进入这个队列
	zeroInQueue := list.New()
	for _, value := range graph.nodes {
		inMap[value] = value.in //记录所有 Node的原始入度值
		if value.in == 0 {
			zeroInQueue.PushBack(value)
		}
	}
	result := make([]*Node, 0)
	for zeroInQueue.Len() > 0 {
		cur := zeroInQueue.Front() //每次从入度为0 的队列取出 Node
		zeroInQueue.Remove(cur)
		result = append(result, cur.Value.(*Node))
		for _, value := range cur.Value.(*Node).nexts {
			inMap[value] = inMap[value] - 1
			if inMap[value] == 0 {
				zeroInQueue.PushBack(value)
			}
		}
	}
	return result
}

/*
动态规划：记忆化搜索，记录走过的路
*/

//Code03_TopologicalOrderDFS1 统计最大深度，存入缓存中
//Code03_TopologicalOrderDFS2 统计点次，存入缓存中

func main() {

}
