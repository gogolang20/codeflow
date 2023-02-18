package main

import "container/list"

// https://lightoj.com/problem/internet-bandwidth

type Edge struct {
	from      int
	to        int
	available int
}

func NewEdge(a, b, c int) *Edge {
	return &Edge{
		from:      a,
		to:        b,
		available: c,
	}
}

type Dinic struct {
	N     int
	nexts *list.List // 放 Edge 和反向边
	edges []*Edge
	depth []int
	cur   []int
}
