package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 编写一个函数 接收两个文件路径 一个是 srcFileName dstFileName
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {

	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0111)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}

	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(writer, reader)
}

func main() {

	srcFile := "e:/test/123.jpg"
	dstFile := "e:/1234.jpg"

	// 调用完成拷贝
	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Printf("拷贝err =%v", err)
	}

}
