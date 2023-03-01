package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "e:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666) // 数字在window系统没用
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()

	str := "abc,ENGLISH!\n"
	// 写入时 使用带缓存的*writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}
