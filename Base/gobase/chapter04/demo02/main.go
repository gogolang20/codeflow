package main

import (
	"fmt"
)

func test() bool {
	fmt.Println("test...")
	return true
}
func main() {
	var i int = 10

	// 短路 与&& 演示
	if i < 9 && test() {
		fmt.Println("ok...")
	}
}
