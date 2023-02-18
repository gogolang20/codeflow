package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "e:/test/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666) //数字在window系统没用
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//即使关闭文件
	defer file.Close()
	str := "你好 上海\n"
	//写入时 使用带缓存的*writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
