package main

/*
https://leetcode.cn/problems/falling-squares/
提交失败 ！！！
*/

type SegmentTree struct {
	max    []int
	change []int
	update []bool
}

func NewSegmentTree(size int) *SegmentTree {
	N := size + 1
	max := make([]int, N<<2)
	change := make([]int, N<<2)
	update := make([]bool, N<<2)
	return &SegmentTree{
		max:    max,
		change: change,
		update: update,
	}
}

func (st *SegmentTree) pushUp(root int) {
	st.max[root] = Max(st.max[root<<1], st.max[root<<1|1]) // i = i*2 + (i*2+1)
}

func (st *SegmentTree) pushDown(rt, ln, rn int) { // 往下分发任务
	if st.update[rt] {
		st.update[rt<<1] = true
		st.update[rt<<1|1] = true
		st.change[rt<<1] = st.change[rt]
		st.change[rt<<1|1] = st.change[rt]
		st.max[rt<<1] = st.change[rt]
		st.max[rt<<1|1] = st.change[rt]
		st.update[rt] = false
	}
}

func (st *SegmentTree) Update(L, R, C, l, r, rt int) {
	if L <= l && r <= R {
		st.update[rt] = true
		st.change[rt] = C
		st.max[rt] = C
		return
	}
	mid := (l + r) >> 1
	st.pushDown(rt, mid-l+1, r-mid)
	if L <= mid {
		st.Update(L, R, C, l, mid, rt<<1)
	}
	if R > mid {
		st.Update(L, R, C, mid+1, r, rt<<1|1)
	}
	st.pushUp(rt)
}

func (st *SegmentTree) Query(L, R, l, r, rt int) int {
	if L <= l && r <= R {
		return st.max[rt]
	}
	mid := (l + r) >> 1
	st.pushDown(rt, mid-l+1, r-mid)
	left, right := 0, 0
	if L <= mid {
		left = st.Query(L, R, l, mid, rt<<1)
	}
	if R > mid {
		right = st.Query(L, R, mid+1, r, rt<<1|1)
	}
	return Max(left, right)
}

func Index(positions [][]int) map[int]int {
	pos := make(map[int]struct{})
	for _, arr := range positions {
		pos[arr[0]] = struct{}{}
		pos[arr[0]+arr[1]-1] = struct{}{}
	}
	maps := make(map[int]int)
	count := 0
	for key, _ := range pos {
		maps[key] = count + 1
		count++
	}
	return maps
}

func fallingSquares(positions [][]int) []int {
	maps := Index(positions)
	N := len(maps)
	segmentTree := NewSegmentTree(N)
	max := 0
	res := make([]int, 0)
	// 每落一个正方形，收集一下，所有东西组成的图像，最高高度是什么
	for _, arr := range positions {
		L := maps[arr[0]]
		R := maps[arr[0]+arr[1]-1]
		height := segmentTree.Query(L, R, 1, N, 1) + arr[1]
		max = Max(max, height)
		res = append(res, max)
		segmentTree.Update(L, R, height, 1, N, 1)
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
