package main

import (
	"testing"
)

func TestGetSub(t *testing.T) {
	res := getSub(10, 4)
	if res != 6 {
		// fmt.Printf("AddUpper计算错误 cal count=%v 实际值=%v\n", res, 55)
		t.Fatalf("AddUpper计算错误 实际值=%v 期望值=%v\n", res, 6)
	}

	// 如果正确 输出日志
	t.Logf("getSub 计算正确!!!")
}
