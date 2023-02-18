package main

import "fmt"

func printStar(N int) {
	leftUp := 0
	rightDown := N - 1
	m := make([][]byte, N)
	for i := range m {
		m[i] = make([]byte, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			m[i][j] = ' '
		}
	}
	for leftUp <= rightDown {
		set(m, leftUp, rightDown)
		leftUp += 2
		rightDown -= 2
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%c ", m[i][j])
		}
		fmt.Println()
	}
}

func set(m [][]byte, leftUp, rightDown int) {
	for col := leftUp; col <= rightDown; col++ {
		m[leftUp][col] = '*'
	}
	for row := leftUp + 1; row <= rightDown; row++ {
		m[row][rightDown] = '*'
	}
	for col := rightDown - 1; col > leftUp; col-- {
		m[rightDown][col] = '*'
	}
	for row := rightDown - 1; row > leftUp+1; row-- {
		m[row][leftUp+1] = '*'
	}
}

func main() {
	printStar(5)
}
