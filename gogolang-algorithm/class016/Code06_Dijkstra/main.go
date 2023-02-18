package main

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
Dijkstra算法

1）Dijkstra算法必须指定一个源点
2）生成一个源点到各个点的最小距离表，一开始只有一条记录，即原点到自己的最小距离为0，源点到其他所有点的最小距离都为正无穷大
3）从距离表中拿出没拿过记录里的最小记录，通过这个点发出的边，更新源点到各个点的最小距离表，不断重复这一步
4）源点到所有的点记录如果都被拿过一遍，过程停止，最小距离表得到了

有向  无负权重的图  可以有环  给定一个出发点
不能有环路为负
*/

func main() {

}
