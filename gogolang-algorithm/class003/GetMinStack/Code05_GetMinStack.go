package main

import (
	"container/list"
	"errors"
)

/*
实现一个特殊的栈。在基本功能的基础上。再实现返回栈中最小元素的功能
1 pop push getMin操作的时间复杂度都是 O(1) （不可以遍历，遍历是O(N)）
2 设计的栈类型可以使用现成的栈结构
    数据栈：正常操作pop push
    最小栈：同步 pop push，如果压入栈的数比最小栈栈顶小，压入，否则重复压入栈底数；同步增长
*/

type stack struct {
	stackData *list.List
	stackMin  *list.List
}

func (s *stack) push(newNum int) {
	//判断是否大于最小栈顶的数
	if s.stackMin.Len() == 0 {
		s.stackMin.PushFront(newNum)
	} else if newNum < s.stackMin.Front().Value.(int) {
		s.stackMin.PushFront(newNum)
	} else {
		newMin := s.stackMin.Front().Value.(int)
		s.stackMin.PushFront(newMin)
	}

	s.stackData.PushFront(newNum)
}
func (s *stack) pop() (int, error) {
	if s.stackData.Len() == 0 {
		return 0, errors.New("Stack is empty!")
	}
	//弹出数据栈和最小栈的数据
	s.stackMin.Remove(s.stackMin.Front())
	res := s.stackData.Front().Value.(int)
	s.stackData.Remove(s.stackData.Front())

	return res, nil
}

func (s *stack) getmin() (int, error) {
	if s.stackMin.Len() == 0 {
		return 0, errors.New("Stack is empty!")
	}
	//查询最小栈的顶，返回
	return s.stackMin.Front().Value.(int), nil
}

func main() {

}
