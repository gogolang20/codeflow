package main

/*
你有无限的1*2的砖块，要铺满M*N的区域，
不同的铺法有多少种?
*/

func ways1(N, M int) int {
	if N < 1 || M < 1 || ((N*M)&1) != 0 {
		return 0
	}
	if N == 1 || M == 1 {
		return 1
	}
	pre := make([]int, M) // pre代表-1行的状况
	for i := 0; i < len(pre); i++ {
		pre[i] = 1
	}
	return process(pre, 0, N)
}

// pre 表示level-1行的状态
// level表示，正在level行做决定
// N 表示一共有多少行 固定的
// level-2行及其之上所有行，都摆满砖了
// level做决定，让所有区域都满，方法数返回
func process(pre []int, level, N int) int {
	if level == N { // base case
		for i := 0; i < len(pre); i++ {
			if pre[i] == 0 {
				return 0
			}
		}
		return 1
	}
	// 没到终止行，可以选择在当前的level行摆瓷砖
	op := getOp(pre)
	return dfs(op, 0, level, N)
}

// op[i] == 0 可以考虑摆砖
// op[i] == 1 只能竖着向上
func dfs(op []int, col, level, N int) int {
	// 在列上自由发挥，玩深度优先遍历，当col来到终止列，i行的决定做完了
	// 轮到i+1行，做决定
	if col == len(op) {
		return process(op, level+1, N)
	}
	ans := 0
	// col位置不横摆
	ans += dfs(op, col+1, level, N) // col位置上不摆横转
	// col位置横摆, 向右
	if col+1 < len(op) && op[col] == 0 && op[col+1] == 0 {
		op[col] = 1
		op[col+1] = 1
		ans += dfs(op, col+2, level, N)
		op[col] = 0
		op[col+1] = 0
	}
	return ans
}

func getOp(pre []int) []int {
	cur := make([]int, len(pre))
	for i := 0; i < len(pre); i++ {
		cur[i] = pre[i] ^ 1
	}
	return cur
}
