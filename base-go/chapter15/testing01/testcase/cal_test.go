package main

import (
	"fmt"
	"testing"
)

// 编写一个测试用例
func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		// fmt.Printf("AddUpper计算错误 cal count=%v 实际值=%v\n", res, 55)
		t.Fatalf("AddUpper计算错误 实际值=%v 期望值=%v\n", res, 55)
	}

	// 如果正确 输出日志
	t.Logf("AddUpper计算正确")
}

func TestHello(t *testing.T) {
	fmt.Println("hello 被调用")
}
