package main

import (
	"fmt"
)

func main() {

	// 演示goto
	fmt.Println("1")
	goto label1
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
label1:
	fmt.Println("5")
	fmt.Println("6")
	fmt.Println("7")
	fmt.Println("8")

	// 演示return
	var n int = 30
	fmt.Println("1")
	if n > 10 {
		return
	}
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	fmt.Println("6")
	fmt.Println("7")
	fmt.Println("8")
}
