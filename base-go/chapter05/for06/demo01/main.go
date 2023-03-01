package main

import (
	"fmt"
)

func main() {

	// 打印299句 hello world

	// for循环
	for i := 1; i <= 10; i++ {
		fmt.Println("hello world", i)
	}

	// for循环的第二种写法
	j := 1        // 循环变量初始化
	for j <= 10 { // 循环条件
		fmt.Println("hello world！！！", j)
		j++ // 循环变量迭代
	}

	// for循环的第三种写法 通常配合 break 使用
	k := 1
	for {
		if k <= 10 {

			fmt.Println("hello")
		} else {
			break // 跳出for 循环
		}
		k++
	}

	// 字符串遍历方式1 传统方式
	// var str string = "hello,world!"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c \n", str[i])
	// }

	// 使用切片
	var str string = "hello,world!中国"
	str2 := []rune(str) // 切片
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i])
	}

	// 字符串遍历方式2 for range
	str = "23tavde67上海"
	for index, val := range str {
		fmt.Printf("index=%d, val=%c \n", index, val)
	}

}
