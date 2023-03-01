package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 打开文件 附带方式
	filePath := "e:/test/abc.txt"
	// os包中的 OpenFile 函数 传入三个参数
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	// 即使关闭文件
	defer file.Close()
	str := "hello,Golang!必胜\n" // \n\r换行两次  //\n是换行一次
	// 带缓存写入
	// 写入时 使用带缓存的*writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 读完之后 必须要添加
	// Flush 方法
	writer.Flush()
	fmt.Println("写入完成！！！")
}
