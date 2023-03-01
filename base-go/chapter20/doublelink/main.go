package main

import (
	"fmt"
)

// 查找功能

type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode
	next     *HeroNode // 表示指向下个结点
}

// 给链表插入一个结点
// 编写第一种方式 在单链表的最后加入
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	// 思路
	// 1 先找到该链表的最后这个结点
	// 2 创建一个辅助结点【跑龙套】
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	// 3 将 newHeroNode 加入到链表的最后
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

// 编写第二种方式 根据 no 的编号从小到大插入
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	// 思路
	// 1 先找到适当的结点
	// 2 创建一个辅助结点【跑龙套】
	temp := head
	flag := true
	// 让插入的结点的 no，和temp 的下个结点的 no比较
	for {
		if temp.next == nil { // 说明到链表的最后
			break
		} else if temp.next.no >= newHeroNode.no { // 加个= 同样的编号 后加入的会在前面
			// 说明 newHeroNode 就应该插入到 temp后面
			break
		} else if temp.next.no == newHeroNode.no {
			// 说明我们链表中已经有这个 no，不能插入
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("对不起，已经存在no=", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}

}

// 删除一个结点【双向链表删除一个结点】
func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	// 找到删除结点的no，和temp的下个结点的 no 比较
	for {
		if temp.next == nil { // 说明到链表的最后
			break
		} else if temp.next.no == id {
			// 说明我们找到了
			flag = true
			break
		}
		temp = temp.next
	}

	if flag { // 找到 删除
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("要删除的id不存在")
	}
}

// 显示链表的所有结点信息
// 这里仍然使用单向链表显示方式
func ListHeroNode(head *HeroNode) {
	// 1 创建一个辅助结点【跑龙套】
	temp := head
	// 先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空空如也...")
		return
	}
	// 2 遍历这个链表
	for {
		fmt.Printf("[%d , %s , %s]==>", temp.next.no,
			temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

// 逆序打印
func ListHeroNode2(head *HeroNode) {
	// 1 创建一个辅助结点【跑龙套】
	temp := head
	// 判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空空如也...")
		return
	}

	// 2 让temp 定位到双向链表的最后结点
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	// 3 遍历这个链表
	for {
		fmt.Printf("[%d , %s , %s]==>", temp.no,
			temp.name, temp.nickname)
		// 判断是否链表头
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}

func main() {
	// 1 先创建一个头结点
	head := &HeroNode{}

	// 2 创建一个新的 HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}

	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero3)
	ListHeroNode(head)
	fmt.Println()
	ListHeroNode2(head)
}
