package main

import (
	"fmt"
)

type Person struct {
	Name string
}

// speak的方法 与Person 类型绑定
func (p Person) speak() {
	fmt.Println(p.Name, "是一个goodman")
}

// 从1加到1000的方法
func (p Person) jisuan() {
	res := 0
	for i := 0; i <= 1000; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算结果是:", res)
}

// 接收一个数计算
func (p Person) jisuan2(n int) {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算结果是:", res)
}

// 计算两个数的和 并返回值
func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 给 Person 类型绑定一个方法
func (person Person) test() {
	person.Name = "jcak"
	fmt.Println("test() name=", person.Name) // 输出 jack
}

func main() {
	var p Person
	p.Name = "tom"
	p.test()                            // 调用方法
	fmt.Println("main() name=", p.Name) // 输出 tom

	// speak和test 两个方法都和 Person 这个结构体类型绑定了
	p.speak()
	fmt.Println(p.Name)
	p.jisuan()
	p.jisuan2(10)

	n1 := 10
	n2 := 20
	res := p.getSum(n1, n2)
	fmt.Println("!res=", res)
}
