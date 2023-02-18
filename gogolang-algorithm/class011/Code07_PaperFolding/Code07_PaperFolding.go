package main

import "fmt"

/*
请把一段纸条竖着放在桌子上，然后从纸条的下边向上方对折1次，压出折痕后展开。此时折痕是凹下去的，即折痕突起的方向指向纸条的背面。 如果从纸条的下边向上方连续对折2次，压出折痕后展开，此时有三条折痕，从上到下依次是下折痕、下折痕和上折痕。
给定一个输入参数N，代表纸条都从下边向上方连续对折N次。 请从上到下打印所有折痕的方向。
例如:N=1时，打印: down N=2时，打印: down down up

二叉树的中序遍历
头节点是 凹
所有左子树的头都是 凹
所有右子树的头都是 凸
*/

// 当前你来了一个节点，脑海中想象的！
// 这个节点在第i层，一共有N层，N固定不变的
// 这个节点如果是凹的话，down = T
// 这个节点如果是凸的话，down = F
// 函数的功能：中序打印以你想象的节点为头的整棵树！

func printAllFolds(n int) {
	process(1, n, true)
}

// 当前你来了一个节点，脑海中想象的！
// 这个节点在第i层，一共有N层，N固定不变的
// 这个节点如果是凹的话，down = T
// 这个节点如果是凸的话，down = F
// 函数的功能：中序打印以你想象的节点为头的整棵树！
func process(i, n int, down bool) {
	if i > n {
		return
	}
	process(i+1, n, true)
	if down {
		fmt.Print("凹 ")
	} else {
		fmt.Print("凸 ")
	}
	process(i+1, n, false)
}

func main() {
	N := 4
	printAllFolds(N)
}
