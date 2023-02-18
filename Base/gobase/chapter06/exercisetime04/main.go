package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	// 计算test03函数运行时间
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("执行test03()耗费的时间是%v秒\n", end-start)
}
