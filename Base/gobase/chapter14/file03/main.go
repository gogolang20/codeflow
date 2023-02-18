package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 使用ioutil.ReadFile 一次性的将文件读取到位
	// 文件不可以太大，适合小文件
	file := "e:/test/test.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err= %v", err)
	}
	fmt.Printf("%v", content) // 输出的是代码
	fmt.Println()
	fmt.Printf("%v", string(content)) // 输出的内容 转换成字符串
}
