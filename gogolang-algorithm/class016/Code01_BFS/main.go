package main

import (
	"container/list"
	"fmt"
)

/*
图的表示方法
1）邻接表法
2）邻接矩阵法
*/

/*
图
1）由点的集合和边的集合构成
2）虽然存在有向图和无向图的概念，但实际上都可以用有向图来表达
3）边上可能带有权值
*/

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

//准备一个set 去重
var set map[*Node]struct{}

//宽度优先遍历：使用队列
func bfs(node *Node) {
	if node == nil {
		return
	}
	queue := list.New()
	queue.PushBack(node)
	set[node] = struct{}{}
	for queue.Len() > 0 {
		cur := queue.Front()                 //弹出
		queue.Remove(cur)                    //真正的移出 queue
		fmt.Println(cur.Value.(*Node).value) //打印
		for _, v := range cur.Value.(*Node).nexts {
			if _, ok := set[v]; !ok { //集合没有的节点加入队列
				set[v] = struct{}{}
				queue.PushBack(v)
			}
		}
	}
}

//深度优先遍历
func dfs(node *Node) {
	if node == nil {
		return
	}
	stack := list.New()
	stack.PushFront(node)   //压入
	set[node] = struct{}{}  //加入集合
	fmt.Println(node.value) //打印
	for stack.Len() > 0 {
		cur := stack.Front()
		stack.Remove(cur)
		for _, v := range cur.Value.(*Node).nexts {
			if _, ok := set[v]; !ok {
				stack.PushFront(cur.Value.(*Node)) //cur 重新压入
				stack.PushFront(v)                 //压入
				set[v] = struct{}{}                //加入集合
				fmt.Println(v.value)               //打印
				break
			}
		}
	}
}

// matrix 所有的边
// N*3 的矩阵
// [weight, from节点上面的值，to节点上面的值]
// [ 5 , 0 , 7]
// [ 3 , 0,  1]
func createGraph(matrix [][]int) *Graph {
	graph := NewGraph()
	for i := 0; i < len(matrix); i++ {
		// 拿到每一条边， matrix[i]
		weight := matrix[i][0]
		from := matrix[i][1]
		to := matrix[i][2]
		if _, ok := graph.nodes[from]; !ok {
			graph.nodes[from] = NewNode(from)
		}
		if _, ok := graph.nodes[to]; !ok {
			graph.nodes[from] = NewNode(to)
		}
		fromNode := graph.nodes[from]
		toNode := graph.nodes[to]
		newEdge := &Edge{weight: weight, from: fromNode, to: toNode}
		fromNode.nexts = append(fromNode.nexts, toNode)
		fromNode.out++
		toNode.in++
		fromNode.edges = append(fromNode.edges, newEdge)
		graph.edges[newEdge] = struct{}{}
	}
	return graph
}
