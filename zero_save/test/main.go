package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3}
	x := a[:1]
	y := a[2:]
	fmt.Println(x, cap(x), len(x)) // [0] 4 1
	fmt.Println(y, cap(y), len(y)) // [2 3] 2 2

	x = append(x, y...)            // 0 2 3
	fmt.Println(a, x)              // [0 2 3 3] [0 2 3]
	fmt.Println(y, cap(y), len(y)) // [3 3] 2 2

	x = append(x, y...)
	fmt.Println(a, x) // [0 2 3 3] [0 2 3 3 3]
}
