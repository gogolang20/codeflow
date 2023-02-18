package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 文件时指针类型
	// 打开文件
	file, err := os.Open("e:/test/test.txt")
	if err != nil {
		fmt.Println("open error:", err)
	}
	defer file.Close() // 即使关闭file
	// 输出文件
	// 创建一个*Reader 带缓存的
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束一次
		if err == io.EOF {                  // io.EOF表示文件的末尾
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束...")
}
