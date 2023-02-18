package main

import "fmt"

/*
给定一个整型数组arr，代表数值不同的纸牌排成一条线
玩家A和玩家B依次拿走每张纸牌
规定玩家A先拿，玩家B后拿
但是每个玩家每次只能拿走最左或最右的纸牌
玩家A和玩家B都绝顶聪明
请返回最后获胜者的分数。
*/

/*
方法一 暴力递归
*/
func win1(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	first := f1(arr, 0, len(arr)-1)  //返回值在二维数组的位置
	second := g1(arr, 0, len(arr)-1) //返回值在二维数组的位置
	return max(first, second)
}

// arr[L..R]，先手获得的最好分数返回
func f1(arr []int, L, R int) int {
	if L == R { //base case 动态规划二维数组填值的参考
		return arr[L]
	}
	p1 := arr[L] + g1(arr, L+1, R)
	p2 := arr[R] + g1(arr, L, R-1)
	return max(p1, p2) //先手选择最好分数返回
}

// arr[L..R]，后手获得的最好分数返回
func g1(arr []int, L, R int) int {
	if L == R { //base case 动态规划二维数组填值的参考
		return 0
	}
	p1 := f1(arr, L+1, R) // 对手拿走了L位置的数
	p2 := f1(arr, L, R-1) // 对手拿走了R位置的数
	return min(p1, p2)    // 对手会给你两个最优中的最小
}

/*
方法二 傻缓存
动态规划
*/
func win2(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	//可变参数的变化范围
	fmap := make([][]int, len(arr))
	gmap := make([][]int, len(arr))
	for i := range fmap {
		fmap[i] = make([]int, len(arr))
		gmap[i] = make([]int, len(arr))
		for j := 0; j < len(arr); j++ {
			fmap[i][j] = -1
			gmap[i][j] = -1
		}
	}
	first := f2(arr, 0, len(arr)-1, fmap, gmap)
	second := g2(arr, 0, len(arr)-1, fmap, gmap)
	return max(first, second)
}
func f2(arr []int, L, R int, fmap, gmap [][]int) int {
	if fmap[L][R] != -1 {
		return fmap[L][R]
	}
	ans := 0
	if L == R {
		ans = arr[L]
	} else {
		p1 := arr[L] + g2(arr, L+1, R, fmap, gmap)
		p2 := arr[R] + g2(arr, L, R-1, fmap, gmap)
		ans = max(p1, p2)
	}
	fmap[L][R] = ans
	return ans
}
func g2(arr []int, L, R int, fmap, gmap [][]int) int {
	if gmap[L][R] != -1 {
		return gmap[L][R]
	}
	ans := 0
	if L != R {
		p1 := f2(arr, L+1, R, fmap, gmap)
		p2 := f2(arr, L, R-1, fmap, gmap)
		ans = min(p1, p2)
	}
	gmap[L][R] = ans
	return ans
}

/*
方法三 动态规划
*/
func win3(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	fmap := make([][]int, len(arr))
	gmap := make([][]int, len(arr))
	for i := range fmap {
		fmap[i] = make([]int, len(arr))
		gmap[i] = make([]int, len(arr))
		fmap[i][i] = arr[i]
	}
	for startCol := 1; startCol < len(arr); startCol++ {
		L := 0
		R := startCol
		for R < len(arr) {
			fmap[L][R] = max(arr[L]+gmap[L+1][R], arr[R]+gmap[L][R-1]) //从表里拿值
			gmap[L][R] = min(fmap[L+1][R], fmap[L][R-1])               //从表里拿值
			L++
			R++
		}
	}
	return max(fmap[0][len(arr)-1], gmap[0][len(arr)-1])
}

func main() {
	arr := []int{5, 7, 4, 5, 8, 1, 6, 0, 3, 4, 6, 1, 7}
	fmt.Println(win1(arr))
	fmt.Println(win2(arr))
	fmt.Println(win3(arr))
}

//返回最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//返回最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
