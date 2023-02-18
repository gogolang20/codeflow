package main

import (
	"errors"
	"fmt"
)

// 使用数组模拟一个栈的使用，栈是先入后出！！！
type Stack struct {
	MaxTop int    // 表示栈最大的可以存放个数
	Top    int    // 表示栈顶，因为栈底是不变的
	arr    [5]int // 数组模拟栈
}

// 入栈
func (this *Stack) Push(val int) (err error) {
	// 先判断栈是否满了
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	// 放入数据
	this.arr[this.Top] = val
	return
}

// 出栈
func (this *Stack) Pop() (val int, err error) {
	// 判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty!")
		return 0, errors.New("stack empty!")
	}
	// 先取值 再 this.Top--
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

// 遍历栈，主要从栈顶开始遍历
func (this *Stack) List() {
	// 先判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack enpty")
		return
	}
	fmt.Println("栈的情况如下:")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("artr[%d]=%d\n", i, this.arr[i])
	}
}

func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1, // 表示栈为空
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	// 显示
	stack.List()
	val, _ := stack.Pop()
	fmt.Println("出栈val=", val)
	// 显示
	stack.List()
}
