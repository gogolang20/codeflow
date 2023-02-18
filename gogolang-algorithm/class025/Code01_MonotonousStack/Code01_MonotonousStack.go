package main

import "container/list"

/*
一种特别设计的栈结构，为了解决如下的问题：

给定一个可能含有重复值的数组arr，i位置的数一定存在如下两个信息
1）arr[i]的左侧离i最近并且小于(或者大于)arr[i]的数在哪？
2）arr[i]的右侧离i最近并且小于(或者大于)arr[i]的数在哪？
如果想得到arr中所有位置的两个信息，怎么能让得到信息的过程尽量快。

那么到底怎么设计呢？
*/

/*
单调栈的实现

1 准备一个从小到大的栈结构
2 往里压入元素，直到一个元素小于栈顶元素
3 弹出栈顶元素，生成信息：元素压着的是左信息，让元素弹出的是右信息
4 弹出所有不符合从小到大条件的元素，依次生成信息
*/

// arr = [ 3, 1, 2, 3]
//         0  1  2  3
//  [
//     0 : [-1,  1]
//     1 : [-1, -1]
//     2 : [ 1, -1]
//     3 : [ 2, -1]
//  ]
func getNearLessNoRepeat(arr []int) [][]int {
	res := make([][]int, len(arr))
	for i := range res {
		res[i] = make([]int, 2)
	}
	// 只存位置！
	stack := list.New()
	for i := 0; i < len(arr); i++ { // 当遍历到i位置的数，arr[i]
		for stack.Len() > 0 && arr[stack.Front().Value.(int)] > arr[i] {
			j := stack.Front()
			stack.Remove(j)
			leftLessIndex := 0 // 左边最小元素的下标
			if stack.Len() > 0 {
				leftLessIndex = stack.Front().Value.(int)
			} else {
				leftLessIndex = -1
			}
			res[j.Value.(int)][0] = leftLessIndex
			res[j.Value.(int)][1] = i
		}
		stack.PushFront(i)
	}
	for stack.Len() > 0 {
		j := stack.Front()
		stack.Remove(j)
		leftLessIndex := 0
		if stack.Len() > 0 {
			leftLessIndex = stack.Front().Value.(int)
		} else {
			leftLessIndex = -1
		}
		res[j.Value.(int)][0] = leftLessIndex
		res[j.Value.(int)][1] = -1
	}
	return res
}

//有重复元素的单调栈
func getNearLess(arr []int) [][]int { // 返回的是记录 左右index 的二维数组
	res := make([][]int, len(arr))
	for i := range res {
		res[i] = make([]int, 2)
	}
	stack := list.New()             //链表里面存slice
	for i := 0; i < len(arr); i++ { // i -> arr[i] 进栈
		for stack.Len() > 0 && arr[stack.Front().Value.([]int)[0]] > arr[i] {
			popIs := stack.Front().Value.([]int)
			stack.Remove(stack.Front()) // 弹出slice
			leftLessIndex := 0          // 左边最小元素的下标
			if stack.Len() > 0 {
				leftLessIndex = stack.Front().Value.([]int)[len(stack.Front().Value.([]int))-1]
			} else {
				leftLessIndex = -1
			}
			for _, value := range popIs { // slice中value是arr的元素index
				res[value][0] = leftLessIndex
				res[value][1] = i
			}
		}
		if stack.Len() > 0 && arr[stack.Front().Value.([]int)[0]] == arr[i] { // 遇到相等的数
			stack.Front().Value = append(stack.Front().Value.([]int), i) // 追加重复的元素的index到slice，语法可能有问题
		} else {
			slice := make([]int, 0)
			slice = append(slice, i)
			stack.PushFront(slice)
		}
	}
	// 栈中剩余的元素单独处理
	for stack.Len() > 0 {
		popIs := stack.Front().Value.([]int)
		stack.Remove(stack.Front()) // 弹出slice
		leftLessIndex := 0
		if stack.Len() > 0 {
			leftLessIndex = stack.Front().Value.([]int)[len(stack.Front().Value.([]int))-1]
		} else {
			leftLessIndex = -1
		}
		for _, value := range popIs { //slice中的值都是相等的
			res[value][0] = leftLessIndex
			res[value][1] = -1
		}
	}
	return res
}
