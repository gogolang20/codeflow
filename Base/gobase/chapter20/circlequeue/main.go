package main

import (
	"errors"
	"fmt"
	"os"
)

// 结构体管理队列
type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int // 队列头部
	tail    int // 队列尾部
}

// 入队列 AddQueue/Push  出队列 GetQueue/Pop
func (cr *CircleQueue) Push(val int) (err error) {
	if cr.IsFull() {
		return errors.New("queue full")
	}
	// 分析出 this.tail 在队列尾部，但是不包含最后一个元素
	cr.array[cr.tail] = val // 把值给尾部
	cr.tail = (cr.tail + 1) % cr.maxSize
	return
}

func (cr *CircleQueue) Pop() (val int, err error) {
	if cr.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	// 取出 head指向队首，并且含队首元素
	val = cr.array[cr.head]
	cr.head = (cr.head + 1) % cr.maxSize
	return
}

// 显示队列
func (cr *CircleQueue) ListQueue() {
	fmt.Println("环形队列情况如下：")
	// 取出当前队列有多少个元素
	size := cr.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	// 设计一个辅助的变量 指向head
	tempHead := cr.head
	for i := cr.head; i < size; i++ {
		fmt.Printf("array[%d]=%d\t", tempHead, cr.array[tempHead])
		tempHead = (tempHead + 1) % cr.maxSize
	}
	fmt.Println()
}

// 判断环形队列为满
func (cr *CircleQueue) IsFull() bool {
	return (cr.tail+1)%cr.maxSize == cr.head
}

// 判断环形队列为空
func (cr *CircleQueue) IsEmpty() bool {
	return cr.tail == cr.head
}

// 取出环形队列有多少个元素
func (cr *CircleQueue) Size() int {
	return (cr.tail + cr.maxSize - cr.head) % cr.maxSize
}

func main() {
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}

	var key string
	var val int
	for {
		fmt.Println("1:输入add 表示添加数据到队列")
		fmt.Println("2:输入get 表示从队列获取数据")
		fmt.Println("3:输入show 表示显示队列")
		fmt.Println("4:输入exit 退出")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入您要如队列数")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("AddQueue ok")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
