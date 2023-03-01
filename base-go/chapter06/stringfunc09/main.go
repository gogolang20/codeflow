package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// 统计字符串的长度 按字节返回
	// 一个汉字占3个字节
	str := "hello 北"
	fmt.Println("str len=", len(str))

	str2 := "hello 上海"
	// 切片
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	// 字符串转整数 strconv内的Atoi函数
	// n, err := strconv.Atoi("12")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换结果是", n)
	}

	// 整数转字符串
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v, str=%T\n", str, str)

	// 字符串转byte切片
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	// byte转成字符串
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)

	// 10进制转 2 8 16进制 返回对应的字符串
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是%v\n", str)
	str = strconv.FormatInt(123, 8)
	fmt.Printf("123对应的八进制是%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的十六进制是%v\n", str)

	// 查找子串是否在指定的字符串中 strings 包的函数
	b := strings.Contains("seafood", "tim")
	fmt.Printf("b=%v\n", b)

	// 统计一个字符串中有几个指定的字串 strings 包的函数
	num := strings.Count("cehesse", "se")
	fmt.Printf("num=%v\n", num)

	// 不区分大小写的字符串比较（==是区分字母大小写的）
	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v\n", b) // true

	fmt.Println("结果", "abc" == "Abc") // false

	// 返回字串在字符串第一次出现的 index 值，如果没有 返回-1
	index := strings.Index("NLT_abc", "abc") // 4
	fmt.Printf("index=%v\n", index)

	// 返回字串在字符串最后一次出现的 index 值，如果没有 返回-1
	index = strings.LastIndex("go golang", "go") // 3
	fmt.Printf("!index=%v\n", index)

	// 将指定的字串替换成另一个字串 如果n=-1 表示替换全部
	str2 = "go go golang"
	str = strings.Replace(str2, "go", "北京", -1)
	fmt.Printf("!str=%v str2=%v\n", str, str2)

	// 按照某个指定的字符，为分割标识，将一个字符串拆分成字符串数组
	strArr := strings.Split("hello,world,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}
	fmt.Printf("strArr=%v\n", strArr)

	// 将字符串的字母进行大小写的转换
	str = "go Lang Hello"
	str = strings.ToLower(str)
	fmt.Printf("str=%v\n", str)
	str = strings.ToUpper(str)
	fmt.Printf("str=%v\n", str)

	// 16:将字符串左右两边的空格去掉
	str = strings.TrimSpace("  hello , wor ld ,  ok  ") // hello , wor ld ,  ok
	fmt.Printf("str=%q\n", str)

	// 17:将字符串左右两边指定的字符去掉
	str = strings.Trim("!  hello , wor ! ld ,  ok  !", " !")
	fmt.Printf("str=%q\n", str)

	// 18:将字符串左边指定的字符去掉
	str = strings.TrimLeft("!  hello , wor ! ld ,  ok  !", " !")
	fmt.Printf("str=%q\n", str)

	// 19:将字符串右边指定的字符去掉
	str = strings.TrimRight("!  hello , wor ! ld ,  ok  !", " !")
	fmt.Printf("str=%q\n", str)

	// 20:判断字符串是否以指定的字符串开头
	b = strings.HasPrefix("ft go see see col", "ft")
	fmt.Printf("b=%v\n", b)

	// 21:判断字符串是否以指定的字符串结束
	b = strings.HasSuffix("Ft go see see col", "ft")
	fmt.Printf("!b=%v\n", b)

}
