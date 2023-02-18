package main

import "math"

/*
派对的最大快乐值
员工信息的定义如下:
class Employee {
    public int happy; // 这名员工可以带来的快乐值
    List<Employee> subordinates; // 这名员工有哪些直接下级
}
*/

type Node struct {
	happy int
	next  []*Node
}

func NewNode(happy int, list []*Node) *Node {
	return &Node{
		happy: happy,
		next:  list,
	}
}

type Info struct {
	no  int //x 不来
	yes int //x 来
}

func NewInfo(no, yes int) *Info {
	return &Info{
		no:  no,
		yes: yes,
	}
}

func maxHappy2(head *Node) int {
	allInfo := process(head)
	return int(math.Max(float64(allInfo.no), float64(allInfo.yes)))
}

func process(x *Node) *Info {
	if x == nil {
		return NewInfo(0, 0)
	}
	no, yes := 0, x.happy
	for index := range x.next {
		nextInfo := process(x.next[index])
		no += int(math.Max(float64(nextInfo.no), float64(nextInfo.yes))) //x 不来
		yes += nextInfo.no                                               //x 来
	}
	return NewInfo(no, yes)
}

func main() {

}
