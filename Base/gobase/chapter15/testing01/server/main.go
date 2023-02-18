package main

import (
	"fmt"
)

func main() {
	res := addUpper(10)
	if res == 55 {
		fmt.Printf("计算正确 cal count=%v 实际值=%v\n", res, 55)
	} else {
		fmt.Printf("计算错误 cal count=%v 实际值=%v\n", res, 55)
	}
}
