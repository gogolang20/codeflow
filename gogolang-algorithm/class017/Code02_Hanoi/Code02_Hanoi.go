package main

import (
	"fmt"
	"strconv"
)

/*
暴力递归就是尝试
1，把问题转化为规模缩小了的同类问题的子问题
2，有明确的不需要继续进行递归的条件(base case)
3，有当得到了子问题的结果之后的决策过程
4，不记录每一个子问题的解
*/

/*
汉诺塔问题
打印n层汉诺塔从最左边移动到最右边的全部过程
*/

func main() {
	Hannoi(3, "left", "right", "mid")
}

/*
三大步
1) 1 - (N-1)  从 from 移动到 other
2) N  从 from 移动到 to
3) 1 - (N-1)  从 other 移动到 to
（移动 2的N次方减一 步）
*/

func Hannoi(n int, from, to, other string) {
	if n == 1 {
		fmt.Println("move 1 from " + from + " to " + to)
	} else {
		Hannoi(n-1, from, other, to)
		fmt.Println("move " + strconv.Itoa(n) + " from " + from + " to " + to)
		Hannoi(n-1, other, to, from)
	}
}
