package main

import (
	"fmt"
)

type MethodUtils struct {
	// 没有字段也可以
}

func (m MethodUtils) Print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (m MethodUtils) Print2(k int, n int) {
	// 不能使用m 会与声明绑定的结构体变量冲突 duplicate argument m
	for i := 1; i <= k; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (m MethodUtils) area(len float64, width float64) float64 {
	return len * width
}

type Calcuator struct {
	Num1 float64
	Num2 float64
}

func (cal *Calcuator) getSum() float64 {
	return cal.Num1 + cal.Num2
}

func (cal *Calcuator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = cal.Num1 + cal.Num2
	case '-':
		res = cal.Num1 - cal.Num2
	case '*':
		res = cal.Num1 * cal.Num2
	case '/':
		res = cal.Num1 / cal.Num2
	default:
		fmt.Println("输入错误")
	}
	return res
}

func main() {

	var m MethodUtils
	m.Print()

	fmt.Println()
	m.Print2(4, 5)

	fmt.Println()
	arearRes := m.area(6.0, 7.9)
	fmt.Println(arearRes)

	//
	var cal Calcuator
	cal.Num1 = 3.5
	cal.Num2 = 4.5
	fmt.Println(cal.getSum())

	res := cal.getRes('-')
	fmt.Println("res=", res)

}
