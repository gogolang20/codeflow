package main

type IndexTree2D struct {
	tree [][]int
	nums [][]int
	N    int
	M    int
}

func NewIndexTree2D(matrix [][]int) *IndexTree2D {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	tree := make([][]int, len(matrix)+1)
	for i := range tree {
		tree[i] = make([]int, len(matrix[0])+1)
	}
	nums := make([][]int, len(matrix))
	for i := range nums {
		nums[i] = make([]int, len(matrix[0]))
	}
	return &IndexTree2D{
		tree: tree,
		nums: nums,
		N:    len(matrix),
		M:    len(matrix[0]),
	}
}

func (it2D *IndexTree2D) sum(row, col int) int {
	sum := 0
	for i := row + 1; i > 0; i -= i & (-i) {
		for j := col + 1; j > 0; j -= j & (-j) {
			sum += it2D.tree[i][j]
		}
	}
	return sum
}

func (it2D *IndexTree2D) update(row, col, val int) {
	if it2D.N == 0 || it2D.M == 0 {
		return
	}
	add := val - it2D.nums[row][col]
	it2D.nums[row][col] = val
	for i := row + 1; i <= it2D.N; i += i & (-i) {
		for j := col + 1; j <= it2D.M; j += j & (-j) {
			it2D.tree[i][j] += add
		}
	}
}

func (it2D *IndexTree2D) sumRegion(row1, col1, row2, col2 int) int {
	if it2D.N == 0 || it2D.M == 0 {
		return 0
	}
	return it2D.sum(row2, col2) + it2D.sum(row1-1, col1-1) - it2D.sum(row1-1, col2) - it2D.sum(row2, col1-1)
}
