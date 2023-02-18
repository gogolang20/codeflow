package main

import (
	"container/list"
	"fmt"
)

/*
给你一个栈，请你逆序这个栈，
不能申请额外的数据结构，
只能使用递归函数。 如何实现?
*/
func reverse(stack *list.List) {
	if stack.Len() == 0 {
		return
	}
	i := f(stack)
	reverse(stack)
	stack.PushFront(i)
}

//将栈底元素弹出，其他位置不变
func f(stack *list.List) int {
	result := stack.Front()
	stack.Remove(result)
	if stack.Len() == 0 {
		return result.Value.(int)
	} else {
		last := f(stack)
		stack.PushFront(result.Value.(int))
		return last
	}
}
func main() {
	res := list.New()
	res.PushFront(1)
	res.PushFront(2)
	res.PushFront(3)
	fmt.Println(res.Front().Value)
	reverse(res)
	fmt.Println(res.Front().Value)
}
