package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 保存统计的结果
type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {

	fileName := "e:/test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open err= %v", err)
	}
	defer file.Close()

	var count CharCount
	reader := bufio.NewReader(file)

	// 开始循环读取内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// !!!  str = []rune(str)
		// 遍历 str 统计
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough // 穿透
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}

	fmt.Printf("字符的个数为=%v 数字的个数为=%v 空格的个数为=%v 其他的个数为=%v",
		count.ChCount, count.SpaceCount, count.NumCount, count.OtherCount)

}
