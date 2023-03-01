package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// 方式2 (推荐)
	// var p2 Person = Person{"mary", 20}
	p2 := Person{"mary", 20}
	// p2.Name = "tom"
	// p2.Age = 5
	fmt.Println(p2)

	// 方式3
	// p3 := new(Person)  和下方一样
	var p3 *Person = new(Person)
	// *p3.Name = "smith"  同下方一样
	// go的设计者 底层对写法除了
	(*p3).Name = "smith"
	p3.Name = "john"

	(*p3).Age = 30
	p3.Age = 40
	fmt.Println(*p3)

	// 方式4
	// var person *Person = &Person{"jake", 56}  直接赋值
	var person *Person = &Person{}
	// 因为 person 是一个指针 标准的访问字段的方法
	// (*person).Name = "scott"
	(*person).Name = "scott"
	person.Name = "scott!!!"

	(*person).Age = 88
	person.Age = 800
	fmt.Println(*person)

	// 指向p2
	var p4 *Person = &p2
	p4.Name = "孙悟空"
	p4.Age = 900

	fmt.Println(p2) // p2的值被修改
	fmt.Printf("p4.Name=%v p4.Age=%v\n", (*p4).Name, p4.Age)
	fmt.Printf("p2的地址=%p\n", &p2)
	fmt.Printf("p4的值是%p, p4的地址=%p\n", p4, &p4)
}
