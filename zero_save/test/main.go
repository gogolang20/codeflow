package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3}
	x := a[:1] // 0
	y := a[2:] // 2 3
	fmt.Println(x)
	fmt.Println(y)

	x = append(x, y...) // 0 2 3
	fmt.Println(a)

	x = append(x, y...) // x = 0 2 3 3 3
	fmt.Println(a, x)   //
}
