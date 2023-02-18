package main

import "fmt"

func in(a int) (b int) {
	defer func() {
		a++
		b++
		fmt.Println("defer", a)
	}()
	a++
	fmt.Println("in", a)
	return a
}

func main() {
	var a, b int
	b = in(a)
	fmt.Println("main", a, b)
}
