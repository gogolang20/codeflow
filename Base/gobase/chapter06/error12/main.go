package main

import (
	"errors"
	"fmt"
)

func test() {
	// 使用 defer 和 recover 除了
	defer func() {
		err := recover() // recover内置函数可以捕获异常
		if err != nil {  // 说明捕获到错误
			fmt.Println("err=", err)
			// 这里将错误发送给管理员
			fmt.Println("将错误邮件发送信息给XXX")
		}
	}() // ()表示匿名函数立即执行
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
	/*
		defer func() {
			if err := recover; err != nil {
				fmt.Println("err", err)
			}
		}()
	*/
}

// 函数去读取以配置文件init.conf的信息
// 如果文件名传入不正确 我们就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		// 读取
		return nil // nil 是没有错误的意思
	} else {
		// 返回一个自定义错误
		return errors.New("读取文件错误。。")
	}
}

func test02() {
	//
	err := readConf("config2.ini")
	if err != nil {
		// 如果读取文件发送错误 就输出这个错误 并终止程序
		panic(err)
	}

	fmt.Println("test02继续执行。。。")
}

func main() {

	// 测试
	// test()

	// fmt.Println("main()下的代码。。。")

	// 测试自定义错误
	test02()
	fmt.Println("main()下的代码。。。")
}
