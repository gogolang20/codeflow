package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "e:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666) // 数字在window系统没用
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	// 即使关闭文件
	defer file.Close()

	// 先读取原来的文件内容 显示在终端
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// 显示到终端
		fmt.Print(str)
	}

	str := "hello,上海！\n"
	// 写入时 使用带缓存的*writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
