package main

/*
给定一个字符串str，返回这个字符串的最长回文子序列长度
比如 ： str = “a12b3c43def2ghi1kpm”
最长回文子序列是“1234321”或者“123c321”，返回长度7
*/

// 测试链接：https://leetcode.com/problems/longest-palindromic-subsequence/
/*
方法一 测试超时
*/
func lpsl1(s string) int {
	if s == "" {
		return 0
	}
	str := []byte(s)
	return f(str, 0, len(str)-1)
}

// str[L..R]最长回文子序列长度返回
func f(str []byte, L, R int) int {
	if L == R {
		return 1
	}
	if L == R-1 {
		if str[L] == str[R] {
			return 2
		} else {
			return 1
		}
	}
	p1 := f(str, L+1, R-1) //第一种 L 不结尾  R 不结尾
	p2 := f(str, L, R-1)   //第二种 L 结尾    R 不结尾
	p3 := f(str, L+1, R)   //第三种 L 不结尾  R 结尾
	p4 := 0                //第四种 L 结尾    R 结尾
	if str[L] == str[R] {  //相等才有可能性
		p4 = 2 + f(str, L+1, R-1)
	}
	return max(max(p1, p2), max(p3, p4))
}

/*
方法二
通过测试
*/
func lpsl2(s string) int {
	if s == "" {
		return 0
	}
	str := []byte(s)
	N := len(str)
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, N)
	}
	dp[N-1][N-1] = 1
	for i := 0; i < N-1; i++ {
		dp[i][i] = 1
		if str[i] == str[i+1] {
			dp[i][i+1] = 2
		} else {
			dp[i][i+1] = 1
		}
	}
	for L := N - 3; L >= 0; L-- {
		for R := L + 2; R < N; R++ {
			dp[L][R] = max(dp[L][R-1], dp[L+1][R]) //不需要比较左下的位置
			if str[L] == str[R] {                  //如果有可能性4
				dp[L][R] = max(dp[L][R], 2+dp[L+1][R-1])
			}
		}
	}
	return dp[0][N-1] //对应 return f(str, 0, len(str)-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
