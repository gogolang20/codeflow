package main

import "fmt"

func printMatrixZigZag(matrix [][]int) {
	tR := 0
	tC := 0
	dR := 0
	dC := 0
	endR := len(matrix) - 1
	endC := len(matrix[0]) - 1
	fromUp := false
	for tR != endR+1 {
		printLevel(matrix, tR, tC, dR, dC, fromUp)
		if tC == endC {
			tR = tR + 1
			tC = tC
		} else {
			tR = tR
			tC = tC + 1
		}
		if dR == endR {
			dC = dC + 1
			dR = dR
		} else {
			dC = dC
			dR = dR + 1
		}
		fromUp = !fromUp
	}
	fmt.Println()
}

func printLevel(m [][]int, tR, tC, dR, dC int, f bool) {
	if f {
		for tR != dR+1 {
			fmt.Printf("%v ", m[tR][tC])
			tR++
			tC--
		}
	} else {
		for dR != tR-1 {
			fmt.Printf("%v ", m[dR][dC])
			dR--
			dC++
		}
	}
}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	printMatrixZigZag(matrix)
}
