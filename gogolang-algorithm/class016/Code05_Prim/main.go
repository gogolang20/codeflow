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
最小生成树算法之Prim

1）可以从任意节点出发来寻找最小生成树
2）某个点加入到被选取的点中后，解锁这个点出发的所有新的边
3）在所有解锁的边中选最小的边，然后看看这个边会不会形成环
4）如果会，不要当前边，继续考察剩下解锁的边中最小的边，重复3）
5）如果不会，要当前边，将该边的指向点加入到被选取的点中，重复2）
6）当所有点都被选取，最小生成树就得到了
*/

func main() {

}
