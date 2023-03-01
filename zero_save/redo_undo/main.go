package main

import (
	"container/list"
	"fmt"
	"strings"
)

func Redo(str string) string {
	arr := strings.Split(str, " ")
	stack := list.New()
	redoStack := list.New()
	fmt.Println(arr)
	for _, val := range arr {
		// 遇到普通词就压入stack，并清空redo栈，因为此时写入了一个新词，再往前的词已经找不回来了
		// 遇到undo就从stack中弹栈至redo
		// 遇到redo就从redo中弹栈至stack
		if val == "undo" {
			if stack.Len() > 0 {
				redoStack.PushFront(stack.Front().Value.(string))
				stack.Remove(stack.Front())
			}
		} else if val == "redo" {
			if redoStack.Len() > 0 {
				stack.PushFront(redoStack.Front().Value.(string))
				redoStack.Remove(redoStack.Front())
			}
		} else {
			if redoStack.Len() > 0 {
				redoStack.Remove(redoStack.Front())
			}
			stack.PushFront(val)
		}
	}

	newArr := make([]string, 0)
	for stack.Len() > 0 {
		if _, ok := stack.Back().Value.(string); !ok {
			break
		}
		newArr = append(newArr, stack.Back().Value.(string))
		stack.Remove(stack.Back())
	}
	return strings.Join(newArr, " ")
}

func main() {
	fmt.Printf("res: [%s]\n", Redo("hello undo redo world."))
}

// https://ac.nowcoder.com/questionTerminal/46badc29891b4294a3b9cc235a96631a?orderByHotValue=1&page=1&onlyReference=false
