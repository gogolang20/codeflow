package main

import (
	"fmt"
	"os"
)

func main() {
	// 文件时指针类型
	// 打开文件
	file, err := os.Open("e:/test/test.txt")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("open error:", err)
		}
	}()
	// 输出文件
	fmt.Printf("file=%v", file) // file=&{0xc0420625a0}
	// 关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close error:", err)
	}
}
