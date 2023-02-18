package main

import (
	"container/heap"
	"fmt"
	"gogolang-algorithm/utils"
	"math"
	"sort"
)

type Line struct {
	start int
	end   int
}

type Lines []*Line

func (l Lines) Len() int {
	return len(l)
}
func (l Lines) Less(i, j int) bool {
	return l[i].start < l[j].start
}
func (l Lines) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (h *Lines) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Line))
}
func (h *Lines) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxCover1(lines [][]int) int {
	min := math.MaxInt
	max := math.MinInt
	for i := 0; i < len(lines); i++ {
		min = utils.Min(min, lines[i][0])
		max = utils.Max(max, lines[i][1])
	}
	cover := 0
	for p := float64(min) + 0.5; p < float64(max); p += 1 {
		cur := 0
		for i := 0; i < len(lines); i++ {
			if float64(lines[i][0]) < p && float64(lines[i][1]) > p {
				cur++
			}
		}
		cover = utils.Max(cover, cur)
	}
	return cover
}

/*
代码有待完善
*/
func maxCover2(m [][]int) int {
	liness := make(Lines, len(m))
	for i := range m {
		liness[i] = &Line{start: m[i][0], end: m[i][1]}
	}
	sort.Sort(liness)
	// 小根堆，每一条线段的结尾数值，使用默认的
	heap.Init(&liness)
	max := 0
	for i := range liness {
		// lines[i] -> cur  在黑盒中，把<=cur.start 东西都弹出
		for len(liness) > 0 && liness[0].start <= liness[i].start {
			heap.Pop(&liness)
		}
		heap.Push(&liness, liness[i].end)
		max = utils.Max(max, len(liness[:i+1]))
	}
	return max
}

func main() {
	lines := make([][]int, 6)
	for i := range lines {
		lines[i] = make([]int, 2)
	}
	lines[0] = []int{4, 9}
	lines[1] = []int{1, 4}
	lines[2] = []int{7, 15}
	lines[3] = []int{2, 4}
	lines[4] = []int{4, 6}
	lines[5] = []int{3, 7}

	fmt.Println(maxCover1(lines)) // 3
	fmt.Println(maxCover2(lines))
}
