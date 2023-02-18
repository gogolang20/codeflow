package main

import (
	"fmt"
)

type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}

func (i *integer) change() {
	*i += 1
}

type Student struct {
	Name string
	Age  int
}

// 给*Student 实现方法String()
func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]\n", stu.Name, stu.Age)
	return str
}

func main() {
	var i integer = 10
	i.print()
	i.change()
	fmt.Println("main() i=", i)

	//
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	fmt.Println(&stu)

}
