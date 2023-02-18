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
最小生成树算法之Kruskal

1）总是从权值最小的边开始考虑，依次考察权值依次变大的边
2）当前的边要么进入最小生成树的集合，要么丢弃
3）如果当前的边进入最小生成树的集合中不会形成环，就要当前边
4）如果当前的边进入最小生成树的集合中会形成环，就不要当前边
5）考察完所有边之后，最小生成树的集合也得到了

使用并查集 + 贪心
*/
func (eg Edge)Less(i,j int) {

}
func kruskalMST(graph Graph) map[*Edge]struct{} {
	result := make(map[*Edge]struct{})
	// 从小的边到大的边，依次弹出，小根堆！
	return result
}

func main() {

}
