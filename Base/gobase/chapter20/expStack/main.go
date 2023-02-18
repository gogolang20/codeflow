package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 使用数组模拟一个栈的使用，栈是先入后出！！！
type Stack struct {
	MaxTop int     // 表示栈最大的可以存放个数
	Top    int     // 表示栈顶，因为栈底是不变的
	arr    [20]int // 数组模拟栈
}

// 入栈
func (stack *Stack) Push(val int) (err error) {
	// 先判断栈是否满了
	if stack.Top == stack.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	stack.Top++
	// 放入数据
	stack.arr[stack.Top] = val
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

// 判断字符是不是运算符（+ - * /） ASCII 码判断
func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

// 运算的方法
func (this *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误")
	}
	return res
}

// 返回运算符的优先级【程序员定义】（*/ 优先级1  +- 优先级0）
func (htis *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

func main() {
	// 数栈
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	// 符号栈
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	exp := "30+20*3-10"
	// 定义一个 index，帮助扫描 exp
	index := 0
	// 为了配合运算 定义需要的遍历
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""

	for {
		// 增加一个判断数 长度的逻辑
		ch := exp[index : index+1] // 切片，一个一个取 是个字符串
		// ch ==> "+"  ==> 43
		temp := int([]byte(ch)[0])  // 字符对应的ASCII码
		if operStack.IsOper(temp) { // 说明是符号
			// 两个逻辑
			// 如果 operStack 是一个空栈 直接入栈
			if operStack.Top == -1 { // 空栈
				operStack.Push(temp)
			} else {
				if operStack.Priority(operStack.arr[operStack.Top]) >=
					operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)
					// 将计算的结果重新入数栈
					numStack.Push(result)
					// 当前的符号压入符号栈
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}
		} else { // 说明是数
			// 处理多位数的思路
			// 1 定义一个变量 keepNum string
			keepNum += ch
			// 2 每次向index 后面字符测试一下

			// 如果已经到表达式最后
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				// 向 index 后面是不是运算符 【index】
				if operStack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
		}
		// 继续扫描
		// 先判断 index 是否已经扫描到计算表达式的最后
		if index+1 == len(exp) {
			break
		}
		index++
	}
	// 如果扫描表达式完毕，依次从符号栈取出符号，然后从数栈取出两个数
	// 运算的结果，入数栈，直到符号栈为空
	for {
		if operStack.Top == -1 {
			break // 退出条件
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		// 将计算的结果重新入数栈
		numStack.Push(result)

	}
	// 如果算法没有问题 表示式也是正确的  则结果就是numStack 最后数
	res, _ := numStack.Pop()
	fmt.Printf("表达式%s=%v", exp, res)
}
