package main

import "fmt"

func t(i int) int {
	return i * i
}

func fp(f func(int) int, num int) int {
	return f(num)
}

func main() {
	res := fp(t, 6)
	fmt.Println(res)
}
