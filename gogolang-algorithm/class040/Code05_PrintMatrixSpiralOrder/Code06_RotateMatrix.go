package main

import "fmt"

func rotate(matrix [][]int) {
	a := 0
	b := 0
	c := len(matrix) - 1
	d := len(matrix[0]) - 1
	for a < c {
		rotateEdge(matrix, a, b, c, d)
		a++
		b++
		c--
		d--
	}
}

func rotateEdge(m [][]int, a, b, c, d int) {
	tmp := 0
	for i := 0; i < d-b; i++ {
		tmp = m[a][b+i]
		m[a][b+i] = m[c-i][b]
		m[c-i][b] = m[c][d-i]
		m[c][d-i] = m[a+i][d]
		m[a+i][d] = tmp
	}
}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	printMatrix(matrix)
	rotate(matrix)

	fmt.Println()
	printMatrix(matrix)
}

func printMatrix(matrix [][]int) {
	for i := 0; i != len(matrix); i++ {
		for j := 0; j != len(matrix[0]); j++ {
			fmt.Printf("%v ", matrix[i][j])
		}
		fmt.Println()
	}
}
