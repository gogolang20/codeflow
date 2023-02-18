package main

import "fmt"

/*
线段树

1，一种支持范围整体修改和范围整体查询的数据结构

2，解决的问题范畴：
大范围信息可以只由左、右两侧信息加工出，
而不必遍历左右两个子范围的具体状况
*/

/*
线段树实例一

给定一个数组arr，用户希望你实现如下三个方法
1）void add(int L, int R, int V) :  让数组arr[L…R]上每个数都加上V
2）void update(int L, int R, int V) :  让数组arr[L…R]上每个数都变成V
3）int sum(int L, int R) :让返回arr[L…R]这个范围整体的累加和
怎么让这三个方法，时间复杂度都是O(logN)
*/

/*
线段树实例二

想象一下标准的俄罗斯方块游戏，X轴是积木最终下落到底的轴线
下面是这个游戏的简化版：
1）只会下落正方形积木
2）[a,b] -> 代表一个边长为b的正方形积木，积木左边缘沿着X = a这条线从上方掉落
3）认为整个X轴都可能接住积木，也就是说简化版游戏是没有整体的左右边界的
4）没有整体的左右边界，所以简化版游戏不会消除积木，因为不会有哪一层被填满。

给定一个N*2的二维数组matrix，可以代表N个积木依次掉落，
返回每一次掉落之后的最大高度
*/



type SegmentTree struct {
	MAXN   int
	arr    []int  // arr[]为原序列的信息从0开始，但在arr里是从1开始的
	sum    []int  // sum[]模拟线段树维护区间和
	lazy   []int  // lazy[]为累加和懒惰标记
	change []int  // change[]为更新的值
	update []bool // update[]为更新慵懒标记
}

func NewSegmentTree(origin []int) *SegmentTree {
	MAXN := len(origin) + 1
	arr := make([]int, MAXN) // arr[0] 不用 从1开始使用
	for i := 1; i < MAXN; i++ {
		arr[i] = origin[i-1]
	}
	sum := make([]int, MAXN<<2)     // 用来支持脑补概念中，某一个范围的累加和信息
	lazy := make([]int, MAXN<<2)    // 用来支持脑补概念中，某一个范围沒有往下傳遞的纍加任務
	change := make([]int, MAXN<<2)  // 用来支持脑补概念中，某一个范围有没有更新操作的任务
	update := make([]bool, MAXN<<2) // 用来支持脑补概念中，某一个范围更新任务，更新成了什么
	return &SegmentTree{
		MAXN:   MAXN,
		arr:    arr,
		sum:    sum,
		lazy:   lazy,
		change: change,
		update: update,
	}
}

func (st *SegmentTree) pushUp(root int) {
	st.sum[root] = st.sum[root<<1] + st.sum[root<<1|1] // i = i*2 + (i*2+1)
}

// 之前的，所有懒增加，和懒更新，从父范围，发给左右两个子范围
// 分发策略是什么
// ln表示左子树元素结点个数，rn表示右子树结点个数
func (st *SegmentTree) pushDown(rt, ln, rn int) { // 往下分发任务
	if st.update[rt] {
		st.update[rt<<1] = true
		st.update[rt<<1|1] = true
		st.change[rt<<1] = st.change[rt]
		st.change[rt<<1|1] = st.change[rt]
		st.lazy[rt<<1] = 0
		st.lazy[rt<<1|1] = 0
		st.sum[rt<<1] = st.change[rt] * ln
		st.sum[rt<<1|1] = st.change[rt] * rn
		st.update[rt] = false
	}
	if st.lazy[rt] != 0 {
		st.lazy[rt<<1] += st.lazy[rt]
		st.sum[rt<<1] += st.lazy[rt] * ln
		st.lazy[rt<<1|1] += st.lazy[rt]
		st.sum[rt<<1|1] += st.lazy[rt] * rn
		st.lazy[rt] = 0
	}
}

// 在初始化阶段，先把sum数组，填好
// 在arr[l~r]范围上，去build，1~N，
// rt : 这个范围在sum中的下标
func (st *SegmentTree) Build(l, r, rt int) {
	if l == r { // 绝对是叶节点
		st.sum[rt] = st.arr[l]
		return
	}
	mid := (l + r) >> 1
	st.Build(l, mid, rt<<1)
	st.Build(mid+1, r, rt<<1|1)
	st.pushUp(rt)
}

// L~R  所有的值变成C
// l~r  rt
func (st *SegmentTree) Update(L, R, C, l, r, rt int) {
	if L <= l && r <= R {
		st.update[rt] = true
		st.change[rt] = C
		st.sum[rt] = C * (r - l + 1)
		st.lazy[rt] = 0
		return
	}
	// 当前任务躲不掉，无法懒更新，要往下发
	mid := (l + r) >> 1
	st.pushDown(rt, mid-l+1, r-mid)
	if L <= mid {
		st.Update(L, R, C, l, mid, rt<<1)
	}
	if R > mid {
		st.Update(L, R, C, mid+1, r, rt<<1|1)
	}
	st.pushUp(rt) // 调整累加和
}

// L~R, C 任务！
// rt，l~r
func (st *SegmentTree) Add(L, R, C, l, r, rt int) {
	// 任务如果把此时的范围全包了！
	if L <= l && r <= R {
		st.sum[rt] += C * (r - l + 1)
		st.lazy[rt] += C
		return
	}
	// 任务没有把你全包！
	// l  r  mid = (l+r)/2
	mid := (l + r) >> 1
	st.pushDown(rt, mid-l+1, r-mid)
	// L~R
	if L <= mid {
		st.Add(L, R, C, l, mid, rt<<1)
	}
	if R > mid {
		st.Add(L, R, C, mid+1, r, rt<<1|1)
	}
	st.pushUp(rt)
}

// 1~6 累加和是多少？ 1~8 rt
func (st *SegmentTree) Query(L, R, l, r, rt int) int {
	if L <= l && r <= R {
		return st.sum[rt]
	}
	mid := (l + r) >> 1
	st.pushDown(rt, mid-l+1, r-mid)
	ans := 0
	if L <= mid {
		ans += st.Query(L, R, l, mid, rt<<1)
	}
	if R > mid {
		ans += st.Query(L, R, mid+1, r, rt<<1|1)
	}
	return ans
}

func main() {
	arr := []int{5, 2, 3, 6, 8, 9, 1, 7}
	st := NewSegmentTree(arr)
	st.Build(1, len(arr), 1)
	fmt.Println(st.sum)

	//st.Add(3, 4, 6, 1, len(arr), 1)
	//
	//st.Update(3, 4, 6, 1, len(arr), 1)
	//fmt.Println(st.change)

	fmt.Println(st.Query(3, 3, 1, len(arr), 1))
}
