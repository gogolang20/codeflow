package main

import "fmt"

//leetcode 305题 Number of Islands II

//方法一 递归
func numIslands3(board [][]byte) int {
	islands := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '1' {
				islands++
				infect(board, i, j)
			}
		}
	}
	return islands
}

// 从(i,j)这个位置出发，把所有练成一片的'1'字符，变成0
func infect(board [][]byte, i, j int) {
	if i < 0 || i == len(board) || j < 0 || j == len(board[0]) || board[i][j] != '1' {
		return
	}
	board[i][j] = 99 //变成不是 '1' 的 ASCII
	infect(board, i-1, j)
	infect(board, i+1, j)
	infect(board, i, j-1)
	infect(board, i, j+1)
}

//方法二 并查集
type UnionFind2 struct {
	parent []int
	size   []int
	help   []int
	colume int
	sets   int
}

func NewUnionFind2(board [][]byte) *UnionFind2 {
	row := len(board)
	colume := len(board[0])
	length := row * colume
	sets := 0
	parent := make([]int, length)
	size := make([]int, length)
	help := make([]int, length)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '1' {
				in := index(i, j, colume)
				parent[in] = in
				size[i] = 1
				sets++
			}
		}
	}
	return &UnionFind2{
		parent: parent,
		size:   size,
		help:   help,
		colume: colume,
		sets:   sets,
	}
}

//[][]byte 转换成一位数组表示每个元素
//index = row * colume + col // (row,col) -> i
func index(row, col int, colume int) int {
	return row*colume + col
}

func (unf *UnionFind2) find() {
	/*
	 */
}

func (unf *UnionFind2) union() {
	/*
		public void union(int r1, int c1, int r2, int c2) {
			int i1 = index(r1, c1);
			int i2 = index(r2, c2);
			int f1 = find(i1);
			int f2 = find(i2);
			if (f1 != f2) {
				if (size[f1] >= size[f2]) {
					size[f1] += size[f2];
					parent[f2] = f1;
				} else {
					size[f2] += size[f1];
					parent[f1] = f2;
				}
				sets--;
			}
		}
	*/
}

func (unf *UnionFind2) setss() int {
	return unf.sets
}

func main() {
	var b byte = '1'
	fmt.Printf("%c %v\n", b, b)

	test1 := struct{}{}
	test2 := struct{}{}
	fmt.Println(test1 == test2)   //true
	fmt.Println(&test1 == &test2) //false
}
