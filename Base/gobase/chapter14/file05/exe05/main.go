package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// 将一个文件的内容写入另一个文件中
	file1Path := "e:/test.txt"
	file2Path := "e:/test2.txt"

	// 先读取
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("read file err=%v", err)
		return
	}

	err = ioutil.WriteFile(file2Path, data, 0333)
	if err != nil {
		fmt.Printf("write file error=%v", err)
	}

}
