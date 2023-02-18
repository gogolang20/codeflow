package main

import (
	"fmt"
	"time"
)

/*
请同学们自行搜索或者想象一个象棋的棋盘，
然后把整个棋盘放入第一象限，棋盘的最左下角是(0,0)位置
那么整个棋盘就是横坐标上9条线、纵坐标上10条线的区域
给你三个 参数 x，y，k
返回“马”从(0,0)位置出发，必须走k步
最后落在(x,y)上的方法数有多少种?
*/

// 当前来到的位置是（x,y）
// 还剩下rest步需要跳
// 跳完rest步，正好跳到a，b的方法数是多少？
// 10 * 9
func jump(a, b, k int) int {
	return process(0, 0, k, a, b)
}
func process(x, y, rest, a, b int) int {
	if x < 0 || x > 9 || y < 0 || y > 8 {
		return 0
	}
	if rest == 0 {
		if x == a && y == b {
			return 1
		} else {
			return 0
		}
	}
	ways := process(x+2, y+1, rest-1, a, b)
	ways += process(x+1, y+2, rest-1, a, b)
	ways += process(x-1, y+2, rest-1, a, b)
	ways += process(x-2, y+1, rest-1, a, b)
	ways += process(x-2, y-1, rest-1, a, b)
	ways += process(x-1, y-2, rest-1, a, b)
	ways += process(x+1, y-2, rest-1, a, b)
	ways += process(x+2, y-1, rest-1, a, b)
	return ways
}

/*
方法二
*/
func dps(a, b, k int) int {
	dp := make([][][]int, 10)
	for i := range dp {
		dp[i] = make([][]int, 9)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
		}
	}
	dp[a][b][0] = 1                    // 对应 if x == a && y == b
	for rest := 1; rest <= k; rest++ { //第0层已经设置完毕
		for x := 0; x < 10; x++ {
			for y := 0; y < 9; y++ {
				ways := pick(dp, x+2, y+1, rest-1)
				ways += pick(dp, x+1, y+2, rest-1)
				ways += pick(dp, x-1, y+2, rest-1)
				ways += pick(dp, x-2, y+1, rest-1)
				ways += pick(dp, x-2, y-1, rest-1)
				ways += pick(dp, x-1, y-2, rest-1)
				ways += pick(dp, x+1, y-2, rest-1)
				ways += pick(dp, x+2, y-1, rest-1)
				dp[x][y][rest] = ways
			}
		}
	}
	return dp[0][0][k] // 对应 return process(0, 0, k, a, b)
}

func pick(dp [][][]int, x, y, rest int) int {
	if x < 0 || x > 9 || y < 0 || y > 8 {
		return 0
	}
	return dp[x][y][rest]
}

func main() {
	start := time.Now()
	fmt.Println(jump(7, 7, 10)) //复杂度 O(8^K)
	fmt.Println(time.Since(start))

	fmt.Println(dps(7, 7, 10)) //复杂度 O(K)
	fmt.Println(time.Since(start))
}
