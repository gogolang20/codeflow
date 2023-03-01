package main

import (
	"fmt"
)

type AInterface interface {
	Say()
}

type Stu struct {
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say() i=", i)
}

// 一个自定义类型可以实现多个接口
type BInterface interface {
	Hello()
}

type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Monster Hello()")
}

func (m Monster) Say() {
	fmt.Println("Monster Say()")
}

func main() {

	var stu Stu // 结构体变量实现了 Say() 也就实现了 AInterface 接口
	var a AInterface = stu
	a.Say()

	var i integer = 10
	var b AInterface = i
	b.Say()

	// Monster 实现了 AInterface 和BInterface 两个接口
	var monster Monster
	var a2 AInterface = monster
	var b2 BInterface = monster
	a2.Say()
	b2.Hello()

}
